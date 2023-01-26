package main

import (
	"errors"
	"flag"
	"io/fs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	var (
		configPath = flag.String("c", "/etc/charv/backend.conf", "backend config file path")
	)
	flag.Parse()

	err := godotenv.Load(*configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	imagesDir := os.Getenv("IMAGES_DIR")
	_, err = os.ReadDir(imagesDir)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}

	guestsDir := os.Getenv("GUESTS_DIR")
	_, err = os.ReadDir(guestsDir)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}

	storageDir := os.Getenv("STORAGE_POOLS_DIR")
	_, err = os.ReadDir(storageDir)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}

	qmpDir := os.Getenv("QMP_DIR")
	_, err = os.ReadDir(qmpDir)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}

	vncDir := os.Getenv("VNC_DIR")
	_, err = os.ReadDir(vncDir)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		log.Fatal(err.Error())
	}
}
