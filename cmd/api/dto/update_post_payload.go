package dto

type UpdatePostPayload struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
