package vo

type EntriesRelationship struct {
	Links struct {
		Self    string `json:"self"`
		Related string `json:"related"`
	} `json:"links"`
}

type Relationship struct {
	Links struct {
		Self    string `json:"self"`
		Related string `json:"related"`
	} `json:"links"`
	Data struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}
