package toncenter

import (
	"fmt"
	"strconv"
)

type Blocker interface {
	MasterchainInfo() (MasterChainInfoStruct, error)
	Blocks(req BlocksStructParams) (BlocksStruct, error)
	MasterchainBlockShardState(seqno int) (MasterChainBlockShardStateStruct, error)
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
func (c *Block) MasterchainBlockShardState(seqno int) (MasterChainBlockShardStateStruct, error) {
	result := MasterChainBlockShardStateStruct{}
	_, err := c.client.resty.R().SetQueryParam("seqno", strconv.Itoa(seqno)).
		SetResult(&result).ForceContentType("application/json").Get(MasterChainBlockShardState + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get MasterChainBlockShardState method error: %v", err)
	}
	return result, err
}

func (c *Block) Blocks(req BlocksStructParams) (BlocksStruct, error) {
	result := BlocksStruct{}
	params := SetParams(req)
	_, err := c.client.resty.R().SetQueryParamsFromValues(params).
		SetResult(&result).ForceContentType("application/json").Get(BlocksUrl + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get Blocks method error: %v", err)
	}
	return result, err
}

func (c *Block) MasterchainInfo() (MasterChainInfoStruct, error) {
	result := MasterChainInfoStruct{}
	_, err := c.client.resty.R().
		SetResult(&result).ForceContentType("application/json").Get(MasterChainUrl + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get MasterChainInfo method error: %v", err)
	}
	return result, err
}
