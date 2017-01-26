package util

import (
	"github.com/tucnak/telebot"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"os"
)

func UrlExt(url string) string {
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '.' {
			return url[i:]
		}
	}
	return ""
}

func SendPhoto(url string, message telebot.Message, bot *telebot.Bot) error {
	imagefile, err := SaveImage(url)
	if err != nil {
		log.Println("Error fetching ")
		log.Println(err)
		bot.SendMessage(message.Chat, url, nil)

		return err
	}
	defer os.Remove(imagefile)

	photo, err := telebot.NewFile(imagefile)

	//	var photo = telebot.Photo{}
	//	photo.Thumbnail.File, err = telebot.NewFile(imagefile)
	if err != nil {
		log.Println("Error creating the new file ")
		log.Println(err)
		bot.SendMessage(message.Chat, url, nil)

		return err
	}
	//photo.filename=imagefile
	picture := telebot.Photo{File: photo}

	err = bot.SendPhoto(message.Chat, &picture, nil)
	if err != nil {
		log.Println("Error sending photo")
		log.Println(err)
		bot.SendMessage(message.Chat, url, nil)

		return err
	}
	return err
}

func SaveImage(url string) (string, error) {
	// don't worry about errors
	log.Println("GET: Saving " + url)
	ext := UrlExt(url)
	log.Println("GET: Extension " + ext)

	response, e := http.Get(url)
	if e != nil {
		log.Println(e)
		return "", e
	}

	defer response.Body.Close()
	f, e := ioutil.TempFile("", "img")
	defer os.Remove(f.Name())

	var path = f.Name() + ext
	if e != nil {
		log.Println(e)
		return "", e
	}
	//open a file for writing
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	file.Close()
	log.Println(url + " Saved in " + path)

	return path, err

}

func DecodeJson(url string, f func(io.ReadCloser) bool) (bool, error) {
	r, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer r.Body.Close()

	return f(r.Body), err
}
