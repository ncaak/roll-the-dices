package response


type Chat struct {
	FirstName string	`json:"first_name"`
	Id int				`json:"id"`
	LastName string		`json:"last_name"`
	Type string			`json:"type"`
	Username string		`json:"username"`
}

type Entities struct {
	Length int	`json:"length"`
	Offset int	`json:"offset"`
	Type string	`json:"type"`
}

type From struct {
    FirstName string    `json:"first_name"`
	Id int              `json:"id"`
	IsBot bool			`json:"is_bot"`
	LanguageCode string	`json:"language_code"`
	LastName string     `json:"last_name"`
	Username string     `json:"username"`
}

type Message struct {
	Chat Chat			`json:"chat"`
	Date int			`json:"date"`
	Entities []Entities	`json:"entities"`
	From From			`json:"from"`
	MessageId int		`json:"message_id"`
	Text string			`json:"text"`
}

type Result struct {
	Message Message	`json:"message"`
	UpdateId int	`json:"update_id"`
}

type Response struct {
	Ok bool			`json:"ok"`
	Result []Result	`json:"result"`
}

