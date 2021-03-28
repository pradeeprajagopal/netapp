package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Luns struct {
	Embedded struct {
		NetappLunInventoryList []struct {
			Policies          []interface{} `json:"policies"`
			Latency           float64       `json:"latency"`
			AvailableCapacity float64       `json:"available_capacity"`
			Volume            struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"volume"`
			Cluster struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			Aggregate struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"aggregate"`
			Svm struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"svm"`
			ClusterFqdn    string `json:"cluster_fqdn"`
			QosPolicyGroup struct {
				ID                        interface{}   `json:"id"`
				Label                     interface{}   `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"qos_policy_group"`
			Iops          float64 `json:"iops"`
			NumberOfHours float64 `json:"number_of_hours"`
			ID            int     `json:"id"`
			Node          struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"node"`
			TotalCapacity float64 `json:"total_capacity"`
			Mbps          float64 `json:"mbps"`
			Status        string  `json:"status"`
			Lun           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"lun"`
		} `json:"netapp:lunInventoryList"`
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


//GetLuns ... retrieves all Luns information
func GetLuns() (Luns, error) {
	var luns Luns
	client := Auth(server,user,password)
	newurl := client.URL + "luns?max_records=15000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return luns, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return luns, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return luns, err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &luns)
	if err != nil {
		return luns, err
	}
	//return output
	return luns, nil

}
