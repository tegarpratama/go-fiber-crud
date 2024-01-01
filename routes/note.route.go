package routes

import (
	"crud/controllers"

	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(api fiber.Router) {
	router := api.Group("/notes")

	router.Get("/", controllers.GetNotes)
	router.Post("/", controllers.CreateNoteHandler)
	router.Get("/:noteId/detail", controllers.FindNoteById)
	router.Put("/:noteId/update", controllers.UpdateNote)
	router.Delete("/:noteId/delete", controllers.DeleteNote)
}
