package toncenter

import (
	"fmt"
)

type Blocks interface {
	MasterchainInfo() (MasterChainInfoStruct, error)
	// Blocks
	// MasterchainBlockShardState
	// AddressBook
	// MastherChainBlockShards
}
type Block struct {
	client Client
}

func NewBlock(c Client) Blocks {
	return &Block{
		client: c,
	}
}

func (c *Block) MasterchainInfo() (MasterChainInfoStruct, error) {
	result := MasterChainInfoStruct{}
	_, err := c.client.resty.R().
		SetResult(&result).ForceContentType("application/json").Get(c.client.url)
	if err != nil {
		return result, fmt.Errorf("Get MasterChainInfo method error: %v", err)
	}
	return result, err
}
