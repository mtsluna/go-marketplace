package repositories

import (
	"bytes"
	"encoding/json"
	"github.com/mtsluna/go-market/src/domain/contracts/dto"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type ImageRepo struct {

}

const IMGUR_URL = "https://api.imgur.com/3/image"

func (repo *ImageRepo) Upload(file *multipart.FileHeader) dto.ImageDto {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	fileOpen, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, fileOpen)
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", IMGUR_URL, body)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Client-ID 9d5405d688dce1d")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var responseDto dto.ImageDto

	if err := json.NewDecoder(response.Body).Decode(&responseDto); err != nil {
		log.Fatal(err)
	}

	return responseDto

}
