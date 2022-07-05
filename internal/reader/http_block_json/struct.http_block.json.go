package http_block_json

// HttpBlockConfig 对应
type HttpBlockConfig struct {
	JsonDesc string `json:"json_desc"`
	BlackIp string `json:"black_ip"`
	WhiteIp string `json:"white_ip"`
	MaxSniffHourIpPV int64 `json:"maxSniffHourIpPV"`
	MaxDDOSHourIpPV int64 `json:"maxDDOSHourIpPV"`
}
