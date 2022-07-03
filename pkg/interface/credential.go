package _interface

import (
	"github.com/XC-Zero/yinwan/pkg/model/mongo_model"
)

type Credential interface {
	TransferToCredential(template mongo_model.CredentialTemplate, perFunc func(credential Credential) Credential) mongo_model.Credential
}
