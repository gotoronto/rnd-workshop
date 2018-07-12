package urlset

var List = NewURLSet("")

func Add(url string) (bool, error) {
	return List.Add(url)
}
