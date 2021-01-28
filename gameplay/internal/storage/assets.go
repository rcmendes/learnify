package storage

import (
	"io/ioutil"
	"path/filepath"
)

//ImageRepository defines the contract of an Image Repository.
type ImageRepository interface {
	GetImageByFilename(filename string) (*[]byte, error)
}

//AudioRepository defines the contract of an Audio Repository.
type AudioRepository interface {
	GetAudioByFilename(filename string) (*[]byte, error)
}

type imageFSRepository struct {
	basePath string
}

type audioFSRepository struct {
	basePath string
}

//TODO Customize errors

// NewImageFSRepository creates a File System based Image repository instance.
func NewImageFSRepository(basePath string) ImageRepository {
	return &imageFSRepository{basePath: basePath}
}

// NewAudioFSRepository creates a File System based Audio repository instance.
func NewAudioFSRepository(basePath string) AudioRepository {
	return &audioFSRepository{basePath: basePath}
}

func (repo *imageFSRepository) GetImageByFilename(filename string) (*[]byte, error) {
	filePath := filepath.Join(repo.basePath, filename)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	return &data, nil
}

func (repo *audioFSRepository) GetAudioByFilename(filename string) (*[]byte, error) {
	filePath := filepath.Join(repo.basePath, filename)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	return &data, nil
}
