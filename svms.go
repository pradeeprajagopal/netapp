package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Svms struct {
	Embedded struct {
		NetappSvmInventoryList []struct {
			Policies          []interface{} `json:"policies"`
			Latency           float64       `json:"latency"`
			AvailableCapacity float64       `json:"available_capacity"`
			Cluster           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			Svm struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"svm"`
			ClusterFqdn   string  `json:"cluster_fqdn"`
			Iops          float64 `json:"iops"`
			NumberOfHours int     `json:"number_of_hours"`
			ID            int     `json:"id"`
			TotalCapacity float64 `json:"total_capacity"`
			Mbps          float64 `json:"mbps"`
			Status        string  `json:"status"`
		} `json:"netapp:svmInventoryList"`
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

//GetAllSvms ...
func GetAllSvms() (Svms, error) {
	var svms Svms
	client := Auth(server,user,password)
	newurl := client.URL + "svms?max_records=30000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return svms, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return svms, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return svms, err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &svms)
	if err != nil {
		return svms, err
	}
	//return output
	return svms, nil

}
