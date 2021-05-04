package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
func GetVolumes() (Volumes, error) {
	var volumes Volumes
	client := Auth(server, user, password)
	newurl := client.URL + "volumes?max_records=1000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return volumes, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return volumes, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return volumes, err
	}
	err = json.Unmarshal(bodyText, &volumes)
	if err != nil {
		return volumes, err
	}
	return volumes, nil

}

type VolumeV2 struct {
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	NumRecords int `json:"num_records"`
	Records    []struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Aggregates []struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"aggregates"`
		Autosize struct {
			Maximum int    `json:"maximum"`
			Mode    string `json:"mode"`
		} `json:"autosize"`
		Cluster struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"cluster"`
		CreateTime time.Time `json:"create_time"`
		Key        string    `json:"key"`
		Language   string    `json:"language"`
		Name       string    `json:"name"`
		Nas        struct {
			ExportPolicy struct {
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
				ID   int64  `json:"id"`
				Key  string `json:"key"`
				Name string `json:"name"`
			} `json:"export_policy"`
		} `json:"nas"`
		Qos struct {
			Policy struct {
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
				Key               string `json:"key"`
				MaxThroughputIops int    `json:"max_throughput_iops"`
				MaxThroughputMbps int    `json:"max_throughput_mbps"`
				MinThroughputIops int    `json:"min_throughput_iops"`
				Name              string `json:"name"`
				UUID              string `json:"uuid"`
			} `json:"policy"`
		} `json:"qos"`
		Snapmirror struct {
			IsProtected bool `json:"is_protected"`
		} `json:"snapmirror"`
		SnapshotPolicy struct {
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"snapshot_policy"`
		Space struct {
			Available int `json:"available"`
			Size      int `json:"size"`
			Used      int `json:"used"`
		} `json:"space"`
		State string `json:"state"`
		Style string `json:"style"`
		Svm   struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"svm"`
		Tiering struct {
			Policy string `json:"policy"`
		} `json:"tiering"`
		Type string `json:"type"`
		UUID string `json:"uuid"`
	} `json:"records"`
	TotalRecords int `json:"total_records"`
}

func getVolumesV2(query string) (VolumeV2, error) {
	var results VolumeV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return VolumeV2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return VolumeV2{}, err
	}
	return results, nil
}

//Retrieving a list of volumes in the datacenter
func GetVolumesV2() (VolumeV2, error) {
	var volumes VolumeV2
	query := "/api/datacenter/storage/volume"
	bodyText, err := getResponseBody(query)
	if err != nil {
		return VolumeV2{}, err
	}
	err = json.Unmarshal(bodyText, &volumes)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return volumes, err
	}
	return volumes, nil
}

//Retrieving a list of volumes in the datacenter filtered using "cluster.name"
func GetVolumesFromClusterNameV2(cluster string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?cluster.name=" + cluster
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "svm.name"
func GetVolumesFromSvmNameV2(svm string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?svm.name=" + svm
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "type"
func GetVolumesFromTypeV2(t string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?type=" + t
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "key"
func GetVolumesFromKeyV2(key string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?key=" + key
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "state"
func GetVolumesFromStateV2(state string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?state=" + state
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "style"
func GetVolumesFromStyleV2(style string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?state=" + style
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "uuid"
func GetVolumesFromUUIDV2(uuid string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?uuid=" + uuid
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "aggregate.key"
func GetVolumesFromAggregateKeyV2(key string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?aggregate.key=" + key
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "cluster.key"
func GetVolumesFromClusterKeyV2(key string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?cluster.key=" + key
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "svm.key"
func GetVolumesFromSVMKeyV2(key string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?svm.key=" + key
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "aggregate.uuid"
func GetVolumesFromAggregateUUIDV2(uuid string, environment string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?aggregate.uuid=" + uuid
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "cluster.uuid"
func GetVolumesFromClusterKeyUUIDV2(uuid string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?cluster.uuid=" + uuid
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "svm.uuid"
func GetVolumesFromSVMUUIDV2(uuid string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?svm.uuid=" + uuid
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "aggregate.name"
func GetVolumesFromAggregateNameV2(name string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?aggregate.name=" + name
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "nas.export_policy.key"
func GetVolumesFromNASKeyV2(key string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?nas.export_policy.key=" + key
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "nas.export_policy.id"
func GetVolumesFromNASIDV2(id string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?nas.export_policy.id=" + id
	return getVolumesV2(query)
}

//Retrieving a list of volumes in the datacenter filtered using "nas.export_policy.name"
func GetVolumesFromNASNameV2(name string) (VolumeV2, error) {
	query := "/api/datacenter/storage/volumes?nas.export_policy.name=" + name
	return getVolumesV2(query)
}
