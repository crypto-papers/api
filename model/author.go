package model

type Author struct {
	Name      string `json:"name"`
	Psuedonym bool   `json:"psuedonym"`
	CreatedAt string `json:"created"`
}
