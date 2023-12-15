package test

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_url(t *testing.T) {
	fileName := "截图 2023-11-30 16-38-16.png"
	encodedFileName := url.QueryEscape(fileName)
	link := fmt.Sprintf("https://xichen-server/your-prefix/%s", encodedFileName)
	fmt.Println("Encoded Link:", link)
}
