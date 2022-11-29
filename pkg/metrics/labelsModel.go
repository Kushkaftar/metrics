package metrics

type CreateLabel struct {
	Label Label `json:"label"`
}

type Labels struct {
	Labels []Label `json:"labels"`
}

type Label struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
