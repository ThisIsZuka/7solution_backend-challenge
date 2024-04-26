// Question_01.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func MainHandler(c *gin.Context) {
	data, err := ReadJsonfile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดขณะอ่านไฟล์"})
		return
	}

	ans := RoadSum(data)

	c.JSON(http.StatusOK, ans)
}

func RoadSum(data [][]int) int {

	// assum := [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 9},
	// 	{26, 53, 6, 34},
	// 	{10, 51, 87, 86, 81},
	// 	{61, 95, 66, 57, 25, 68},
	// 	{90, 81, 80, 38, 92, 67, 73},
	// 	{30, 28, 51, 76, 81, 18, 75, 44},
	// 	{84, 14, 95, 87, 62, 81, 17, 78, 58},
	// 	{21, 46, 71, 58, 2, 79, 62, 39, 31, 9},
	// 	{56, 34, 35, 53, 78, 31, 81, 18, 90, 93, 15},
	// }

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
		fmt.Println("Round Sum:", ans_sum)
	}

	fmt.Println("Sum:", ans_sum)

	return ans_sum
}

func ReadJsonfile() ([][]int, error) {
	// อ่านข้อมูล JSON จากไฟล์
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
