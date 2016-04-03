package elus

import (
	"encoding/json"
	"log"

    "fmt"
    "time"
	cmodel "github.com/Cepave/common/model"
	"github.com/Cepave/query/graph"
)

type server struct{}

func setq(startTs int64, endTs int64) cmodel.GraphQueryParam {
    request := cmodel.GraphQueryParam{
        Start:    startTs, 
        End:        endTs,
        ConsolFun: "AVERAGE",
        Endpoint:  "",
        Counter:   "",
    }
    return request
}

func rrdQueryForErl(endpointList []string, counterList []string) (resp []*cmodel.GraphQueryResponse) {
    queryParams := setq(time.Now().Unix() - 86400, time.Now().Unix())
    for _, enp := range endpointList {
		queryParams.Endpoint = enp
		for _, con := range counterList {
			queryParams.Counter = con
            fmt.Printf("%v", queryParams)
			res, err := graph.QueryOne(queryParams)
			if err != nil {
				log.Printf("@@@@@-> %v", err.Error())
			}
			resp = append(resp, res)
		}
	}
	return
}

func Query(endpoints []string, counterList []string) (result string){
    resTmp := rrdQueryForErl(endpoints, counterList)
    
    for idx, result := range resTmp {
		if result.Values == nil {
			result.Values = []*cmodel.RRDData{}
		}
		resTmp[idx] = result
	}
	res, _ := json.Marshal(resTmp)
    result = string(res)
    return
}
