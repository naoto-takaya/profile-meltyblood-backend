package service

import (
	"fmt"

	"github.com/naoto-takaya/profile-meltyblood-backend/validators"
)

func Upload(image validators.Image) {
	fmt.Print(image)
}
