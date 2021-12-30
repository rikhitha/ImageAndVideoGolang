package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func UploadFiles(response http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(10 * 1024 * 1024)

	files := request.MultipartForm.File["myFile"]

	for _, file := range files {

		fmt.Println("File Name", file.Filename)
		fmt.Println("File Size", file.Size)
		fmt.Println("File Type", file.Header.Get("Content-Type"))
		fmt.Println("********************************")

		//Saving files to server

		f, _ := file.Open()

		tempFile, err := ioutil.TempFile("uploads", "upload-*.MP4")

		if err != nil {
			fmt.Println(err)
		}

		defer tempFile.Close()

		fileBytes, err2 := ioutil.ReadAll(f)

		if err2 != nil {
			fmt.Println(err)

		}

		tempFile.Write(fileBytes)

		//response.Header().Add("Content-Type", "video/mp4")
		//response.Write(fileBytes)

		http.ServeContent(response, request, "video1.MP4", time.Now(), f)
		//http.ServeFile(response, request, "video1.MP4")

		//contentType := http.DetectContentType(fileBytes)
		//response.Header().Set("Content-type", contentType)
		//response.Write(fileBytes)
	}

	fmt.Println("check uploads folder to view videos")
}
