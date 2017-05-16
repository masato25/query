package model

import (
	cmodel "github.com/Cepave/open-falcon-backend/common/model"
)

type Result struct {
	Dstype   string
	Step     int
	Endpoint string
	Counter  string
	Values   []*TimeSeriesData
}

type CostomGraphQueryResponse struct {
	Endpoint string            `json:"endpoint"`
	Counter  string            `json:"counter"`
	DsType   string            `json:"dstype"`
	Step     int               `json:"step"`
	Min      float64           `json:"min"`
	Mean     float64           `json:"mean"`
	Max      float64           `json:"max"`
	Values   []*cmodel.RRDData `json:"Values"` //大写为了兼容已经再用这个api的用户
}
