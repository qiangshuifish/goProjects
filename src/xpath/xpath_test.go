package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	xpath := ""
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	exe(srt, xpath, html, "")
}

func TestXpath(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	xpath := "//title"
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	exe(srt, xpath, html, "")
}

func TestXpathAndHandle(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title id=\"1\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title id=\"2\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title id=\"3\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title id=\"4\">asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title >asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	xpath := "//title"
	//config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	resultHandle := `{"value_content":"//title","id_attr":"//title"}`
	exe(srt, xpath, html, resultHandle)
}

func TestConfig(t *testing.T) {
	html := "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	//xpath := "//title"
	config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	srt := ""
	exe(config, srt, html, "")
}
func TestMap(t *testing.T) {
	config := "{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\",\"obj\":{\"titleArr\":[\"//title\"],\"textArr_content\":[\"//title\"],\"content_attr\":[\"//meta\"],\"title\":\"//title\",\"text_content\":\"//title\",\"head\":\"//head\",\"headText_content\":\"//head\"}}"
	configJson := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(config), &configJson) //第二个参数要地址传递
	if err != nil {
		printlnErr(err, "Parser config err")
		return
	}
	fmt.Println(configJson)
}
