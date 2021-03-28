package netapp

import (
	"encoding/json"
	"io/ioutil"
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
