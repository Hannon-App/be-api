package service

import (
	"Hannon-app/features/items"

	"github.com/go-playground/validator"
)

type ItemService struct {
	itemData items.ItemDataInterface
	validate *validator.Validate
}

func New(repo items.ItemDataInterface) items.ItemServiceInterface {
	return &ItemService{
		itemData: repo,
	}
}

func (service *ItemService) GetAllItem() ([]items.ItemCore, error) {
	result, err := service.itemData.ReadAll()
	if err != nil {
		return []items.ItemCore{}, err
	}
	return result, err
}

func (service *ItemService) Delete(id uint) error {
	err := service.itemData.Delete(id)
	return err
}

func (service *ItemService) GetById(id uint) (items.ItemCore, error) {
	return service.itemData.SelectById(id)
}

func (service *ItemService) Create(input items.ItemCore) (items.ItemCore, error) {

	result, err := service.itemData.Insert(input)
	if err != nil {
		return items.ItemCore{}, err
	}
	return result, err
}

func (service *ItemService) Update(id uint, input items.ItemCore) (items.ItemCore, error) {
	result, err := service.itemData.UpdateDataItem(id, input)
	if err != nil {
		return items.ItemCore{}, err
	}
	return result, nil
}
