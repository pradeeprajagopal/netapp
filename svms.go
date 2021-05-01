package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

type Items struct {
	links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Key  string `json:"key"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type SvmV2 struct {
	links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	numRecords int `json:"num_records"`
	Records    []struct {
		links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Aggregates []Items `json:"aggregates"`
		Cifs       struct {
			AdDomain struct {
				Fqdn string `json:"fqdn"`
			} `json:"ad_domain"`
			Enabled bool   `json:"enabled"`
			Name    string `json:"name"`
		} `json:"cifs"`
		Cluster struct {
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"cluster"`
		DNS struct {
			Domains []string `json:"domains"`
			Servers []string `json:"servers"`
		} `json:"dns"`
		Fcp struct {
			Enabled bool `json:"enabled"`
		} `json:"fcp"`
		Ipspace struct {
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"ipspace"`
		Iscsi struct {
			Enabled bool `json:"enabled"`
		} `json:"iscsi"`
		Key      string `json:"key"`
		Language string `json:"language"`
		Ldap     struct {
			Enabled bool `json:"enabled"`
		} `json:"ldap"`
		Name string `json:"name"`
		Nfs  struct {
			Enabled bool `json:"enabled"`
		} `json:"nfs"`
		Nis struct {
			Domain  string   `json:"domain"`
			Enabled bool     `json:"enabled"`
			Servers []string `json:"servers"`
		} `json:"nis"`
		Nvme struct {
			Enabled bool `json:"enabled"`
		} `json:"nvme"`
		SnapshotPolicy struct {
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"snapshot_policy"`
		State   string `json:"state"`
		Subtype string `json:"subtype"`
		UUID    string `json:"uuid"`
	} `json:"records"`
	totalRecords int `json:"total_records"`
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



func getSvmsV2(query string) (SvmV2, error) {
	var results SvmV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return SvmV2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return SvmV2{}, err
	}
	return results, nil
}

//GetSvmsV2 Retrieving a list of SVMs in the datacenter
func GetSvmsV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms"
	return getSvmsV2(query)

}

//GetSvmsNFSEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "nfs.enabled":
func GetSvmsNFSEnabledV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms?nfs.enabled=true"
	return getSvmsV2(query)
}

//GetSvmsFromClusterV2 Retrieving a list of SVMs in the datacenter filtered using "cluster.name"
func GetSvmsFromClusterV2(cluster string) (SvmV2, error) {
	query := "/api/datacenter/svm/svms?cluster.name=" + cluster
	return getSvmsV2(query)
}

//GetSvmsFromNFSEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "nfs.enabled"
func GetSvmsFromNFSEnabledV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms?nfs.enabled=true"
	return getSvmsV2(query)
}

//GetSvmsFromNVMeEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "nvme.enabled"
func GetSvmsFromNVMeEnabledV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms?nvme.enabled=true"
	return getSvmsV2(query)

}

//GetSvmsFromFcpEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "fcp.enabled"
func GetSvmsFromFcpEnabledV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms?fcp.enabled=true"
	return getSvmsV2(query)
}

//GetSvmsFromFcpEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "iscsi.enabled"
func GetSvmsFromiScsiEnabledV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms?iscsi.enabled=true"
	return getSvmsV2(query)
}

//GetSvmsFromNameV2 Retrieving a list of SVMs in the datacenter filtered using "name"
func GetSvmsFromNameV2(name string) (SvmV2, error) {
	query := "/api/datacenter/svm/svms?name=" + name
	return getSvmsV2(query)

}

//GetSvmsFromCifsEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "cifs.enabled"
func GetSvmsFromCifsEnabledV2() (SvmV2, error) {
	query := "/api/datacenter/svm/svms?cifs.enabled=true"
	return getSvmsV2(query)

}

//GetSvmsFromCifsEnabledV2 Retrieving a list of SVMs in the datacenter filtered using "cifs.name"
func GetSvmsFromCifsNameV2(name string) (SvmV2, error) {
	query := "/api/datacenter/svm/svms?cifs.name=" + name
	return getSvmsV2(query)

}
