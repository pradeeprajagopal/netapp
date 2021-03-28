package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Volumes struct {
	Embedded struct {
		NetappVolumeInventoryList []struct {
			SystemManagerURL             string `json:"system_manager_url"`
			RecommendationGenerationTime string `json:"recommendation_generation_time"`
			TieringPolicy                string `json:"tiering_policy"`
			Policies                     []struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"policies"`
			CloudRecommendation string      `json:"cloud_recommendation"`
			RecommendationData  string      `json:"recommendation_data"`
			Latency             interface{} `json:"latency"`
			AvailableCapacity   float64     `json:"available_capacity"`
			DiskType            string      `json:"disk_type"`
			Volume              struct {
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
			Iops          interface{} `json:"iops"`
			NumberOfHours int         `json:"number_of_hours"`
			ID            int         `json:"id"`
			Node          struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"node"`
			TotalCapacity float64     `json:"total_capacity"`
			Iopspertb     interface{} `json:"iopspertb"`
			Mbps          interface{} `json:"mbps"`
			VolumeType    string      `json:"volume_type"`
			ColdData      interface{} `json:"cold_data"`
			Status        string      `json:"status"`
		} `json:"netapp:volumeInventoryList"`
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


//GetVolumes ...
//ToDo: Enable Pagination
func GetVolumes() (Volumes,error) {
	var volumes Volumes
	client := Auth(server,user,password)
	newurl := client.URL + "volumes?max_records=1000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return volumes,err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return volumes,err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return volumes,err
	}
	err = json.Unmarshal(bodyText, &volumes)
	if err != nil {
		return volumes,err
	}
	return volumes,nil

}
