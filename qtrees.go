package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Qtrees struct {
	Embedded struct {
		NetappQtreeInventoryList []struct {
			Qtree struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"qtree"`
			FilesSoftLimit int `json:"files_soft_limit"`
			Volume         struct {
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
			Svm struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"svm"`
			ClusterFqdn         string      `json:"cluster_fqdn"`
			QuotaType           string      `json:"quota_type"`
			UserOrGroup         string      `json:"user_or_group"`
			DiskUsedPercentage  float64     `json:"disk_used_percentage"`
			FilesUsedPercentage interface{} `json:"files_used_percentage"`
			DiskHardLimit       float64     `json:"disk_hard_limit"`
			DiskSoftLimit       float64     `json:"disk_soft_limit"`
			ID                  int         `json:"id"`
			FilesHardLimit      float64     `json:"files_hard_limit"`
			Status              string      `json:"status"`
		} `json:"netapp:qtreeInventoryList"`
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


//GetAllQtrees ...
func GetAllQtrees() (Qtrees, error) {
	var qtrees Qtrees
	client := Auth(server,user,password)
	newUrl := client.URL + "qtrees?max_records=30000"
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return qtrees, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.capacity.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return qtrees, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return qtrees, err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &qtrees)
	if err != nil {
		return qtrees, err
	}
	//return output

	return qtrees, nil

}
