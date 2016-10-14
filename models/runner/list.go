package runner

// ListReq represents the request to get a list of jobs
type ListReq struct {
	Page     string `URIParam:"yes" json:"-"`
	PageSize string `URIParam:"yes" json:"-"`
}

// ListRes represents the response from a get list of jobs request
type ListRes struct {
	TotalSize int64 `json:"totalSize"`
	Page      int64 `json:"page"`
	PageSize  int64 `json:"size"`
	Jobs      []Job `json:"results"`
}
