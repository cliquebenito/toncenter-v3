package toncenter

import (
	"fmt"
	"net/url"
	"strconv"
)

type Blocker interface {
	MasterchainInfo() (MasterChainInfoStruct, error)
	Blocks(req BlocksStructParams) (BlocksStruct, error)
	MasterchainBlockShardState(seqno int) (MasterChainBlockShardStateStruct, error)
	AddressBook(query AddressBookParams) (AddressBookResponse, error)
	MasterchainBlocksShards(seqno int) (MasterChainBlockShardsStruct, error)
}
type Block struct {
	client Client
}

func NewBlock(c *Client) Blocker {
	return &Block{
		client: *c,
	}
}
func (c *Block) MasterchainBlocksShards(seqno int) (MasterChainBlockShardsStruct, error) {
	result := MasterChainBlockShardsStruct{}
	_, err := c.client.resty.R().SetQueryParam("seqno", strconv.Itoa(seqno)).
		SetResult(&result).ForceContentType("application/json").Get(MasterChainBlockShards + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get MasterchainBlocksShards method error: %v", err)
	}
	return result, err

}

func (c *Block) AddressBook(query AddressBookParams) (AddressBookResponse, error) {
	result := AddressBookResponse{}
	params := make([]string, 0, len(query.Address))
	for _, v := range query.Address {
		params = append(params, v)
	}
	queryParams := url.Values{}

	for _, v := range params {
		queryParams.Add("address", v)
	}
	_, err := c.client.resty.R().SetQueryParamsFromValues(queryParams).
		SetResult(&result).ForceContentType("application/json").Get(AddressBookURL + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get AddressBook method error: %v", err)
	}
	return result, err
}

func (c *Block) MasterchainBlockShardState(seqno int) (MasterChainBlockShardStateStruct, error) {
	result := MasterChainBlockShardStateStruct{}
	_, err := c.client.resty.R().SetQueryParam("seqno", strconv.Itoa(seqno)).
		SetResult(&result).ForceContentType("application/json").Get(MasterChainBlockShardState + c.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get MasterchainBlockShardState method error: %v", err)
	}
	return result, err
}

func (c *Block) Blocks(req BlocksStructParams) (BlocksStruct, error) {
	result := BlocksStruct{}
	params := SetBlockParams(req)
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
		return result, fmt.Errorf("Get MasterchainInfo method error: %v", err)
	}
	return result, err
}
