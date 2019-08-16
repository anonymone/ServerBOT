package modules

import "testing"

func TestConfig(t *testing.T) {
	// var a JsonOperator
	path := "./config.json"
	data := ConfigInformation{
		Receiver:   []string{"1"},
		Sender:     "1",
		ServerHost: "1",
		ServerPort: 1,
		APICredit:  "1",
		TaskSend:   "1",
		TaskCheck:  "1",
		TaskTry:    "1"}
	data.WriteJSON(path)
	data.ReadJSON(path)
}
