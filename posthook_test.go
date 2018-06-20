package posthook

import (
	"fmt"
	"testing"
)

func TestDeserializeSingle(t *testing.T) {

	resp := `{
		"data": {
		  "id": "0525baf7-79f9-4991-bb39-6a601ead00e5",
		  "path": "/webhooks/ph/event_reminder",
		  "data": {
			"eventID": 25
		  },
		  "postAt": "2018-06-15T21:00:00Z",
		  "status": "PENDING",
		  "createdAt": "2018-04-09T20:43:24.321778Z",
		  "updatedAt": "2018-04-09T20:43:24.321778Z"
		}
	  }`

	hook, err := single([]byte(resp))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hook.ID)

	if hook.ID == "" {
		t.Fatal("empty id")
	}
	if hook.ID != "0525baf7-79f9-4991-bb39-6a601ead00e5" {
		t.Fatal("ID mistmatch")
	}
	if hook.Data == nil {
		t.Fatal("empty data")
	}
}
func TestDeserializeList(t *testing.T) {
	resp := `{
		"data": [
		  {
			"id": "8e3aa909-fb84-4495-944d-a4c192defe66",
			"path": "/webhooks/ph/event_reminder",
			"domain": "posthook.io",
			"data": {
			  "eventID": 25
			},
			"postAt": "2018-07-19T21:00:00Z",
			"status": "PENDING",
			"createdAt": "2018-04-30T17:45:20.68114Z",
			"updatedAt": "2018-04-30T17:45:20.68114Z"
		  },
		  {
			"id": "c1ec9560-65fc-4b88-bfe0-1bc6e56cb3db",
			"path": "/webhooks/ph/event_reminder",
			"domain": "posthook.io",
			"data": {
			  "eventID": 25
			},
			"postAt": "2018-07-17T21:00:00Z",
			"status": "PENDING",
			"createdAt": "2018-04-30T17:45:16.097812Z",
			"updatedAt": "2018-04-30T17:45:16.097812Z"
		  },
		  {
			"id": "6ab4d4d5-4767-452d-b72f-6ec40562b27e",
			"path": "/webhooks/ph/event_reminder",
			"domain": "posthook.io",
			"data": {
			  "eventID": 25
			},
			"postAt": "2018-07-15T21:00:00Z",
			"status": "PENDING",
			"createdAt": "2018-04-30T17:45:06.475047Z",
			"updatedAt": "2018-04-30T17:45:06.475047Z"
		  }
		]
	  }
	`
	hooks, err := list([]byte(resp))
	if err != nil {
		t.Fatal(err)
	}
	if hooks == nil {
		t.Fatal("Empty list of hooks")
	}
	want := 3
	got := len(hooks)

	if got != want {
		t.Fatalf("got %d want: %d", got, want)
	}
}
