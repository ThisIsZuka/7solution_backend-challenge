// Question_01.go

package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	data, err := ReadJsonfile()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "เกิดข้อผิดพลาดขณะอ่านไฟล์"})
		return
	}

	ans, err := RoadSum(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ans)
}

func RoadSum(data [][]int) (int, error) {

	var index = 0
	var ans_sum = 0

	// loop แต่ละ node
	for i := range data {

		maxValue := 0
		f_index := index
		l_index := index + 1

		// loop หาค่ามากสุดของ node นั้นๆ
		for j := f_index; j <= l_index && j < len(data[i]); j++ {
			if data[i][j] > maxValue {
				maxValue = data[i][j]
				index = j
			}
		}
		ans_sum += maxValue
	}

	return ans_sum, nil
}

func ReadJsonfile() ([][]int, error) {

	file, err := os.ReadFile("./hard.json")
	if err != nil {
		return nil, err
	}

	var data [][]int
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
