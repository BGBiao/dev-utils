package uuid

import (
	"github.com/google/uuid"
)

func NewUUID() (uint32, string) {

	// a random uuid with the crypto/rand package
	uid := uuid.New()
	return uid.ID(), uid.String()

}
