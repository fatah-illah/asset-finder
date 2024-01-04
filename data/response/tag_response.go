package response

type TagResponse struct {
	ID    uint           `json:"id"`
	Label string         `json:"label"`
	Posts []PostResponse `json:"posts"`
}
