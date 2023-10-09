package lecture6

type BookGenre struct {
	GenreName string `json:"genre_name"`
	Subgenre  string `json:"subgenre"`
}

type Book struct {
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Genre  BookGenre `json:"genre"`
}
