package extensions

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
)

type Action interface {
	_action()
}

type Response struct {
	ReplyType string   `json:"reply_type"`
	Actions   []Action `json:"actions"`
}

type NewNoteAction struct {
	Action string `json:"action"`
	Body   string `json:"body"`
}

func (nna *NewNoteAction) _action() {}

type ChangeEventAction struct {
	Action      string      `json:"action"`
	ChangeEvent ChangeEvent `json:"change_event"`
}

type ReplyToChatAction struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ChangeEvent struct {
	Summary string `json:"summary"`
}

func ReplyForRestart(old, new *appsv1.Deployment) (*Response, error) {
	newNoteAction := &NewNoteAction{
		Action: "new_note",
		Body:   fmt.Sprintf("Restarting deployment '%s' on namespace '%s' - current resource version is '%s'", old.Name, old.Namespace, old.ResourceVersion),
	}

	response := &Response{
		ReplyType: "actions",
		Actions:   []Action{newNoteAction},
	}

	return response, nil
}
