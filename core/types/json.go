package json

// https://mholt.github.io/json-to-go/

type GoogleImage struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
		Status  string `json:"status"`
		Details []struct {
			Type     string `json:"@type"`
			Reason   string `json:"reason"`
			Domain   string `json:"domain"`
			Metadata struct {
				Consumer string `json:"consumer"`
				Service  string `json:"service"`
			} `json:"metadata"`
		} `json:"details"`
	} `json:"error"`
}

