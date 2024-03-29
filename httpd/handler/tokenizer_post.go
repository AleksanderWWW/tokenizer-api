package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AleksanderWWW/tokenizer-api/backend/tokenizer"
)

type tokenizerRequest struct {
	Text           string `json:"text"`
	Model          string `json:"model"`
	AddPrefixSpace bool   `json:"addPrefixSpace"`
	TrimOffsets    bool   `json:"trimOffsets"`
}

func TokenizerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := tokenizerRequest{}
		c.Bind(&requestBody)

		tk, err := tokenizer.GetModelSwitch(requestBody.Model, requestBody.AddPrefixSpace, requestBody.TrimOffsets)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("unrecognized model: %s", requestBody.Model),
			})
			return
		}

		en, err := tk.EncodeSingle(requestBody.Text)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"text":           requestBody.Text,
			"model":          requestBody.Model,
			"addPrefixSpace": requestBody.AddPrefixSpace,
			"trimOffsets":    requestBody.TrimOffsets,
			"tokens":         en.Tokens,
		})
	}
}
