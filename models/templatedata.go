package models

// TemplateData hold the data sent by handlers
// This is the way to avoid import cycle problem in Go
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRSFtoken string
	Flash     string
	Warning   string
	Error     string
}
