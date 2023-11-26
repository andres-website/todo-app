package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	// fmt.Println("createList()")
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {

		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// call service method
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	/*c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})*/
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllList(c *gin.Context) {
	// test_str := viper.GetString("port")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	test_str := "Дароу"
	sendTelegramMessage(test_str)

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

func (h *Handler) updateList(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

func sendTelegramMessage(text string) {
	// token_for_telegram_bot := viper.GetString("token_for_telegram_bot")
	token_for_telegram_bot := os.Getenv("TOKEN_TELEGRAM_FOR_BOT")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=91005356&text=%s&parse_mode=Markdown", token_for_telegram_bot, text)
	http.Get(url)
	// return err
}
