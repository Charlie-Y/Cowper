package bibleweb

type BookRef struct {
	Path string `json:"path"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

type BookRefWrap struct {
	Book BookRef `json:"book"`
}
