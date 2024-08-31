package toncenter

import "net/url"

func SetParams(req BlocksStructParams) url.Values {
	queryParams := url.Values{}
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
	return queryParams
}
