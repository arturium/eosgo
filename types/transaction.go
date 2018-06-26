package types

// TransactionHeader ...
type TransactionHeader struct {
	Status        string `json:"status"`
	CPUUsageUs    int    `json:"cpu_usage_us"`
	NetUsageWords int    `json:"net_usage_words"`
	Trx           Trx    `json:"trx"`
}

// Trx ...
type Trx struct {
	ID                    string   `json:"id"`
	Signatures            []string `json:"signatures"`
	Compression           string   `json:"compression"`
	PackedContextFreeData string   `json:"packed_context_free_data"`
	// ContextFreeDAta       []string  `json:"context_free_data"`
	PackedTrx   string      `json:"packed_trx"`
	Transaction Transaction `json:"transaction"`
}

// Transaction ...
type Transaction struct {
	Expiration      Time `json:"expiration"`
	RefBlockNum     int  `json:"ref_block_num"`
	RefBlockPrefix  int  `json:"ref_block_prefix"`
	MaxNetUsagWords int  `json:"max_net_usage_words"`
	MaxCPUUsageMs   int  `json:"max_cpu_usage_ms"`
	DelaySec        int  `json:"delay_sec"`
	// ContextFreeActions []Action `json:"context_free_actions"`
	Actions []Action `json:"actions"`
}