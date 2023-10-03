package helpers

import gonanoid "github.com/matoous/go-nanoid"

func GenerateName() (string, error) {
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	newid, errID := gonanoid.Generate(str, 30)
	if errID != nil {
		return "", errID
	}
	return newid, nil
}
