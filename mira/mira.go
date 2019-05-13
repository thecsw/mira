package mira

// When we initialize the Reddit instance,
// automatically start a goroutine that will
// update the token every 45 minutes. The
// auto_refresh should not be accessible to
// the end user as it is an internal method
func Init(c Credentials) *Reddit {
	auth, _ := Authenticate(&c)
	go auth.auto_refresh()
	return auth
}
