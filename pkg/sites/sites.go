package sites

type Annonces interface {
	GetAnnonces() string
}

func Get(a Annonces) string {
	return a.GetAnnonces()
}
