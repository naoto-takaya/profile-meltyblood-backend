package infrastructure

import (
	"fmt"
	"net/http"
)

func Authentication() {

}

func UploadMedia(data string) {
	url := "https://upload.twitter.com/1.1/media/upload.json"
	req, _ := http.NewRequest("POST", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	fmt.Print(resp)

}
