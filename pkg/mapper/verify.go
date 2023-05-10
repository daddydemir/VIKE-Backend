package mapper

import (
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
)

type VerifyMapper struct {
	Code string
}

func (v VerifyMapper) ToVerify(userId uuid.UUID) models.Verify {
	var verify = models.Verify{
		VerifyId:   uuid.New(),
		CustomerId: userId,
		Code:       v.Code,
	}
	return verify
}
