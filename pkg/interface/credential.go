package _interface

import (
	"github.com/gin-gonic/gin"
)

//type Credential interface {
//	TransferToCredential(template mongo_model.CredentialTemplate, perFunc func(credential Credential) Credential) mongo_model.Credential
//}

type Credential interface {
	ToCredential(ctx *gin.Context)
}
