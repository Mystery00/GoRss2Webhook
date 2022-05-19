package webhook

type Webhook struct {
	SubscribeUrl string
	Http         Http
}

type Http struct {
	Url    string
	Method string
	Body   string
	Header map[string]string
	Auth   Auth
}

type Auth struct {
	Username string
	Password string
}
