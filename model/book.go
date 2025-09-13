package model

type Book struct {
	Id      string `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Editora string `json:"editora"`
	ISBN    string `json:"isbn"`
}
