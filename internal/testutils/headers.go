package testutils

import "fmt"

const (
	CookieHeader    = "Cookie"
	SetCookieHeader = "Set-Cookie"
	UidHeader       = "uid"
	UidHeaderValue  = "123"
)

func UidCookieHeaderResponse() string {
	return fmt.Sprintf("%s=%s; path=/", UidHeader, UidHeaderValue)
}

func UidCookieHeaderRequest() string {
	return fmt.Sprintf("%s=%s", UidHeader, UidHeaderValue)
}
