package main

import (
	"fmt"
	"go-starter/utils"
)

func main() {
	square := calculateSquare(20, 30)
	fmt.Println("square: ", square)

	grade := checkGrade(50)
	fmt.Println("grade: ", grade)

	tax := calculateTax(234523.24)
	fmt.Println("tax: ", tax)

	student := manageScoreStudent()
	fmt.Println("student: ", student)

	// m = make(map[string]int)
	// test :=
}

// Variables and Data Types
func calculateSquare(width int, height int) int {
	fmt.Println("width:", width, "height:", height)
	return width * height
}

// Control Structures
func checkGrade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	}
	return "F"
}

// Functions
func calculateTax(salary float64) string {
	if salary > 500000 {
		return "15%"
	} else if salary <= 500000 {
		return "10%"
	}
	return "5%"
}

// Arrays, Slices, Maps
// โจทย์: ระบบจัดการคะแนนนักเรียน
// - สร้าง slice เก็บคะแนนนักเรียน
// - สร้าง map เก็บชื่อนักเรียนและคะแนน
// - เพิ่มคะแนน 5 คน
// - แสดงค่าเฉลี่ย
// - แสดงคะแนนสูงสุด-ต่ำสุด
// - แสดงรายชื่อนักเรียนที่ได้คะแนนมากกว่าค่าเฉลี่ย
type Student struct {
	score int
}

func manageScoreStudent() any {

	scores := make(map[string]int, 5)
	student_score := make([]int, 0, 5)

	scores["Alex"] = 24
	scores["Otto"] = 1
	scores["Michel"] = 23
	scores["Daniel"] = 52
	scores["Milly"] = 34

	for _, score := range scores {
		student_score = append(student_score, score)
	}

	min, max, avg := utils.FindMinMaxAvg(student_score)

	fmt.Println("min:", min, "max:", max, "avg:", avg)
	studentScoreGreaterThanAvg := utils.FindStudentScoreGreaterThanAvg(scores, avg)

	for name := range studentScoreGreaterThanAvg {
		fmt.Println("studentScoreGreaterThanAvg: ", name)
	}
	return scores
}
