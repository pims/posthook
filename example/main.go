package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pims/posthook"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// create posthook client
	ph := posthook.New(os.Getenv("POSTHOOK_API_KEY"))

	// verify endpoint being used
	fmt.Println(ph.Endpoint())

	// data to be sent back to us when posthook executes the hook
	data := map[string]interface{}{
		"foobar": "baz",
	}

	// Step 1: schedule a hook
	hook, err := ph.Schedule("ping", time.Now().Add(1*time.Minute).UTC(), data)
	handleErr(err)

	fmt.Printf("Hook id: %s, status: %s\n", hook.ID, hook.Status)

	filters := posthook.Filters{
		Limit: posthook.Int(10),
	}

	// Step 2: verify our hook has been scheduled
	hooks, err := ph.List(filters)
	handleErr(err)

	for _, hook := range hooks {
		in := hook.PostAt.Sub(time.Now().UTC())
		fmt.Println(hook.ID, hook.Status, hook.PostAt, in.Minutes(), hook.Path, hook.Data)
	}

	if len(hooks) == 0 {
		handleErr(fmt.Errorf("expected at least 1 hook"))
	}
	// Step 3: double check the info about the hook we just scheduled
	hook, err = ph.Get(hooks[0].ID)
	handleErr(err)
	fmt.Printf("Hook id: %s, path: %s, scheduled for: %s\n", hook.ID, hook.Path, hook.PostAt)

	// (optional) Step 4: decide we no longer want this hook scheduled
	err = ph.Delete(hooks[0].ID)
	handleErr(err)
	fmt.Printf("Hook %s was deleted", hooks[0].ID)
}
