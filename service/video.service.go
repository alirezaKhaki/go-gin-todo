package service

import "github.com/alirezaKhaki/go-gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{}
}

func (s *videoService) Save(video entity.Video) entity.Video {
	s.videos = append(s.videos, video)
	return video
}

func (s *videoService) FindAll() []entity.Video {
	return s.videos
}
