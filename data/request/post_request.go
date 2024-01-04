package request

type PostRequest struct {
	Title   string   `validate:"required,min=1,max=255" json:"title"`
	Content string   `validate:"required" json:"content"`
	Tags    []string `json:"tags"`
}
