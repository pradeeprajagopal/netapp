package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Interfaces struct {
	Embedded struct {
		NetappInterfaceInventoryList []struct {
			LifType  string `json:"lif_type"`
			HomeNode struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"home_node"`
			CurrentNode struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"current_node"`
			Policies []interface{} `json:"policies"`
			HomePort struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"home_port"`
			CurrentPort struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"current_port"`
			Role    string      `json:"role"`
			Latency interface{} `json:"latency"`
			Cluster struct {
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
			ClusterFqdn   string      `json:"cluster_fqdn"`
			Iops          interface{} `json:"iops"`
			HomeLocation  string      `json:"home_location"`
			NumberOfHours int         `json:"number_of_hours"`
			ID            int         `json:"id"`
			Mbps          float64     `json:"mbps"`
			Lif           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"lif"`
			Status          string `json:"status"`
			CurrentLocation string `json:"current_location"`
		} `json:"netapp:interfaceInventoryList"`
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


//GetInterfaces ... Provides interfaces information
func GetInterfaces() (Interfaces, error) {
	var interfaces Interfaces
	client := Auth(server,user,password)
	newUrl := client.URL + "interfaces?max_records=1000"
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return interfaces, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return interfaces, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return interfaces, err
	}
	//fmt.Println(string(bodyText))
	var v interface{}
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return interfaces, err
	}
	err = json.Unmarshal(bodyText, &interfaces)
	if err != nil {
		return interfaces, err
	}
	return interfaces,nil

}
