package services

import "github.com/dhikramagroup/gin-server/entity"

type VideosService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideosService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video

}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
