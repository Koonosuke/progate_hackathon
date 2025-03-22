package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/matthewTechCom/progate_hackathon/internal/usecase"
)

type RAGController struct {
	Usecase *usecase.RAGUsecase
}

// /update_embedding?id=1&query=〜 に対応
func (rc *RAGController) UpdateEmbedding(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "query is required"})
	}

	if err := rc.Usecase.EnsureEmbedding(id, query); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update embedding"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Embedding updated"})
}

// /search?query=〜 に対応
func (rc *RAGController) SearchWithRAG(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "queｚry is required"})
	}

	results, err := rc.Usecase.SearchWithRAG(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search data"})
	}

	return c.JSON(http.StatusOK, results)
}

// /search_with_advice?query=〜 に対応
func (rc *RAGController) SearchWithAdvice(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "query is required"})
	}

	data, err := rc.Usecase.SearchWithRAGAndAdvice(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to search or generate advice"})
	}

	return c.JSON(http.StatusOK, data)
}
