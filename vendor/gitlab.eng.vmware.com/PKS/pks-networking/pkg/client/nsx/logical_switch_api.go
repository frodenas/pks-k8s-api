/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	lsw "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/logical_switching"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// CreateLogicalSwitch creates logical switch
func (nc *client) CreateLogicalSwitch(logicalSwitchModel *models.LogicalSwitch) (*models.LogicalSwitch, error) {
	params := lsw.NewCreateLogicalSwitchParams().WithLogicalSwitch(logicalSwitchModel)
	res, err := nc.client.LogicalSwitching.CreateLogicalSwitch(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// UpdateLogicalSwitch updates logical switch
func (nc *client) UpdateLogicalSwitch(logicalSwitchModel *models.LogicalSwitch) (*models.LogicalSwitch, error) {
	params := lsw.NewUpdateLogicalSwitchParams().WithLswitchID(logicalSwitchModel.ID).WithLogicalSwitch(logicalSwitchModel)
	res, err := nc.client.LogicalSwitching.UpdateLogicalSwitch(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetLogicalSwitch gets logical switch
func (nc *client) GetLogicalSwitch(logicalSwitchID string) (*models.LogicalSwitch, error) {
	params := lsw.NewGetLogicalSwitchParams().WithLswitchID(logicalSwitchID)
	res, err := nc.client.LogicalSwitching.GetLogicalSwitch(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteLogicalSwitch deletes the specified logical switch
func (nc *client) DeleteLogicalSwitch(LSwitchID string) error {
	params := lsw.NewDeleteLogicalSwitchParams()

	params.WithCascade(util.BoolPtr(true))
	params.WithLswitchID(LSwitchID)
	params.WithDetach(util.BoolPtr(true))

	_, err := nc.client.LogicalSwitching.DeleteLogicalSwitch(params, nc.auth)

	return err
}

// ListLogicalSwitches list logical switches
func (nc *client) ListLogicalSwitches() (*models.LogicalSwitchListResult, error) {
	params := lsw.NewListLogicalSwitchesParams()
	res, err := nc.client.LogicalSwitching.ListLogicalSwitches(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListLogicalPorts list logical switch ports
func (nc *client) ListLogicalPorts() (*models.LogicalPortListResult, error) {
	params := lsw.NewListLogicalPortsParams()
	res, err := nc.client.LogicalSwitching.ListLogicalPorts(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) GetLogicalPortsForLogicalSwitch(LogicalSwitchID string) ([]*models.LogicalPort, error) {
	var err error
	ports, err := nc.ListLogicalPorts()
	if err != nil {
		return nil, err
	}
	var res []*models.LogicalPort
	for _, port := range ports.Results {
		if util.StringVal(port.LogicalSwitchID) == LogicalSwitchID {
			res = append(res, port)
		}
	}
	return res, nil
}

// CreateLogicalPort creates a logical port
func (nc *client) CreateLogicalPort(logicalPortModel *models.LogicalPort) (*models.LogicalPort, error) {
	params := lsw.NewCreateLogicalPortParams().WithLogicalPort(logicalPortModel)
	res, err := nc.client.LogicalSwitching.CreateLogicalPort(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteLogicalPort will delete logical port based on provided ID with
// associated vif attachment detached
func (nc *client) DeleteLogicalPort(LogicalPortID string) error {
	params := lsw.NewDeleteLogicalPortParams().WithDetach(util.BoolPtr(true)).WithLportID(LogicalPortID)
	_, err := nc.client.LogicalSwitching.DeleteLogicalPort(params, nc.auth)
	return err
}

// GetLogicalPort gets the logical port based on provided ID
func (nc *client) GetLogicalPort(logicalPortID string) (*models.LogicalPort, error) {
	params := lsw.NewGetLogicalPortParams().WithLportID(logicalPortID)
	logicalPort, err := nc.client.LogicalSwitching.GetLogicalPort(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return logicalPort.Payload, err
}

// GetLogicalSwitchGivenName gets the logical switch matching display name
func (nc *client) GetLogicalSwitchGivenName(DisplayName string) (*models.LogicalSwitch, error) {
	res, err := nc.ListLogicalSwitches()
	if err != nil {
		return nil, err
	}
	for _, logicalSwitch := range res.Results {
		if logicalSwitch.DisplayName == DisplayName {
			return logicalSwitch, err
		}
	}
	return nil, err
}

// CreateSwitchingProfile creates a switching profile
func (nc *client) CreateSpoofGuardSwitchingProfile(switchingProfileModel *models.SpoofGuardSwitchingProfile) (*models.SpoofGuardSwitchingProfile, error) {
	params := lsw.NewCreateSwitchingProfileParams().WithSpoofGuardSwitchingProfile(switchingProfileModel)
	res, err := nc.client.LogicalSwitching.CreateSwitchingProfile(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListSwitchingProfilesByType lists switching profiles for a given type
func (nc *client) ListSwitchingProfilesByType(switchingProfileType string) (*models.SwitchingProfilesListResult, error) {
	params := lsw.NewListSwitchingProfilesParams().WithSwitchingProfileType(&switchingProfileType)
	res, err := nc.client.LogicalSwitching.ListSwitchingProfiles(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteSwitchingProfile deletes switching profile given a ID
func (nc *client) DeleteSwitchingProfile(switchingProfileID string) error {
	params := lsw.NewDeleteSwitchingProfileParams().WithSwitchingProfileID(switchingProfileID)
	_, err := nc.client.LogicalSwitching.DeleteSwitchingProfile(params, nc.auth)
	return err
}
