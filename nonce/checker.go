package nonce

type Checker interface {
	Check(key string) bool
}
