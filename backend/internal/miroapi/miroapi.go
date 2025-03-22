package miroapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"html"
)

type MiroAPIInterface interface {
	GetWidgets(boardID, accessToken string) ([]Widget, error)
}

// miroの全体の情報
type Widget struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type MiroAPI struct{
	Token string
}

// MiroAPIのコンストラクタ
func NewMiroAPI(token string) MiroAPIInterface {
	return &MiroAPI{Token: token}
}

// HTMLタグを除く処理
func removeHTMLTags(input string) string {
	decoded := html.UnescapeString(input)
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllString(decoded, "")
}

// 指定されたボードIDとアクセストークンでMiro APIから全体情報を取得する
func (api *MiroAPI) GetWidgets(boardID, accessToken string) ([]Widget, error) {
	url := fmt.Sprintf("https://api.miro.com/v1/boards/%s/widgets", boardID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ボードの情報を取れませんでした status code: %d", resp.StatusCode)
	}

	var widgetsResp struct {
		Data []Widget `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&widgetsResp); err != nil {
		return nil, err
	}

	// HTMLタグを削除
	for i := range widgetsResp.Data {
		widgetsResp.Data[i].Text = removeHTMLTags(widgetsResp.Data[i].Text)
	}

	// 取得したウィジェット数をログ出力してデバッグ
	fmt.Printf("取得したウィジェット件数: %d\n", len(widgetsResp.Data))

	// 各ウィジェットの情報をデバッグ出力
	for _, widget := range widgetsResp.Data {
		fmt.Printf("Widget ID: %s,  Text: %s\n", widget.ID, widget.Text)
	}

	return widgetsResp.Data, nil
}