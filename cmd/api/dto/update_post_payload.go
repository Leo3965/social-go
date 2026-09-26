package dto

type UpdatePostPayload struct {
	Content string   `json:"content" validate:"omitempty,max=100"`
	Title   string   `json:"title" validate:"omitempty,max=100"`
	Tags    []string `json:"tags"`
}
