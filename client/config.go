package client

type ClientConfig struct {
	URL               string
	Bearer            string
	Headers           map[string]string
	RequestId         string  `default:"request_id"`
	Timeout           float64 `default:"5"`
	Backoff           float64 `default:"0.1"`
	Retry             int     `default:"3"`
	IgnoreStatusCodes []int
}
