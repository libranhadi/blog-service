package requests

type PostCreateRequest struct {
	Title    string `validate:"required" json:"title"`
	Content  string `validate:"required" json:"content"`
	Category string `validate:"required" json:"category"`
}

type PostUpdateRequest struct {
	Id       string `validate:"required" json:"id"`
	Title    string `validate:"required" json:"title"`
	Content  string `validate:"required" json:"content"`
	Category string `validate:"required" json:"category"`
}
