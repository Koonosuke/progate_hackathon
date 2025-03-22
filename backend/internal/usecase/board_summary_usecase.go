package usecase

import (
	"fmt"
	"time"

	"github.com/matthewTechCom/progate_hackathon/internal/repository"
	"github.com/matthewTechCom/progate_hackathon/internal/miroapi"
	"github.com/matthewTechCom/progate_hackathon/internal/chatgptapi"
	"github.com/matthewTechCom/progate_hackathon/internal/model"
)

type BoardSummaryUsecaseInterface interface {
	ProcessAndSave(boardID, accessToken string) ([]int, error)
}

type BoardSummaryUsecase struct {
	BoardSummaryRepo repository.BoardSummaryRepositoryInterface
	MiroAPI          miroapi.MiroAPIInterface
	ChatGPTAPI       chatgptapi.ChatGPTAPIInterface
}

func NewBoardSummaryUsecase(repo repository.BoardSummaryRepositoryInterface, miro miroapi.MiroAPIInterface, chatgpt chatgptapi.ChatGPTAPIInterface) BoardSummaryUsecaseInterface {
	return &BoardSummaryUsecase{
		BoardSummaryRepo: repo,
		MiroAPI:          miro,
		ChatGPTAPI:       chatgpt,
	}
}

// miroapiから情報を取得し、chatgptapiで要約したものをDBに保存する
func (u *BoardSummaryUsecase) ProcessAndSave(boardID, accessToken string) ([]int, error) {
	// miroapiから情報を取得
	widgets, err := u.MiroAPI.GetWidgets(boardID, accessToken)
	if err != nil {
		return nil, fmt.Errorf("miro API から情報の取得に失敗: %v", err)
	}

	var savedIDs []int

	// 複数のウィジェットを処理
	for _, widget := range widgets {
		// chatgptで要約
		summary, err := u.ChatGPTAPI.SummarizeText(widget.Text)
		if err != nil {
			return nil, fmt.Errorf("ChatGPT API による要約に失敗: %v", err)
		}

		// ウィジェットIDと要約内容をログ出力してデバッグ
		fmt.Printf("ウィジェットID: %s の要約: %s\n", widget.ID, summary)

		// modelのインスタンスを作成
		boardSummary := &model.BoardSummary{
			Summary:   summary,
			CreatedAt: time.Now(),
		}

		// リポジトリを通じてDBへ保存する
		savedID, err := u.BoardSummaryRepo.Save(boardSummary)
		if err != nil {
			return nil, fmt.Errorf("要約の保存に失敗: %v", err)
		}

		// 保存したIDを保存
		savedIDs = append(savedIDs, savedID)
	}

	// 保存されたIDのスライスを返す
	return savedIDs, nil
}