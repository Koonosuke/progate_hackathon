package miroapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MiroAPIInterface interface {
	GetWidgets(boardID, accessToken string) ([]Widget, error)
}

// miroの全体の情報
type Widget struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
	WidgetID string `json:"widgetId"`
}

type MiroAPI struct{
	Token string
}

// MiroAPIのコンストラクタ
func NewMiroAPI(token string) MiroAPIInterface {
	return &MiroAPI{Token: token}
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

	// 取得したウィジェット数をログ出力してデバッグ
	fmt.Printf("取得したウィジェット件数: %d\n", len(widgetsResp.Data))
	
	return widgetsResp.Data, nil
}