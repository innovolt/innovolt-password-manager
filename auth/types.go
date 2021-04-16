package auth

type UserCredential struct {
	Username string
	Password string
}

type AppCredential struct {
	ApiKey string
}

type Credential struct {
	User interface{}
	App  interface{}
}
