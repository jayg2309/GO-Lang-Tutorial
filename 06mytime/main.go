package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package")

	// time package will help us to get the current time
	// time.Now() will return the current time
	presentTime := time.Now()                     // the format is 2024-04-04 14:13:08.6437134 +0530 IST m=+0.003261301
	fmt.Println(presentTime)                      // pretty big and not so readable
	fmt.Println(presentTime.Format("01-02-2004")) // 04-04-2024 // 01-02-2006 is the reference date
	// we get the format
	fmt.Println(presentTime.Format("01-02-2004 Monday"))          // 04-04-2024 Thursday
	fmt.Println(presentTime.Format("01-02-2004 11:04:05 Monday")) // 04-04-2024 14:13:08 Thursday

	createdDate := time.Date(2020, time.April, 23, 9, 16, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println((createdDate.Format("01-01-2006 12:12:12"))) // 23-04-2020 09:16:00

	// build for different OS
	// GOOS="darwin" go build // this is for mac , we can click on the executable file and run it on mac
	// GOOS="linux" go build // this is for linux , we can click on the executable file and run it on linux
	// a new mytime file will be created in the same folder

}
