package service

import (
	"github.com/alirezaKhaki/go-gin/entity"
	"github.com/jinzhu/gorm"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
	db     *gorm.DB
}

func NewVideoService(db *gorm.DB) VideoService {
	return &videoService{db: db}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.videos = append(s.videos, video)
	return video
}

func (s *videoService) FindAll() []entity.Video {
	video := entity.Video{Title: "test", Description: "test", URL: "www.test.com"}
	s.videos = append(s.videos, video)
	return s.videos
}
