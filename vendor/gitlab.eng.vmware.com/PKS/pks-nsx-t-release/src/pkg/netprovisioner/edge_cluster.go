/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

// ExtractTransportZoneIDFromEdgeCluster gets transport zone ID from edge cluster ID
func (p *nsxNetworkProvisioner) ExtractTransportZoneIDFromEdgeCluster(edgeClusterID string) (string, error) {
	if err := util.EnsureParams(edgeClusterID); err != nil {
		return "", err
	}

	edgeCluster, err := p.ReadEdgeCluster(edgeClusterID)
	if err != nil {
		return "", err
	}
	if len(edgeCluster.Members) == 0 {
		return "", fmt.Errorf("No transport node associated with edge cluster %s", edgeClusterID)
	}
	transportNodeID := util.StringVal(edgeCluster.Members[0].TransportNodeID)

	transportZoneID, err := p.GetOverlayTransportZoneIDFromTransportNode(transportNodeID)
	if err != nil {
		return "", err
	}

	return transportZoneID, nil
}
