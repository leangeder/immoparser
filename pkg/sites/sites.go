package sites

type Annonces interface {
	GetAnnonces() string
}

// func Get(a Annonces) string {
// 	//  listAnnonceleboncoin := leboncoin.Annonce{}
// 	// listAnnoncedomimmo := domimmo.Annonce{}
//
// 	// log.Println(sites.Get(listAnnonceleboncoin))
// 	// log.Println(sites.Get(listAnnoncedomimmo))
// 	return a.GetAnnonces()
// }

type site struct{}

func GetAnnoncesFromSites(a Annonces) string {
}

func Get() string {
	return GetAnnoncesFromSites(&site{})
}
