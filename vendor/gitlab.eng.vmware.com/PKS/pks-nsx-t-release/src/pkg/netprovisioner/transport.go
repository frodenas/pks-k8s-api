/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

// CheckTransportZone checks if given transport zone is valid
func (p *nsxNetworkProvisioner) CheckTransportZone(transportZoneID string) error {
	if err := util.EnsureParams(transportZoneID); err != nil {
		return err
	}
	_, otzErr := p.GetTransportZone(transportZoneID)
	return otzErr
}

// UntagTransportZone removes the cluster tag from the overlay transport zone
func (p *nsxNetworkProvisioner) UntagTransportZone(transportZoneID string, clusterName string) error {
	var err error
	var TZ *models.TransportZone

	if err = util.EnsureParams(transportZoneID, clusterName); err != nil {
		return err
	}
	if TZ, err = p.GetTransportZone(transportZoneID); err != nil {
		return err
	}
	nsx.RemoveTag(&TZ.ManagedResource, models.Tag{
		Scope: nsx.NcpTagKeyCluster,
		Tag:   clusterName,
	})
	if _, err = p.UpdateTransportZone(TZ); err != nil {
		return err
	}
	return nil
}

// GetOverlayTransportZoneIDFromTransportNode retrieves transport zone ID with overlay type from transport node ID
func (p *nsxNetworkProvisioner) GetOverlayTransportZoneIDFromTransportNode(transportNodeID string) (string, error) {
	var err error
	if err = util.EnsureParams(transportNodeID); err != nil {
		return "", err
	}

	var res *models.TransportNode
	if res, err = p.GetTransportNode(transportNodeID); err != nil {
		return "", err
	}

	var transportZoneID string
	var transportZone *models.TransportZone
	for _, transportZoneEndpoint := range res.TransportZoneEndpoints {
		transportZoneID = util.StringVal(transportZoneEndpoint.TransportZoneID)
		if transportZone, err = p.GetTransportZone(transportZoneID); err != nil {
			return "", err
		}
		if util.StringVal(transportZone.TransportType) == nsx.TransportZoneTypeOverlay {
			return transportZoneID, nil
		}
	}

	return "", fmt.Errorf("No transport zone with overlay type found in transport node %s", transportNodeID)
}
