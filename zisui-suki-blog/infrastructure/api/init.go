package api

type API struct {
	AuthApiDomain string
}

func Init() (*API, error) {
	return &API{
		AuthApiDomain: "http://localhost:3001/",
	}, nil
}
