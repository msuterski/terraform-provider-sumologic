// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------\
package sumologic

import (
	"encoding/json"
	"fmt"
)

func (s *Client) CreateRole(role Role) (string, error) {
	data, err := s.Post("v1/roles", role)
	if err != nil {
		return "", err
	}

	var createdrole Role
	err = json.Unmarshal(data, &createdrole)
	if err != nil {
		return "", err
	}

	return createdrole.ID, nil
}

func (s *Client) DeleteRole(id string) error {
	_, err := s.Delete(fmt.Sprintf("v1/roles/%s", id))
	return err
}

func (s *Client) GetRole(id string) (*Role, error) {
	data, _, err := s.Get(fmt.Sprintf("v1/roles/%s", id))
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	var role Role
	err = json.Unmarshal(data, &role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *Client) GetRoleName(name string) (*Role, error) {
	data, _, err := s.Get(fmt.Sprintf("v1/roles?name=%s", name))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("role with name '%s' does not exist", name)
	}

	var response RoleResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response.Roles[0], nil
}

func (s *Client) UpdateRole(role Role) error {
	url := fmt.Sprintf("v1/roles/%s", role.ID)

	role.ID = ""

	_, err := s.Put(url, role)
	return err
}

type RoleResponse struct {
	Roles []Role `json:"data"`
}

// models
type Role struct {
	ID string `json:"id,omitempty"`
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description"`
	// A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.
	FilterPredicate string `json:"filterPredicate"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are   ### Connections   - manageConnections   ### Collectors   - manageCollectors   - viewCollectors   ### Dashboards   - shareDashboardWhitelist   - shareDashboardWorld   ### Data Management   - manageContent   - manageDataVolumeFeed   - manageFieldExtractionRules   - manageIndexes   - manageS3DataForwarding   ### Metrics   - manageMonitors   - metricsExtraction   ### Security   - ipWhitelisting   - manageAccessKeys   - manageAuditDataFeed   - managePasswordPolicy   - manageSaml   - manageSupportAccountAccess   - manageUsersAndRoles   - shareDashboardOutsideOrg
	Capabilities []string `json:"capabilities"`
}
