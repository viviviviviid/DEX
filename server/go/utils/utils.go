package utils

import "time"

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}
