package models

import (
	"fmt"
	"time"
)

// Photo model
type Photo struct {
	ID           int64     `xorm:"pk autoincr" json:"id"`
	PhotoName    string    `json:"photo_name"`
	PhotoURL     string    `xorm:"notnull unique" json:"photo_url"`
	PhotoAlbumID int64     `xorm:"index notnull"`
	CreatedAt    time.Time `xorm:"created" json:"created"`
	UpdatedAt    time.Time `xorm:"updated" json:"updated"`
	DeletedAt    time.Time `xorm:"deleted" json:"deleted"`
}

// SavePhotoResult is save photo result
func SavePhotoResult(photo Photo) {
	_, err := engine.InsertOne(photo)
	if err != nil {
		fmt.Printf("Photo.SavePhotoResult, insert data err: %v \n", err)
	}
}
