package usecase

import (
	"fmt"
	"strings"

	"github.com/matthewTechCom/progate_hackathon/internal/model"
	"github.com/matthewTechCom/progate_hackathon/internal/miroapi"
	"github.com/matthewTechCom/progate_hackathon/internal/repository"
)

type WidgetUsecaseInterface interface {
	ProcessAndSave(boardID, accessToken string) ([]int, error)
}

type WidgetUsecase struct {
	BoardRepo  repository.BoardRepositoryInterface
	StickyRepo repository.StickyRepositoryInterface
	MiroAPI    miroapi.MiroAPIInterface
	Embedder   EmbeddingService // ✅ Embedding 用に追加
}

func NewWidgetUsecase(
	boardRepo repository.BoardRepositoryInterface,
	stickyRepo repository.StickyRepositoryInterface,
	miro miroapi.MiroAPIInterface,
	embedder EmbeddingService, // ✅ 追加
) WidgetUsecaseInterface {
	return &WidgetUsecase{
		BoardRepo:  boardRepo,
		StickyRepo: stickyRepo,
		MiroAPI:    miro,
		Embedder:   embedder, // ✅ 追加
	}
}

func (u *WidgetUsecase) ProcessAndSave(boardID, accessToken string) ([]int, error) {
	// DBに存在するかチェックし、なければ新規保存する
	board, err := u.BoardRepo.GetByMiroID(boardID)
	if err != nil {
		newBoard := &model.Board{
			MiroBoardID: boardID,
			Title:       "",
			Description: "",
		}
		boardIDInt, err := u.BoardRepo.Save(newBoard)
		if err != nil {
			return nil, fmt.Errorf("boardの保存に失敗: %v", err)
		}
		board = &model.Board{ID: boardIDInt, MiroBoardID: boardID}
	}

	// Miro APIからウィジェット情報を取得
	widgets, err := u.MiroAPI.GetWidgets(boardID, accessToken)
	if err != nil {
		return nil, fmt.Errorf("miro APIから情報取得に失敗: %v", err)
	}

	var stickies []*model.Sticky
	for _, widget := range widgets {
		var category string
		if strings.Contains(widget.Text, "改善") {
			category = "改善点"
		} else if strings.Contains(widget.Text, "反省") {
			category = "反省点"
		} else {
			continue
		}

		// ✅ embedding 生成
		embedding, err := u.Embedder.GetEmbedding(widget.Text)
		if err != nil {
			fmt.Printf("embedding取得失敗（スキップ）: %v\n", err)
			continue
		}

		sticky := &model.Sticky{
			BoardID:      board.ID,
			MiroStickyID: widget.ID,
			Content:      widget.Text,
			Category:     category,
			Embedding:    embedding, // ✅ embedding セット
		}

		stickies = append(stickies, sticky)
	}

	// DB保存
	savedIDs, err := u.StickyRepo.Save(stickies)
	if err != nil {
		return nil, fmt.Errorf("stickyの保存に失敗: %v", err)
	}

	return savedIDs, nil
}
