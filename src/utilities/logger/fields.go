package logger

type Facet interface {
	AddField(string, interface{}) Facet
	GetFields() map[string]interface{}
}

type Field struct {
	Key   string
	Value interface{}
}

func NewFacets(fields ...Field) Facet {
	facet := &facet{
		fields: make(map[string]interface{}, 0),
	}

	for _, field := range fields {
		facet.fields[field.Key] = field.Value
	}

	return facet
}
func (f facet) AddField(key string, value interface{}) Facet {
	f.fields[key] = value
	return f
}

func (f facet) GetFields() map[string]interface{} {
	return f.fields
}

type facet struct {
	fields map[string]interface{}
}
