package _const

type CredentialType int

func (c *CredentialType) Display() string {
	switch c {
	default:
		return "其他"
	}
}
