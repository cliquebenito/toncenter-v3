package toncenter

import (
	"fmt"
)

type Blocker interface {
	MasterchainInfo() (MasterChainInfoStruct, error)
	// Blocks
	// MasterchainBlockShardState
	// AddressBook
	// MastherChainBlockShards
}
type Block struct {
	client Client
}

func NewBlock(c *Client) Blocker {
	return &Block{
		client: *c,
	}
}

//func (c *Block) Blocks() (BlocksStruct, error) {
//	result := BlocksStruct{}
//	_, err := c.client.resty.R().
//		SetResult(&result).ForceContentType("application/json").Get(c.client.url)
//	if err != nil {
//		return result, fmt.Errorf("Get MasterChainInfo method error: %v", err)
//	}
//	return result, err
//}

func (c *Block) MasterchainInfo() (MasterChainInfoStruct, error) {
	result := MasterChainInfoStruct{}
	_, err := c.client.resty.R().
		SetResult(&result).ForceContentType("application/json").Get(MasterChainUrl + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get MasterChainInfo method error: %v", err)
	}
	return result, err
}
