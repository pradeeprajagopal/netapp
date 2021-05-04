package netapp

import (
	"encoding/json"
)

type AccessEndPoints struct {
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
		DataProtocols []string `json:"data_protocols"`
		Gateways      []string `json:"gateways"`
		IP            struct {
			Address string `json:"address"`
			Netmask string `json:"netmask"`
		} `json:"ip"`
		Key  string `json:"key"`
		Mtu  int    `json:"mtu"`
		Name string `json:"name"`
		Svm  struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"svm"`
		UUID string `json:"uuid"`
		Vlan int    `json:"vlan"`
		Wwpn string `json:"wwpn"`
	} `json:"records"`
	TotalRecords int `json:"total_records"`
}

/*
GetAccessEndPointsV2...

Retrieves all access endpoints for an SVM, file share, or LUN. Network interfaces are provided with attributes, such as "ip.address", "ip.netmask", “gateway” and so on while FCP interfaces are provided with the “wwpn” attribute. You can provide the properties svm.key, fileshare.key, and lun.key to filter out the interfaces as follows:

svm.key - all types of interfaces for SAN and NAS (CIFS, NFS, ISCSI, and FCP protocols) are filtered.
fileshare.key - all interfaces with NAS protocols (CIFS and NFS) are filtered.
lun.key - all interfaces with SAN protocols (iSCSI and FCP) are filtered.
*/
func GetAccessEndPointsV2(key string) (AccessEndPoints, error) {
	var clusters AccessEndPoints
	query := "/api/storage-provider/access-endpoints?resource.key=" + key
	bodyText, err := getResponseBody(query)
	if err != nil {
		return clusters, err
	}
	err = json.Unmarshal(bodyText, &clusters)
	if err != nil {
		return clusters, err
	}
	return clusters, nil

}
