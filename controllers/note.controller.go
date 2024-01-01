package controllers

import (
	"crud/dto"
	"crud/entity"
	"crud/models"
	"crud/validator"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateNoteHandler(c *fiber.Ctx) error {
	var payload dto.CreateNoteSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	newNote := entity.Note{
		ID:        uuid.New().String(),
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		Published: payload.Published,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := models.CreateNote(&newNote)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "fail",
				"message": "Title already exist, please use another title",
			})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"status":  "error",
				"message": result.Error.Error(),
			})
		}
	} else if result.RowsAffected < 1 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newNote,
	})
}

func GetNotes(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var notes []entity.Note
	result := models.GetAllNotes(&notes, intLimit, offset)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   notes,
	})
}

func FindNoteById(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	var note entity.Note

	result := models.GetNoteById(&note, noteId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "Note not found",
			})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	var payload dto.UpdateNoteSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	var note entity.Note
	if result := models.GetNoteById(&note, noteId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "fail",
				"message": "Note not found",
			})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error.Error(),
		})
	}

	if payload.Title != "" {
		note.Title = payload.Title
	}
	if payload.Category != "" {
		note.Category = payload.Category
	}
	if payload.Content != "" {
		note.Content = payload.Content
	}
	if payload.Published {
		note.Published = payload.Published
	}

	note.UpdatedAt = time.Now()

	if result := models.UpdateNote(&note, noteId); result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "fail",
				"message": "Title already exist, please use another title",
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "fail",
				"message": result.Error.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   note,
	})
}

func DeleteNote(c *fiber.Ctx) error {
	noteId := c.Params("noteId")
	result := models.DeleteNote(&entity.Note{}, noteId)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "fail",
			"message": result.Error,
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "Note not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully deleted note",
	})
}
