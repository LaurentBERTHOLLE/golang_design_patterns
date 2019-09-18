package virtual_constructor

func init() {
	RegisterExemplar(Dummy{})
}

type Dummy struct {
	Field string
}

func (d Dummy) Tag() string {
	return "dummy"
}

func (d Dummy) Make(input []string) Exemplar {
	return Dummy{Field: input[0]}
}
