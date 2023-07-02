package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Получение списка книг автора по его Id
func (h *Handler) getBookByAuthor(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	books, err := h.services.GetBooksByAuthor(authorID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

// Получение кол-ва книг автора по его Id
func (h *Handler) getBookCountByAuthor(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	booksCount, err := h.services.GetAuthorBooksCount(authorID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"count": booksCount,
	})
}
