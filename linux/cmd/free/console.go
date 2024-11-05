package free

type MemoryInfo struct {
	Total     string `json:"total"`
	Used      string `json:"used"`
	Free      string `json:"free"`
	Shared    string `json:"shared"`
	Buffer    string `json:"buffer"`
	Available string `json:"available"`
}
