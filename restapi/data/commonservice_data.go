package data

type Administrative struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Slug string `json:"slug"`
	NameWithType string `json:"name_with_type"`
	Path string `json:"path"`
	PathWithType string `json:"path_with_type"`
	Code string `json:"code"`
	ParentCode string `json:"parent_code"`
}