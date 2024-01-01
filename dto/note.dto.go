package dto

type CreateNoteSchema struct {
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Category  string `json:"category" validate:"required"`
	Published bool   `json:"published"`
}

type UpdateNoteSchema struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Published bool   `json:"published"`
}
