package service

import (
	"Hannon-app/features/items"
	"mime/multipart"
)

type ItemService struct {
	itemData items.ItemDataInterface
}

// UnarchiveItem implements items.ItemServiceInterface.
func (service *ItemService) UnarchiveItem(tenantID uint, id uint, input items.ItemCore) error {
	err := service.itemData.UnarchiveItem(tenantID, id, input)
	return err
}

// ArchiveItem implements items.ItemServiceInterface.
func (service *ItemService) ArchiveItem(tenantID uint, id uint, input items.ItemCore) error {
	err := service.itemData.ArchiveItem(tenantID, id, input)
	return err
}

// GetArchiveItem implements items.ItemServiceInterface.
func (service *ItemService) GetArchiveItem(tenantID uint) ([]items.ItemCore, error) {
	result, err := service.itemData.ReadArchiveItem(tenantID)
	return result, err
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

func (service *ItemService) Delete(tenantID uint, id uint) error {
	return service.itemData.Delete(tenantID, id)

}

func (service *ItemService) GetById(id uint) (items.ItemCore, error) {
	return service.itemData.SelectById(id)
}

func (service *ItemService) Create(tenantID uint, input items.ItemCore, file multipart.File, filename string) error {
	return service.itemData.Insert(tenantID, input, file, filename)
}

func (service *ItemService) Update(tenantID uint, id uint, input items.ItemCore, file multipart.File, filename string) error {
	return service.itemData.UpdateDataItem(tenantID, id, input, file, filename)
}

func (service *ItemService) GetItemsByTenant(tenantID uint, page, item uint, searchName string) ([]items.ItemCore, bool, error) {

	result, count, err := service.itemData.ReadItemsByTenant(tenantID, page, item, searchName)

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
