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

func (c *Block) Blocks(req BlocksStructParams) (BlocksStruct, error) {
	result := BlocksStruct{}

	_, err := c.client.resty.R().SetQueryParams(map[string]string{
		req.WorkChain:  req.WorkChain,
		req.Shard:      req.Shard,
		req.Seqno:      req.Seqno,
		req.StartUtime: req.StartUtime,
		req.EndUtime:   req.EndUtime,
		req.StartLt:    req.StartLt,
		req.EndLt:      req.EndLt,
		req.Limit:      req.Limit,
		req.Offset:     req.Offset,
		req.Sort:       req.Sort,
	}).
		SetResult(&result).ForceContentType("application/json").Get(BlocksUrl + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get MasterChainInfo method error: %v", err)
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
