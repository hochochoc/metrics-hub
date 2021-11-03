package model

type HttpResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

type ReportStatistics struct {
	Mean   float64 `json:"mean"`
	Stddev float64 `json:"stddev"`
}
