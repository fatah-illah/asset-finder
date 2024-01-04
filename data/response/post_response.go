package response

type PostResponse struct {
	ID      uint          `json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Tags    []TagResponse `json:"tags"`
}
