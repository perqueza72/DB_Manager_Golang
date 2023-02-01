package filehandlers

import (
	"fmt"
	"io/fs"
	models "models_zinc"
	"path/filepath"
)

func PathFile2Email(path string) (*models.Email, error) {
	lines, errGetLines := GetLines(path)
	if errGetLines != nil {
		return nil, errGetLines
	}

	email, errParsing := ParseStrings2Email(lines)
	if errParsing != nil {
		return nil, errParsing
	}

	return email, nil
}

func FolderFiles2Email(path string) ([]*models.Email, error) {

	emails := []*models.Email{}

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Error go thoought folders, %v", err)
		}

		if !d.IsDir() {
			email, err := PathFile2Email(path)
			if err != nil {
				return err
			}
			emails = append(emails, email)
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return emails, nil
}
