package repository

import (
	"assignment3/models"
	"encoding/json"
	"io/ioutil"
)

type Repository interface {
	GetData() (models.Response, error)
	UpdateData(models.Response) error
}

type repository struct {
	filePath string
}

func NewDataRepository(filePath string) Repository {
	return &repository{
		filePath: filePath,
	}
}

func (r *repository) GetData() (models.Response, error) {
	byte, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return models.Response{}, err
	}

	var response models.Response
	err = json.Unmarshal(byte, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) UpdateData(response models.Response) error {
	byte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.filePath, byte, 0777)
	if err != nil {
		return err
	}

	return nil
}
