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
	affected, err := engine.InsertOne(photoAlbum)
	fmt.Println(affected)
	if err != nil {
		fmt.Printf("PhotoAlbum.SaveSearchResult, insert data err: %v \n", err)
	}
}
