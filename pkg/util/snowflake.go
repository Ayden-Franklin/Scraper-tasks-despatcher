package util

import "github.com/bwmarrin/snowflake"

func Generate() int64{
	node, err := snowflake.NewNode(1)
	if err != nil {
		return -1
	}
	id := node.Generate()
	return id.Int64()
}
