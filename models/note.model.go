package models

import (
	"crud/config"
	"crud/entity"

	"gorm.io/gorm"
)

func CreateNote(note *entity.Note) *gorm.DB {
	return config.DB.Create(&note)
}

func UpdateNote(note *entity.Note, noteId string) *gorm.DB {
	return config.DB.Where("id = ?", noteId).Updates(&note)
}

func GetAllNotes(note *[]entity.Note, limit int, offset int) *gorm.DB {
	return config.DB.Limit(limit).Offset(offset).Find(&note)
}

func GetNoteById(note *entity.Note, noteId string) *gorm.DB {
	return config.DB.First(&note, "id = ?", noteId)
}

func DeleteNote(note *entity.Note, noteId string) *gorm.DB {
	return config.DB.Where("id = ?", noteId).Delete(&note)
}
