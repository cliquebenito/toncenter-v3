package toncenter

import (
	"fmt"
)

type Transactioner interface {
	Transactions(params TransactionsParams) (TransactionStruct, error)
}
type Transaction struct {
	client Client
}

func NewTransaction(c *Client) Transactioner {
	return &Transaction{
		client: *c,
	}
}

func (t *Transaction) TransactionByMasterchainBlock(params TransactionByMasterchainBlockParams) (TransactionByMasterChainBlockStruct, error) {
	result := TransactionByMasterChainBlockStruct{}
	param := SetParamsTxMasterChain(params)
	_, err := t.client.resty.R().SetQueryParamsFromValues(param).
		SetResult(&result).ForceContentType("application/json").Get(TransactionByMasterchainBlock + t.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get TransactionByMasterchainBlock method error: %v", err)
	}
	return result, err
}

func (t *Transaction) Transactions(params TransactionsParams) (TransactionStruct, error) {
	result := TransactionStruct{}
	parsmeters := SetParamsTx(params)
	_, err := t.client.resty.R().SetQueryParamsFromValues(parsmeters).
		SetResult(&result).ForceContentType("application/json").Get(Transactions + t.client.apikey)
	if err != nil {
		return result, fmt.Errorf("Get Transactions method error: %v", err)
	}
	return result, err

}
