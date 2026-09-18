package idgen

import (
	"fmt"
	"strconv"
	"strings"

	snowflake "github.com/bwmarrin/snowflake"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Generator struct {
	node *snowflake.Node
}

func NewGenerator(maelNode *maelstrom.Node) (*Generator, error) {
	stringFreeId := strings.TrimPrefix(maelNode.ID(), "n")
	nodeId, err := strconv.ParseInt(stringFreeId, 10, 64)

	if err != nil {
		fmt.Println("error parsing Node ID: ", err)
		return nil, err
	}

	snowflakeNode, err := snowflake.NewNode(nodeId)

	if err != nil {
		fmt.Println("error creating a new Snowflake Node: ", err)
		return nil, err
	}

	return &Generator{node: snowflakeNode}, nil
}

func (g *Generator) GetNextId() string {
	return g.node.Generate().String()
}
