package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) createList(c *gin.Context) {

}

func (h *Handler) getAllList(c *gin.Context) {
	// test_str := viper.GetString("port")
	test_str := "Дароу"
	sendTelegramMessage(test_str)
}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}

func sendTelegramMessage(text string) error {
	token_for_telegram_bot := viper.GetString("token_for_telegram_bot")
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=91005356&text=%s&parse_mode=Markdown", token_for_telegram_bot, text)
	_, err := http.Get(url)
	return err
}
