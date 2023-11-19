package handler

import (
	"net/http"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/gin-gonic/gin"
)

// Регистрация
func (h *Handler) signUp(c *gin.Context) {

	// Создание переменной типа User (из файла pkg/user.go
	var input todo.User

	// Берём данные из контекста *gin.Context и распаршиваем JSON в переменную User
	if err := c.BindJSON(&input); err != nil {

		// Наша функция вывода ошибки, на случай, если пришли не валидные данные
		//	правила валидации опредиляются при создании типа User: `binding:"required"`
		// Функция создаётся в файле pkg/handler/response.go
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

// Аунтитификация
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	// Берём данные из контекста *gin.Context и распаршиваем JSON в переменную User
	if err := c.BindJSON(&input); err != nil {

		// Наша функция вывода ошибки, на случай, если пришли не валидные данные
		//	правила валидации опредиляются при создании типа User: `binding:"required"`
		// Функция создаётся в файле pkg/handler/response.go
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Autorization.GentrateToken(input.Username, input.Password)
	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
