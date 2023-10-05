package service

import (
	"Hannon-app/features/items"
	"mime/multipart"
)

type ItemService struct {
	itemData items.ItemDataInterface
}

func New(repo items.ItemDataInterface) items.ItemServiceInterface {
	return &ItemService{
		itemData: repo,
	}
}

func (service *ItemService) GetAllItem(page, item uint, search_name string) ([]items.ItemCore, bool, error) {
	result, count, err := service.itemData.ReadAll(page, item, search_name)

	next := true
	var pages int64
	if item != 0 {
		pages = count / int64(item)
		if count%int64(item) != 0 {
			pages += 1
		}
		if page == uint(pages) {
			next = false
		}
	}

	return result, next, err
}

func (service *ItemService) Delete(id uint) error {
	return service.itemData.Delete(id)

}

func (service *ItemService) GetById(id uint) (items.ItemCore, error) {
	return service.itemData.SelectById(id)
}

func (service *ItemService) Create(input items.ItemCore, file multipart.File, filename string) error {
	return service.itemData.Insert(input, file, filename)
}

func (service *ItemService) Update(id uint, input items.ItemCore, file multipart.File, filename string) error {
	return service.itemData.UpdateDataItem(id, input, file, filename)
}
