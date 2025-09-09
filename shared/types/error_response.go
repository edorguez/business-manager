package types

type Error struct {
	Status int64  `json:"status"`
	Error  string `json:"error"`
}
