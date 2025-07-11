package method

type Method int

const (
	Get Method = iota
	Post
	Put
	Patch
	Delete
	Options
)

var stringValues = map[Method]string{
	Get:     "GET",
	Post:    "POST",
	Put:     "PUT",
	Patch:   "PATCH",
	Delete:  "DELETE",
	Options: "OPTIONS",
}

func (method Method) ToString() string {
	return stringValues[method]
}
