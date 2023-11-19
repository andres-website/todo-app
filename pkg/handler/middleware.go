package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autotizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {

	fmt.Println("userIdentity")
	header := c.GetHeader(autotizationHeader)

	fmt.Println(header)
	if header == "" {

		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {

		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	// parse token
	userId, err := h.services.Autorization.ParseToken(headerParts[1])
	if err != nil {

		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
