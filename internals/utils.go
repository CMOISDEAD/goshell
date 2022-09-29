package internals

import (
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

func GetPwd() string {
	var res string
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ad := strings.Split(path, "/")
	if len(ad) >= 3 {
		res = "../"
		res += ad[len(ad)-3] + "/"
		res += ad[len(ad)-2] + "/"
		res += ad[len(ad)-1] + "/"
	}
	return res
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
