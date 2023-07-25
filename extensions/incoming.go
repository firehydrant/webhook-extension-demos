package extensions

type Payload struct {
	Data  Data  `json:"data"`
	Event Event `json:"event"`
}

type Data struct {
	Callback         Callback         `json:"callback"`
	CommandArguments []string         `json:"command_arguments"`
	CommandExtension CommandExtension `json:"command_extension"`
	FirehydrantUser  FirehydrantUser  `json:"firehydrant_user"`
	Payload          interface{}      `json:"payload"`
	PayloadType      string           `json:"payload_type"`
}

type Callback struct {
	Expiration string `json:"expiration"`
	URL        string `json:"url"`
}

type CommandExtension struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type FirehydrantUser struct {
	Email  string `json:"email"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type Event struct {
	Operation    string `json:"operation"`
	ResourceType string `json:"resource_type"`
}
