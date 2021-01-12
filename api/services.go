package api

//Services comment generic
type Services struct {
	search MovieSearch
}

//WebServices comment generic
type WebServices struct {
	s Services
}

//NewServices comment generic
func NewServices() Services {
	return Services{
		search: &MovieService{},
	}
}

func start() *WebServices {
	return &WebServices{s: NewServices()}
}
