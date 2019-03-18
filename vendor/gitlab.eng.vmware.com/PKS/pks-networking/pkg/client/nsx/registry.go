/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"errors"
	"fmt"
)

func (nc *client) ResourceCollectFunc(resourceType string) ResourceCollectFunc {
	switch resourceType {
	case ResourceTypePersistenceProfile:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLoadBalancerPersistenceProfiles()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeFirewallSection:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListFirewallSections()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeNSGroup:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListNSGroups()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLbService:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLoadBalancerServices()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLbRule:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLoadBalancerRules()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLbPool:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLoadBalancerPools()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLogicalPort:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLogicalPorts()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeIPSet:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				var err error
				lres, err := nc.ListIPSets()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLbVirtualServer:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLoadBalancerVirtualServers()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case RouterTypeTier1:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListT1LogicalRouters()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case RouterTypeTier0:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListT0LogicalRouters()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					if IsNcpSharedResource(&r.ManagedResource) {
						res = append(res, interface{}(r))
					}
				}
				if len(res) == 0 {
					nc.Warn("missing cluster tier-0 router\n")
					return nil, errors.New("missing cluster tier-0 router")
				} else if len(res) > 1 {
					nc.Warn("found multiple(%d) tier-0 routers\n", len(res))
					return nil, errors.New(fmt.Sprintf("found multiple(%d) tier-0 routers\n", len(res)))
				}
				return res, nil
			})
		}
	case ResourceTypeLogicalRouter:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLogicalRouters()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLogicalSwitch:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLogicalSwitches()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeIPPool:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListIPPools()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypeLbProfile:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				lres, err := nc.ListLoadBalancerApplicationProfiles()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, r := range lres.Results {
					res = append(res, interface{}(r))
				}
				return res, nil
			})
		}
	case ResourceTypePrincipalIdentity:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				pis, err := nc.GetPrincipalIdentities()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, pi := range pis.Results {
					res = append(res, interface{}(pi))
				}
				return res, nil
			})
		}
	case ResourceTypeCertificateSelfSigned:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				certs, err := nc.GetCertificates()
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, cert := range certs.Results {
					res = append(res, interface{}(cert))
				}
				return res, nil
			})
		}
	case ResourceTypeSpoofGuardSwitchingProfile:
		{
			return ResourceCollectFunc(func() ([]interface{}, error) {
				switchingProfiles, err := nc.ListSwitchingProfilesByType(ResourceTypeSpoofGuardSwitchingProfile)
				if err != nil {
					return nil, err
				}
				var res []interface{}
				for _, switchingProfile := range switchingProfiles.Results {
					res = append(res, interface{}(switchingProfile))
				}
				return res, nil
			})
		}
	default:
		nc.Debug(fmt.Sprintf("ResourceCollectFunc(): unrecognized resource type: %s\n", resourceType))
	}
	// no default is set
	return nil
}

func (nc *client) ResourceDeleteFunc(resourceType string) ResourceDeleteFunc {
	switch resourceType {
	case ResourceTypePersistenceProfile:
		{
			return ResourceDeleteFunc(nc.DeleteLoadBalancerPersistenceProfile)
		}
	case ResourceTypeFirewallSection:
		{
			return ResourceDeleteFunc(nc.DeleteFirewallSection)
		}
	case ResourceTypeNSGroup:
		{
			return ResourceDeleteFunc(nc.DeleteNSGroup)
		}
	case ResourceTypeIPSet:
		{
			return ResourceDeleteFunc(nc.DeleteIPSet)
		}
	case ResourceTypeLbService:
		{
			return ResourceDeleteFunc(nc.DeleteLoadBalancerService)
		}
	case ResourceTypeLbRule:
		{
			return ResourceDeleteFunc(nc.DeleteLoadBalancerRule)
		}
	case ResourceTypeLbPool:
		{
			return ResourceDeleteFunc(nc.DeleteLoadBalancerPool)
		}
	case ResourceTypeLogicalPort:
		{
			return ResourceDeleteFunc(nc.DeleteLogicalPort)
		}
	case ResourceTypeLbVirtualServer:
		{
			return ResourceDeleteFunc(nc.DeleteLoadBalancerVirtualServer)
		}
	case ResourceTypeLogicalRouter:
		{
			return ResourceDeleteFunc(nc.DeleteLogicalRouter)
		}
	case ResourceTypeLogicalRouterPort:
		{
			return ResourceDeleteFunc(nc.DeleteLogicalRouterPort)
		}
	case ResourceTypeLogicalSwitch:
		{
			return ResourceDeleteFunc(nc.DeleteLogicalSwitch)
		}
	case ResourceTypePrincipalIdentity:
		{
			return ResourceDeleteFunc(nc.DeletePrincipalIdentity)
		}
	case ResourceTypeCertificateSelfSigned:
		{
			return ResourceDeleteFunc(nc.DeleteCertificate)
		}
	case ResourceTypeSpoofGuardSwitchingProfile:
		{
			return ResourceDeleteFunc(nc.DeleteSwitchingProfile)
		}
	default:
		nc.Debug(fmt.Sprintf("ResourceDeleteFunc(): unrecognized resource type: %s\n", resourceType))
	}
	// no default is set
	return nil
}
