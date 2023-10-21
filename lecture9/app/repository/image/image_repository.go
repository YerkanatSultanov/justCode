package image

import (
	"errors"
	"justCode/lecture9/app/model"
)

type ImageRepository interface {
	GetAll() []model.Image
	GetByID(id int) (*model.Image, error)
	Create(image *model.Image) error
	Update(id int, image *model.Image) error
	Delete(id int) error
}

type ImageRepositoryImpl struct {
	images []model.Image
}

func NewImageRepository() *ImageRepositoryImpl {
	return &ImageRepositoryImpl{}
}

func (r *ImageRepositoryImpl) GetAll() []model.Image {
	return r.images
}

func (r *ImageRepositoryImpl) GetByID(id int) (*model.Image, error) {
	for i := range r.images {
		if r.images[i].Image_id == id {
			return &r.images[i], nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *ImageRepositoryImpl) Create(image *model.Image) error {
	image.Image_id = len(r.images) + 1
	r.images = append(r.images, *image)
	return nil
}

func (r *ImageRepositoryImpl) Update(id int, image *model.Image) error {
	for i := range r.images {
		if r.images[i].Image_id == id {
			r.images[i] = *image
			return nil
		}
	}
	return errors.New("Not found")
}

func (r *ImageRepositoryImpl) Delete(id int) error {
	for i := range r.images {
		if r.images[i].Image_id == id {
			r.images = append(r.images[:i], r.images[i+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}
