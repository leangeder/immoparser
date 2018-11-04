package leboncoin

import "fmt"

type Annonce struct{}

func (a Annonce) GetAnnonces() string {
	return fmt.Sprintf("Get Annonces from leboncoin")
}
