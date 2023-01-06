package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/patyumi/video-encoder/domain"
	uuid "github.com/satori/go.uuid"
)

type VideoRepostory interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

// Vai servir para injetar algo que eu nao saquei
type VideoRepostoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepostoryDb {
	return &VideoRepostoryDb{Db: db}
}

func (repo VideoRepostoryDb) Insert(video *domain.Video) (*domain.Video, error) {

	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := repo.Db.Create(video).Error

	if err != nil {
		return nil, err
	}

	return video, nil

}

func (repo VideoRepostoryDb) Find(id string) (*domain.Video, error) {

	var video domain.Video

	repo.Db.Preload("Jobs").First(&video, "id = ?", id)

	if video.ID == "" {
		return nil, fmt.Errorf("Video does not exist")
	}

	return &video, nil

}
