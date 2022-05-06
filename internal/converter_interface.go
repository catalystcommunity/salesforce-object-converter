package internal

type Converter interface {
	SetObject(object string)
	SetRawFieldMap(rawFieldMap map[string]string)
	Convert()
}
