package virtual_constructor

import "strings"

var exemplars = make(map[string]Exemplar)

type Exemplar interface {
	Tag() string
	Make([]string) Exemplar
}

func RegisterExemplar(exemplar Exemplar) {
	exemplars[exemplar.Tag()] = exemplar
}

func Create(input string) (Exemplar, bool) {
	parts := strings.Split(input, " ")
	if exemplar, present := exemplars[parts[0]]; present {
		return exemplar.Make(parts[1:]), true
	} else {
		return nil, false
	}
}
