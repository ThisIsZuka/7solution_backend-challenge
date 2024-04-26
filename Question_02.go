package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Main02Handler(c *gin.Context) {

	input := c.PostForm("input")
	if input == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request input"})
		return
	}

	err := VerifyInput(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DecodeText(input)

	c.JSON(http.StatusOK, input)
}

func VerifyInput(input string) error {

	sp := []string{"L", "R", "="}

	for _, char := range input {
		found := false
		for _, s := range sp {
			if string(char) == s {
				found = true
				break
			}
		}

		if !found {
			return errors.New("invalid input")
		}
	}
	return nil
}

func DecodeText(input string) {
	// สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
	// สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
	// สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา

	// i_f := 0
	// i_m := 0
	// i_l := 0

	ans := ""

	for f := 0; f <= 9; f++ {
		for i, char := range input {

			if i == 0 {
				if string(char) == "L" {
					ans += strconv.Itoa(1)
				} else {
					ans += strconv.Itoa(f)
				}
				continue
			}

			i_r_char := string(ans[i-1])
			i_r_charInt := strconv.Atoi(string(char))
			if string(char) == "L" {
				ans += strconv.Itoa(1)
			} else if string(char) == "R" {
				ans += strconv.Itoa(2)
			} else {
				ans += strconv.Itoa(3)
			}
		}
	}

	fmt.Println(ans)
}
