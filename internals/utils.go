package internals

import (
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

func GetPwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func GetUser() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return user.Username
}

func GetTime() string {
	now := time.Now()
	return now.Format("3:4:5 PM")
}

func FormatPath(path string) string {
	var correct_path string
	ad := strings.Split(path, "/")
	if len(ad) >= 3 {
		correct_path += ad[len(ad)-3] + "/"
		correct_path += ad[len(ad)-2] + "/"
		correct_path += ad[len(ad)-1] + "/"
	}
	return correct_path
}
