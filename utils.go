package toncenter

import (
	"net/url"
	"strconv"
)

func SetBlockParams(req BlocksStructParams) url.Values {
	queryParams := &url.Values{}
	params := map[string]string{
		"workchain":   req.WorkChain,
		"shard":       req.Shard,
		"seqno":       req.Seqno,
		"start_utime": req.StartUtime,
		"end_utime":   req.EndUtime,
		"start_lt":    req.StartLt,
		"end_lt":      req.EndLt,
		"limit":       req.Limit,
		"offset":      req.Offset,
		"sort":        req.Sort,
	}
	for key, value := range params {
		if value != "" {
			queryParams.Set(key, value)
		}
	}
	return *queryParams
}

func SetParamsTxMasterChain(req TransactionByMasterchainBlockParams) url.Values {
	queryParams := &url.Values{}
	params := map[string]string{
		"seqno":  strconv.Itoa(req.Seqno),
		"limit":  strconv.Itoa(req.Limit),
		"offset": strconv.Itoa(req.Offset),
		"sort":   req.Sort,
	}
	for key, value := range params {
		if value != "" {
			queryParams.Set(key, value)
		}
	}
	return *queryParams
}
func SetParamsTx(req TransactionsParams) url.Values {
	queryParams := &url.Values{}
	var account, exclude_account string
	for _, v := range req.Account {
		account += v
	}
	for _, v := range req.ExcludeAccount {
		exclude_account += v
	}
	params := map[string]string{
		"workchain":       req.WorkChain,
		"shard":           req.Shard,
		"seqno":           req.Seqno,
		"account":         account,
		"exclude_account": exclude_account,
		"hash":            req.Hash,
		"lt":              req.Lt,
		"start_utime":     req.StartUtime,
		"end_utime":       req.EndUtime,
		"start_lt":        req.StartLt,
		"end_lt":          req.EndLt,
		"limit":           req.Limit,
		"offset":          req.Offset,
		"sort":            req.Sort,
	}
	for key, value := range params {
		if value != "" {
			queryParams.Set(key, value)
		}
	}
	return *queryParams
}
