package _const

type CredentialType int

func (c *CredentialType) Display() string {
	switch c {
	default:
		return "其他"
	}
}

type CredentialStatus string

const (
	DEPRECATED CredentialStatus = "1"
	NORMAL     CredentialStatus = "2"
)

func (receiver CredentialStatus) Display() string {
	switch receiver {
	case DEPRECATED:
		return "已作废"
	default:
		return ""
	}
}
