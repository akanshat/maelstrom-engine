package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"github.com/akanshat/maelstrom-engine/pkg/idgen"
)

func main() {
	n := maelstrom.NewNode()

	var idGenerator *idgen.Generator
	var initOnce sync.Once
	var initErr error

	n.Handle("echo", func(msg maelstrom.Message) error {
		var body map[string]any

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		return n.Reply(msg, body)
	})

	n.Handle("generate", func(msg maelstrom.Message) error {
		initOnce.Do(func() {
			idGenerator, initErr = idgen.NewGenerator(n)
		})

		if initErr != nil {
			return initErr
		}

		var body map[string]any

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		uniqueId := idGenerator.GetNextId()

		body["type"] = "generate_ok"
		body["id"] = uniqueId

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
