package usecases

import (
	"github.com/mtsluna/go-market/src/domain/contracts"
	"github.com/mtsluna/go-market/src/infrastructure/repositories"
	"mime/multipart"
)

type ImageUc struct {
	repo repositories.ImageRepo
}

func (uc *ImageUc) Upload(file *multipart.FileHeader) contracts.Image{

	imageDto := uc.repo.Upload(file)

	image := contracts.Image{
		Url: imageDto.Data.Link,
	}

	return image

}