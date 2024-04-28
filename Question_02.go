package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main02Handler(c *gin.Context) {

	input := c.PostForm("input")
	if input == "" {
		c.JSON(http.StatusOK, gin.H{"error": "request input"})
		return
	}

	err := VerifyInput(input)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	ans := decode(input)
	c.JSON(http.StatusOK, ans)
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

func decode(input string) []int {
	// สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
	// สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
	// สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา

	Numbers := make([]int, 0)
	Numbers = append(Numbers, 0)
	for i := 1; i < len(input); i++ {
		fmt.Println(i)
		switch input[i] {
		case 'L':
			if input[i-1] == 'L' && input[i] != 'L' {
				Numbers[len(Numbers)-1]++
				Numbers = append(Numbers, Numbers[len(Numbers)-1]-1)
			} else {
				Numbers[len(Numbers)-1] += 2
				Numbers = append(Numbers, Numbers[len(Numbers)-1]-1)
			}
		case 'R':
			Numbers = append(Numbers, i)
			Numbers = append(Numbers, i+1)
		case '=':
			Numbers = append(Numbers, i)
			Numbers = append(Numbers, i)
		}
		break
	}
	return Numbers
}
