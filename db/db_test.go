package db

import (
	"os"
	"testing"
)

func TestNewDb(t *testing.T) {

	testDb, err := NewDb("newDb.db")
	if err != nil {
		t.Fatal(err)
	}
	defer testDb.Close()

	if _, err := os.Stat(testDb.DbFile); os.IsNotExist(err) {
		t.Fatal("Failed to create database.")
	}

	os.Remove("newDb.db")
}

func TestNewImage(t *testing.T) {
	testDb, err := NewDb("newImage.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}
	os.Remove("newImage.db")
}

func BenchmarkNewImage(b *testing.B) {
	_, _ := NewDb("BenchnewImage.db")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testDb.NewImage("hash", "file.jpg", "thumb.jpg", "http://google.de", "nerdlife", "#rumkugel", "user")
	}
	os.Remove("BenchmarknewImage.db")
}

func TestGetImage(t *testing.T) {
	testDb, err := NewDb("getImage.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	_, err = testDb.GetImage(1)
	if err != nil {
		t.Fatal("Failed to select image")
	}
	os.Remove("getImage.db")
}

func TestGetImages(t *testing.T) {
	testDb, err := NewDb("newImages.db")
	if err != nil {
		t.Fatal(err)
	}

	// insert paar bilder
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	img, err := testDb.GetImages(1, 5)
	if err != nil {
		t.Fatal("Failed to select image")
	}
	if len(img) != 5 {
		t.Fatal("Requested 5 Images, got %d", len(img))
	}
	os.Remove("newImages.db")
}

func TestDeleteImage(t *testing.T) {
	testDb, err := NewDb("deleteImage.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	if !testDb.DeleteImage(1) {
		t.Fatal("Failed to delete image.")
	}
	os.Remove("deleteImage.db")
}

func TestGetImageCount(t *testing.T) {
	testDb, err := NewDb("getImageCount.db")
	if err != nil {
		t.Fatal(err)
	}

	_, err = testDb.NewImage("hash", "datei.jpg", "thumb.jpg", "http://example.tld", "nerdlife", "#rumkugel", "Churchill")
	if err != nil {
		t.Fatal("Failed to insert image.")
	}

	c, err := testDb.GetImageCount()
	if err != nil {
		t.Fatal("Failed to count images in db.")
	}
	if c != 1 {
		t.Fatalf("Failed to count images correctly, got %d expected 1\n", c)
	}

	os.Remove("getImageCount.db")
}
