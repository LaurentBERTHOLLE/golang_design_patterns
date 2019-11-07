package builder

func GetBuilder() *builder {
	return &builder{}
}

func (b *builder) Field1(f string) *builder {
	b.dummy.Field1 = f
	return b
}

func (b *builder) Field2(f string) *builder {
	b.dummy.Field2 = f
	return b
}

func (b *builder) Field3(f int) *builder {
	b.dummy.Field3 = f
	return b
}

func (b *builder) Build() Dummy {
	return b.dummy
}

type builder struct {
	dummy Dummy
}
