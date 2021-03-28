package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Nodes struct {
	Embedded struct {
		NetappNodeInventoryList []struct {
			Policies []struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"policies"`
			Utilization       float64 `json:"utilization"`
			Latency           float64 `json:"latency"`
			AvailableCapacity float64 `json:"available_capacity"`
			Cluster           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			ClusterFqdn            string      `json:"cluster_fqdn"`
			Iops                   float64     `json:"iops"`
			FlashCacheReadsPercent interface{} `json:"flash_cache_reads_percent"`
			NumberOfHours          int         `json:"number_of_hours"`
			UsedHeadroom           float64     `json:"used_headroom"`
			ID                     int         `json:"id"`
			Node                   struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"node"`
			TotalCapacity float64 `json:"total_capacity"`
			Mbps          float64 `json:"mbps"`
			Status        string  `json:"status"`
		} `json:"netapp:nodeInventoryList"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
	TotalCount int `json:"totalCount"`
}
//GetAllNodes ... retrieves all node information
func GetAllNodes() (Nodes, error) {
	var nodes Nodes
	client := Auth(server,user,password)
	newUrl := client.URL + "nodes?max_records=15000"
	//logger.Log(newUrl)
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return nodes, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return nodes, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nodes, err
	}
	err = json.Unmarshal(bodyText, &nodes)
	if err != nil {
		return nodes, err
	}

	return nodes, nil

}
