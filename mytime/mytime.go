package mytime

import (
	"fmt"
	"time"
)

func MyTime() {
	fmt.Println("Welcome to time study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))                 // mm-dd-YYYY
	fmt.Println(presentTime.Format("01-02-2006 Monday"))          // mm-dd-YYYY Day
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // mm-dd-YYYY hh:mm:ss Day

	createdDate := time.Date(2015, time.September, 25, 18, 23, 0, 0, time.Local)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
