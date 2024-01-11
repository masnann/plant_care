package email

import (
	"fmt"
	"math/rand"
	"time"
)

type VerificationCodeGenerator interface {
	GenerateCode(length int) string
}

type DefaultCodeGenerator struct{}

func NewCodeGenerator() VerificationCodeGenerator {
	return DefaultCodeGenerator{}
}

func (dcg DefaultCodeGenerator) GenerateCode(length int) string {
	rand.Seed(time.Now().UnixNano())

	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}

	return code
}
