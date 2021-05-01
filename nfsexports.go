package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type NFS struct {
	Embedded struct {
		NetappNfsexportInventoryList []struct {
			Qtree struct {
				ID                        int           `json:"id"`
				Label                     interface{}   `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"qtree"`
			SecurityStyle string `json:"security_style"`
			ExportPolicy  string `json:"export_policy"`
			RuleIndex     string `json:"rule_index"`
			Volume        struct {
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
			ClusterFqdn        string `json:"cluster_fqdn"`
			ClientMatch        string `json:"client_match"`
			VolumeState        string `json:"volume_state"`
			AccessProtocols    string `json:"access_protocols"`
			ReadOnlyAccess     string `json:"read_only_access"`
			ReadWriteAccess    string `json:"read_write_access"`
			UnixPermission     string `json:"unix_permission"`
			JunctionPathActive string `json:"junction_path_active"`
			Status             string `json:"status"`
			JunctionPath       string `json:"junction_path"`
		} `json:"netapp:nfsexportInventoryList"`
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

type NfsV2 struct {
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
		ID    int  `json:"id"`
		Key   string `json:"key"`
		Name  string `json:"name"`
		Rules []struct {
			AnonymousUser string `json:"anonymous_user"`
			Clients       []struct {
				Match string `json:"match"`
			} `json:"clients"`
			Index     int      `json:"index"`
			Protocols []string `json:"protocols"`
			RoRule    []string `json:"ro_rule"`
			RwRule    []string `json:"rw_rule"`
			Superuser []string `json:"superuser"`
		} `json:"rules"`
		Svm struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"svm"`
	} `json:"records"`
	TotalRecords int `json:"total_records"`
}

//GetAllNfsInfo ... retreives all NFS export information
func GetAllNfsInfo() (NFS, error) {
	var nfs NFS
	client := Auth(server,user,password)
	newurl := client.URL + "nfsexports?max_records=15000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return nfs, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.health.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return nfs, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nfs, err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &nfs)
	if err != nil {
		return nfs, err
	}
	return nfs, nil

}


func getNfsInfoV2(query string) (NfsV2, error) {
	var results NfsV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return NfsV2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return NfsV2{}, err
	}
	return results, nil
}


//GetAllNfsInfoV2 get all NFs info
func GetAllNfsInfoV2() (NfsV2, error) {
	query := "/api/datacenter/protocols/nfs/export-policies?offset=1000"
	return getNfsInfoV2(query)
}

//GetNfsInfoV2FromName get NFSExport from name
func GetNfsInfoV2FromName(name string) (NfsV2, error) {
	query := "/api/datacenter/protocols/nfs/export-policies?name="+name
	return getNfsInfoV2(query)
}

//GetNfsInfoV2FromCluster get NFS from Cluster
func GetNfsInfoV2FromCluster(name string) (NfsV2, error) {
	query := "/api/datacenter/protocols/nfs/export-policies?cluster.name="+name
	return getNfsInfoV2(query)
}

//GetNfsInfoV2FromCluster get NFS from Svm
func GetNfsInfoV2FromSvm(name string) (NfsV2, error) {
	query := "/api/datacenter/protocols/nfs/export-policies?svm.name="+name
	return getNfsInfoV2(query)
}