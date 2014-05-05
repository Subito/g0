package db

import (
	"errors"
	_ "fmt"
	"g0/config"
	"time"
)

type Image struct {

	// public Fields
	Id        int
	Hash      string
	Name      string
	Thumbnail string
	Timestamp time.Time
	Url       string
	Network   string
	Channel   string
	User      string

	// Config
	table string
	db    Db
}

func Image() *Image {
	var result *Image = &Image{}

	result.db = NewDb("g0.db")
}

// Fetch($PicId) returns one Image-struct
func (i *Image) Fetch(id int) (Image, error) {
	image, err := i.FetchLatest(id, 1)
	if err != nil {
		return Image{}, err
	}
	return image[0], nil
}

// FetchLatest($id, $n) returns n Image-structs, start at $id
// and n previous db-entries ordered by timestamp desc.
func (i *Image) FetchLatest(id, n int) ([]Image, error) {

}
