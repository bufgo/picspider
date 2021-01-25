package models

import (
	"fmt"
	"time"
)

// PhotoAlbum model
type PhotoAlbum struct {
	ID        int64     `xorm:"pk autoincr" json:"id"`
	AlbumName string    `json:"album_name"`
	AlbumURL  string    `xorm:"notnull unique" json:"album_url"`
	CreatedAt time.Time `xorm:"created" json:"created"`
	UpdatedAt time.Time `xorm:"updated" json:"updated"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted"`
}

// SaveSearchResult is save search result
func SaveSearchResult(photoAlbum PhotoAlbum) {
	_, err := engine.InsertOne(photoAlbum)
	if err != nil {
		fmt.Printf("PhotoAlbum.SaveSearchResult, insert data err: %v \n", err)
	}
}

// GetPhotoAlbumID is get the photo album id
func GetPhotoAlbumID(band string) int64 {
	photoAlbum := &PhotoAlbum{AlbumURL: band}
	has, err := engine.Get(photoAlbum)
	if !has {
		fmt.Printf("PhotoAlbum.GetPhotoAlbumID, fail to get album obj: %v", err)
	}

	return photoAlbum.ID
}
