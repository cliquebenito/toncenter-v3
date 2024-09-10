package toncenter

type AddressBookStruct struct {
	UserFriendly string `json:"user_friendly"`
}
type AddressBookResponse map[string]AddressBookStruct

//	type TxLastRest struct {
//		AddressBookResponse `json:"address_book"`
//	}

type TransactionByMasterchainBlockParams struct {
	Seqno  int    `json:"seqno" validate:"required"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Sort   string `json:"sort"`
}
type TransactionByMasterChainBlockStruct struct {
	Transactions []struct {
		Account       string `json:"account"`
		Hash          string `json:"hash"`
		Lt            string `json:"lt"`
		Now           int    `json:"now"`
		OrigStatus    string `json:"orig_status"`
		EndStatus     string `json:"end_status"`
		TotalFees     string `json:"total_fees"`
		PrevTransHash string `json:"prev_trans_hash"`
		PrevTransLt   string `json:"prev_trans_lt"`
		Description   struct {
			Type   string `json:"type"`
			Action struct {
				Valid       bool `json:"valid"`
				Success     bool `json:"success"`
				NoFunds     bool `json:"no_funds"`
				ResultCode  int  `json:"result_code"`
				TotActions  int  `json:"tot_actions"`
				MsgsCreated int  `json:"msgs_created"`
				SpecActions int  `json:"spec_actions"`
				TotMsgSize  struct {
					Bits  string `json:"bits"`
					Cells string `json:"cells"`
				} `json:"tot_msg_size"`
				StatusChange   string `json:"status_change"`
				SkippedActions int    `json:"skipped_actions"`
				ActionListHash string `json:"action_list_hash"`
			} `json:"action"`
			Aborted   bool `json:"aborted"`
			IsTock    bool `json:"is_tock"`
			Destroyed bool `json:"destroyed"`
			ComputePh struct {
				Mode             int    `json:"mode"`
				Type             string `json:"type"`
				Success          bool   `json:"success"`
				GasFees          string `json:"gas_fees"`
				GasUsed          string `json:"gas_used"`
				VMSteps          int    `json:"vm_steps"`
				ExitCode         int    `json:"exit_code"`
				GasLimit         string `json:"gas_limit"`
				MsgStateUsed     bool   `json:"msg_state_used"`
				AccountActivated bool   `json:"account_activated"`
				VMInitStateHash  string `json:"vm_init_state_hash"`
				VMFinalStateHash string `json:"vm_final_state_hash"`
			} `json:"compute_ph"`
			StoragePh struct {
				StatusChange         string `json:"status_change"`
				StorageFeesCollected string `json:"storage_fees_collected"`
			} `json:"storage_ph"`
		} `json:"description"`
		BlockRef struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"block_ref"`
		InMsg              any   `json:"in_msg"`
		OutMsgs            []any `json:"out_msgs"`
		AccountStateBefore struct {
			Hash          string `json:"hash"`
			Balance       string `json:"balance"`
			AccountStatus string `json:"account_status"`
			FrozenHash    any    `json:"frozen_hash"`
			CodeHash      string `json:"code_hash"`
			DataHash      string `json:"data_hash"`
		} `json:"account_state_before"`
		AccountStateAfter struct {
			Hash          string `json:"hash"`
			Balance       string `json:"balance"`
			AccountStatus string `json:"account_status"`
			FrozenHash    any    `json:"frozen_hash"`
			CodeHash      string `json:"code_hash"`
			DataHash      string `json:"data_hash"`
		} `json:"account_state_after"`
		McBlockSeqno int `json:"mc_block_seqno"`
	} `json:"transactions"`
	AddressBookResponse `json:"address_book"`
}

