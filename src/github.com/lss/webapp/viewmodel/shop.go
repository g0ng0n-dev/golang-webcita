package viewmodel

type Shop struct {
	Title string
	Active string
	Categories []Category
}

type Category struct {
	URL string
	ImageUrl string
	Title string
	Description string
}
