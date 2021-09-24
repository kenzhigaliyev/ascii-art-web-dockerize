package main

import (
	"fmt"
	student "student/web"
)

func main() {
	fmt.Println("http://localhost:7777/ is listening...")
	student.MainFunc()
}
