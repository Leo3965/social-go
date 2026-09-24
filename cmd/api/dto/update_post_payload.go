package dto

type UpdatePostPayload struct {
	Content string   `json:"content" validate:"omitempty,max=100"`
	Tags    []string `json:"tags"`
}
