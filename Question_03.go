package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Main03Handler(c *gin.Context) {

	text, err := getTextFromURL()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	lowerStr := strings.ToLower(text)
	jsonData := groupText(lowerStr)

	c.JSON(http.StatusOK, jsonData)
}

func getTextFromURL() (string, error) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func groupText(text string) map[string]map[string]int {
	groups := make(map[string]map[string]int)

	group_word := strings.FieldsFunc(text, func(r rune) bool {
		return r == ',' || r == '.' || r == ' ' || r == '\n'
	})

	// fmt.Println(group_word)

	// นับจำนวนข้อความในแต่ละกลุ่ม
	for _, word := range group_word {
		if groups["beef"] == nil {
			groups["beef"] = make(map[string]int)
		}
		groups["beef"][word]++
	}

	return groups
}
