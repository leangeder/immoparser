package domimmo

import "fmt"

type Annonce struct{}

func (a Annonce) GetAnnonces() string {
	return fmt.Sprintf("Get Annonces from domimmo")
}
