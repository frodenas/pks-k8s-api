/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	nt "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/network_transport"
)

// ListTranportZones list transport zones
func (nc *client) ListTranportZones() (*models.TransportZoneListResult, error) {
	params := nt.NewListTransportZonesParams()
	res, err := nc.client.NetworkTransport.ListTransportZones(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetTransportZone returns a Transport Zone with the specified ID
func (nc *client) GetTransportZone(ZoneID string) (*models.TransportZone, error) {
	params := nt.NewGetTransportZoneParams().WithZoneID(ZoneID)
	res, err := nc.client.NetworkTransport.GetTransportZone(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// UpdateTransportZone updates a transport zone
func (nc *client) UpdateTransportZone(transportZone *models.TransportZone) (*models.TransportZone, error) {
	params := nt.NewUpdateTransportZoneParams().WithTransportZone(transportZone).WithZoneID(transportZone.ID)
	res, err := nc.client.NetworkTransport.UpdateTransportZone(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// TagTransportZone adds tags to a transport zone
func (nc *client) TagTransportZone(transportZoneID string, tags []*models.Tag) (*models.TransportZone, error) {
	transportZone, err := nc.GetTransportZone(transportZoneID)
	if err != nil {
		return nil, err
	}
	err = ValidateTags(transportZone.ManagedResource, tags)
	if err != nil {
		return nil, err
	}

	transportZone.Tags = append(transportZone.Tags, tags...)
	res, err := nc.UpdateTransportZone(transportZone)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTransportNode gets transport node information from transport node ID
func (nc *client) GetTransportNode(transportNodeID string) (*models.TransportNode, error) {
	params := nt.NewGetTransportNodeParams().WithTransportnodeID(transportNodeID)
	res, err := nc.client.NetworkTransport.GetTransportNode(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ReadEdgeCluster reads edge cluster
func (nc *client) ReadEdgeCluster(edgeClusterID string) (*models.EdgeCluster, error) {
	params := nt.NewReadEdgeClusterParams().WithEdgeClusterID(edgeClusterID)

	res, err := nc.client.NetworkTransport.ReadEdgeCluster(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}
