package tickets_test

import "testing"

func TestRespond(t *testing.T) {
	// Setup
	_, ticketsController := makeController()

	err := ticketsController.Respond("", "", "Hello bitch")

	if err != nil {
		t.Error(err)
	}
}
