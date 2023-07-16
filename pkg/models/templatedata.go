package models

// Template holds data sent from handlers to templates

type TemplateData struct {
	// what kind of things we are sending
	// string information
	// we need a generic datatype

	StringMap map[string]string

	IntMap   map[string]int
	FloatMap map[string]float32

	// objects

	Data map[string]interface{} // we don't know what the data type is so interface

	// when we get into the forms part we have the Cross Site Forgery Token

	CSRFToken string

	// success and error messages

	Flash string

	Warning string
	Error   string
}
