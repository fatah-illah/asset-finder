package request

type TagRequest struct {
	Label string   `validate:"required,min=1,max=255" json:"label"`
	Posts []string `json:"posts"`
}
