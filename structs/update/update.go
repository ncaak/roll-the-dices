package update


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

type Update struct {
	Ok bool			`json:"ok"`
	Result []Result	`json:"result"`
}

func (result *Result) IsCommand() bool {
	var command = false
	if ent := result.Message.Entities; len(ent) > 0 && ent[0].Type == "bot_command"{
		command = true
	}
	return command
}

func (result *Result) GetCommand() string {
	return result.Message.Text;
}

func (result *Result) GetChatId() int {
	return result.Message.Chat.Id;
}

func (result *Result) GetReplyId() int {
	return result.Message.MessageId;
}