//	type TransactionByMasterChainBlockStruct struct {
//		Transactions []struct {
//			Account       string `json:"account"`
//			Hash          string `json:"hash"`
//			Lt            string `json:"lt"`
//			Now           int    `json:"now"`
//			OrigStatus    string `json:"orig_status"`
//			EndStatus     string `json:"end_status"`
//			TotalFees     string `json:"total_fees"`
//			PrevTransHash string `json:"prev_trans_hash"`
//			PrevTransLt   string `json:"prev_trans_lt"`
//			Description   string `json:"description"`
//			BlockRef      struct {
//				Workchain int    `json:"workchain"`
//				Shard     string `json:"shard"`
//				Seqno     int    `json:"seqno"`
//			} `json:"block_ref"`
//			InMsg struct {
//				Hash           string `json:"hash"`
//				Source         string `json:"source"`
//				Destination    string `json:"destination"`
//				Value          string `json:"value"`
//				FwdFee         string `json:"fwd_fee"`
//				IhrFee         string `json:"ihr_fee"`
//				CreatedLt      string `json:"created_lt"`
//				CreatedAt      string `json:"created_at"`
//				Opcode         string `json:"opcode"`
//				IhrDisabled    bool   `json:"ihr_disabled"`
//				Bounce         bool   `json:"bounce"`
//				Bounced        bool   `json:"bounced"`
//				ImportFee      string `json:"import_fee"`
//				MessageContent struct {
//					Hash    string `json:"hash"`
//					Body    string `json:"body"`
//					Decoded struct {
//					} `json:"decoded"`
//				} `json:"message_content"`
//				InitState struct {
//					Hash string `json:"hash"`
//					Body string `json:"body"`
//				} `json:"init_state"`
//			} `json:"in_msg"`
//			OutMsgs []struct {
//				Hash           string `json:"hash"`
//				Source         string `json:"source"`
//				Destination    string `json:"destination"`
//				Value          string `json:"value"`
//				FwdFee         string `json:"fwd_fee"`
//				IhrFee         string `json:"ihr_fee"`
//				CreatedLt      string `json:"created_lt"`
//				CreatedAt      string `json:"created_at"`
//				Opcode         string `json:"opcode"`
//				IhrDisabled    bool   `json:"ihr_disabled"`
//				Bounce         bool   `json:"bounce"`
//				Bounced        bool   `json:"bounced"`
//				ImportFee      string `json:"import_fee"`
//				MessageContent struct {
//					Hash    string `json:"hash"`
//					Body    string `json:"body"`
//					Decoded struct {
//					} `json:"decoded"`
//				} `json:"message_content"`
//				InitState struct {
//					Hash string `json:"hash"`
//					Body string `json:"body"`
//				} `json:"init_state"`
//			} `json:"out_msgs"`
//			AccountStateBefore struct {
//				Hash          string `json:"hash"`
//				Balance       string `json:"balance"`
//				AccountStatus string `json:"account_status"`
//				FrozenHash    string `json:"frozen_hash"`
//				CodeHash      string `json:"code_hash"`
//				DataHash      string `json:"data_hash"`
//			} `json:"account_state_before"`
//			AccountStateAfter struct {
//				Hash          string `json:"hash"`
//				Balance       string `json:"balance"`
//				AccountStatus string `json:"account_status"`
//				FrozenHash    string `json:"frozen_hash"`
//				CodeHash      string `json:"code_hash"`
//				DataHash      string `json:"data_hash"`
//			} `json:"account_state_after"`
//			McBlockSeqno int `json:"mc_block_seqno"`
//		} `json:"transactions"`
//		AddressBookResponse `json:"address_book"`
//	}
type TransactionStruct struct {
	Transactions []struct {
		Account       string `json:"account"`
		Hash          string `json:"hash"`
		Lt            string `json:"lt"`
		Now           int    `json:"now"`
		McBlockSeqno  int    `json:"mc_block_seqno"`
		TraceID       string `json:"trace_id"`
		PrevTransHash string `json:"prev_trans_hash"`
		PrevTransLt   string `json:"prev_trans_lt"`
		OrigStatus    string `json:"orig_status"`
		EndStatus     string `json:"end_status"`
		TotalFees     string `json:"total_fees"`
		Description   struct {
			Type      string `json:"type"`
			Aborted   bool   `json:"aborted"`
			Destroyed bool   `json:"destroyed"`
			IsTock    bool   `json:"is_tock"`
			StoragePh struct {
				StorageFeesCollected string `json:"storage_fees_collected"`
				StatusChange         string `json:"status_change"`
			} `json:"storage_ph"`
			ComputePh struct {
				Skipped          bool   `json:"skipped"`
				Success          bool   `json:"success"`
				MsgStateUsed     bool   `json:"msg_state_used"`
				AccountActivated bool   `json:"account_activated"`
				GasFees          string `json:"gas_fees"`
				GasUsed          string `json:"gas_used"`
				GasLimit         string `json:"gas_limit"`
				Mode             int    `json:"mode"`
				ExitCode         int    `json:"exit_code"`
				VMSteps          int    `json:"vm_steps"`
				VMInitStateHash  string `json:"vm_init_state_hash"`
				VMFinalStateHash string `json:"vm_final_state_hash"`
			} `json:"compute_ph"`
			Action struct {
				Success        bool   `json:"success"`
				Valid          bool   `json:"valid"`
				NoFunds        bool   `json:"no_funds"`
				StatusChange   string `json:"status_change"`
				ResultCode     int    `json:"result_code"`
				TotActions     int    `json:"tot_actions"`
				SpecActions    int    `json:"spec_actions"`
				SkippedActions int    `json:"skipped_actions"`
				MsgsCreated    int    `json:"msgs_created"`
				ActionListHash string `json:"action_list_hash"`
				TotMsgSize     struct {
					Cells string `json:"cells"`
					Bits  string `json:"bits"`
				} `json:"tot_msg_size"`
			} `json:"action"`
		} `json:"description"`
		BlockRef struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			Seqno     int    `json:"seqno"`
		} `json:"block_ref"`
		InMsg              interface{}   `json:"in_msg"`
		OutMsgs            []interface{} `json:"out_msgs"`
		AccountStateBefore struct {
			Hash          string      `json:"hash"`
			Balance       string      `json:"balance"`
			AccountStatus string      `json:"account_status"`
			FrozenHash    interface{} `json:"frozen_hash"`
			DataHash      string      `json:"data_hash"`
			CodeHash      string      `json:"code_hash"`
		} `json:"account_state_before"`
		AccountStateAfter struct {
			Hash          string      `json:"hash"`
			Balance       string      `json:"balance"`
			AccountStatus string      `json:"account_status"`
			FrozenHash    interface{} `json:"frozen_hash"`
			DataHash      string      `json:"data_hash"`
			CodeHash      string      `json:"code_hash"`
		} `json:"account_state_after"`
	} `json:"transactions"`
	AddressBookResponse `json:"address_book"`
}

type MasterChainBlockShardsStruct struct {
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
type AddressBookParams struct {
	Address []string `json:"address" validate:"required"`
}

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
type TransactionsParams struct {
	WorkChain      string   `json:"workchain"`
	Shard          string   `json:"shard"`
	Seqno          string   `json:"seqno"`
	Account        []string `json:"account"`
	ExcludeAccount []string `json:"exclude_account"`
	Hash           string   `json:"hash"`
	Lt             string   `json:"lt"`
	StartUtime     string   `json:"start_utime"`
	EndUtime       string   `json:"end_utime"`
	StartLt        string   `json:"start_lt"`
	EndLt          string   `json:"end_lt"`
	Limit          string   `json:"limit"`
	Offset         string   `json:"offset"`
	Sort           string   `json:"sort"`
}

type MasterChainBlockShardStateStruct struct {
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
