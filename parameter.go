package nibe

// Parameter -
type Parameter struct {
  ParameterID uint32
  Name string
  Title string
  Designation string
  unit string
  DisplayValue string
  RawValue int8
}

// GetParameters -
func (api *API) GetParameters(args...string) []Parameter {
  return []Parameter{}
}
