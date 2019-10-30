package eth

// Web3Config contains settings that control the operation of the Web3 JSON-RPC method exposed
// via the /eth endpoint.
type Web3Config struct {
	// GetLogsMaxBlockRange specifies the maximum number of blocks eth_getLogs will query per request
	GetLogsMaxBlockRange uint64
	// GetFilterMaxBlockRange specifies the maximum number of blocks eth_getFilterChanges and eth_getFilterLogs will query per request
	GetFilterMaxBlockRange uint64
}

func DefaultWeb3Config() *Web3Config {
	return &Web3Config{
		GetLogsMaxBlockRange:   20,
		GetFilterMaxBlockRange: 20,
	}
}
