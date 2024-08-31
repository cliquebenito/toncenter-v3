package toncenter

type MasterChainInfoStruct struct {
	Last struct {
		Workchain              int    `json:"workchain"`
		Shard                  string `json:"shard"`
		Seqno                  int    `json:"seqno"`
		RootHash               string `json:"root_hash"`
		FileHash               string `json:"file_hash"`
		GlobalID               int    `json:"global_id"`
		Version                int    `json:"version"`
		AfterMerge             bool   `json:"after_merge"`
		BeforeSplit            bool   `json:"before_split"`
		AfterSplit             bool   `json:"after_split"`
		WantMerge              bool   `json:"want_merge"`
		WantSplit              bool   `json:"want_split"`
		KeyBlock               bool   `json:"key_block"`
		VertSeqnoIncr          bool   `json:"vert_seqno_incr"`
		Flags                  int    `json:"flags"`
		GenUtime               string `json:"gen_utime"`
		StartLt                string `json:"start_lt"`
		EndLt                  string `json:"end_lt"`
		ValidatorListHashShort int    `json:"validator_list_hash_short"`
		GenCatchainSeqno       int    `json:"gen_catchain_seqno"`
		MinRefMcSeqno          int    `json:"min_ref_mc_seqno"`
		PrevKeyBlockSeqno      int    `json:"prev_key_block_seqno"`
		VertSeqno              int    `json:"vert_seqno"`
		MasterRefSeqno         int    `json:"master_ref_seqno"`
		RandSeed               string `json:"rand_seed"`
		CreatedBy              string `json:"created_by"`
		TxCount                int    `json:"tx_count"`
		MasterchainBlockRef    struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"masterchain_block_ref"`
		PrevBlocks []struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"prev_blocks"`
	} `json:"last"`
	First struct {
		Workchain              int    `json:"workchain"`
		Shard                  string `json:"shard"`
		Seqno                  int    `json:"seqno"`
		RootHash               string `json:"root_hash"`
		FileHash               string `json:"file_hash"`
		GlobalID               int    `json:"global_id"`
		Version                int    `json:"version"`
		AfterMerge             bool   `json:"after_merge"`
		BeforeSplit            bool   `json:"before_split"`
		AfterSplit             bool   `json:"after_split"`
		WantMerge              bool   `json:"want_merge"`
		WantSplit              bool   `json:"want_split"`
		KeyBlock               bool   `json:"key_block"`
		VertSeqnoIncr          bool   `json:"vert_seqno_incr"`
		Flags                  int    `json:"flags"`
		GenUtime               string `json:"gen_utime"`
		StartLt                string `json:"start_lt"`
		EndLt                  string `json:"end_lt"`
		ValidatorListHashShort int    `json:"validator_list_hash_short"`
		GenCatchainSeqno       int    `json:"gen_catchain_seqno"`
		MinRefMcSeqno          int    `json:"min_ref_mc_seqno"`
		PrevKeyBlockSeqno      int    `json:"prev_key_block_seqno"`
		VertSeqno              int    `json:"vert_seqno"`
		MasterRefSeqno         int    `json:"master_ref_seqno"`
		RandSeed               string `json:"rand_seed"`
		CreatedBy              string `json:"created_by"`
		TxCount                int    `json:"tx_count"`
		MasterchainBlockRef    struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"masterchain_block_ref"`
		PrevBlocks []struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"prev_blocks"`
	} `json:"first"`
}
type BlocksStructParams struct {
	WorkChain  string `json:"workchain"`
	Shard      string `json:"shard"`
	Seqno      string `json:"seqno"`
	StartUtime string `json:"start_utime"`
	EndUtime   string `json:"end_utime"`
	StartLt    string `json:"start_lt"`
	EndLt      string `json:"end_lt"`
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	Sort       string `json:"sort"`
}

type BlocksStruct struct {
	Blocks []struct {
		Workchain              int    `json:"workchain"`
		Shard                  string `json:"shard"`
		Seqno                  int    `json:"seqno"`
		RootHash               string `json:"root_hash"`
		FileHash               string `json:"file_hash"`
		GlobalID               int    `json:"global_id"`
		Version                int    `json:"version"`
		AfterMerge             bool   `json:"after_merge"`
		BeforeSplit            bool   `json:"before_split"`
		AfterSplit             bool   `json:"after_split"`
		WantMerge              bool   `json:"want_merge"`
		WantSplit              bool   `json:"want_split"`
		KeyBlock               bool   `json:"key_block"`
		VertSeqnoIncr          bool   `json:"vert_seqno_incr"`
		Flags                  int    `json:"flags"`
		GenUtime               string `json:"gen_utime"`
		StartLt                string `json:"start_lt"`
		EndLt                  string `json:"end_lt"`
		ValidatorListHashShort int    `json:"validator_list_hash_short"`
		GenCatchainSeqno       int    `json:"gen_catchain_seqno"`
		MinRefMcSeqno          int    `json:"min_ref_mc_seqno"`
		PrevKeyBlockSeqno      int    `json:"prev_key_block_seqno"`
		VertSeqno              int    `json:"vert_seqno"`
		MasterRefSeqno         int    `json:"master_ref_seqno"`
		RandSeed               string `json:"rand_seed"`
		CreatedBy              string `json:"created_by"`
		TxCount                int    `json:"tx_count"`
		MasterchainBlockRef    struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"masterchain_block_ref"`
		PrevBlocks []struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"prev_blocks"`
	} `json:"blocks"`
}
