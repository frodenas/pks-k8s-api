/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// EvaluateTag searches for a tag with scope matching the scope parameter
// and returns a pointer to the value of the first matching tag
// or nil if a match is not found
func EvaluateTag(resource *models.ManagedResource, targetScope string) *string {
	for _, tag := range resource.Tags {
		if tag.Scope == targetScope {
			return util.StringPtr(tag.Tag)
		}
	}
	return nil
}

// TagExist returns true if the target tag exist in the list
func TagExist(tagList []*models.Tag, targetTags []*models.Tag) bool {
	for _, tag := range tagList {
		for _, targetTag := range targetTags {
			if tag.Scope == targetTag.Scope && tag.Tag == targetTag.Tag {
				return true
			}
		}
	}
	return false
}

// ScopeExist returns true if the target tag exist in the list
func ScopeExist(resource *models.ManagedResource, targetScope string) bool {
	for _, tag := range resource.Tags {
		if tag.Scope == targetScope {
			return true
		}
	}
	return false
}

// RemoveTag removes a tag from a ManagedResource if it has that tag
func RemoveTag(resource *models.ManagedResource, target models.Tag) {
	var newTags []*models.Tag
	for _, tag := range resource.Tags {
		if tag.Scope != target.Scope || tag.Tag != target.Tag {
			newTags = append(newTags, tag)
		}
	}
	resource.Tags = newTags
}

// ValidateTags performs various checks on existing resource related to tagging
func ValidateTags(managedResource models.ManagedResource, tags []*models.Tag) error {
	if len(managedResource.Tags)+len(tags) > NsxMaxTagsAllowed {
		err := fmt.Errorf("Adding new tags to %s: %s exceed maximum number of tags", managedResource.ResourceType, managedResource.ID)
		return err
	}
	if TagExist(managedResource.Tags, tags) {
		err := fmt.Errorf("Tag already exist in %s: %s", managedResource.ResourceType, managedResource.ID)
		return err
	}
	return nil
}

// for given cluster, ncp created resource should satisfy following requirements:
// 1. has tag ncp/cluster
// 2. value of ncp/cluster tag should match provided clusterName
// should be in sync with: https://opengrok.eng.vmware.com/source/xref/nsx-ujo.git/sample_scripts/nsx_cleanup.py#244
func IsNcpResource(resource *models.ManagedResource, clusterName string) bool {
	c := EvaluateTag(resource, NcpTagKeyCluster)
	return c != nil && *c == clusterName
}

// ncp shared resource should satisfy following requirements:
// 1. has tag ncp/shared_resource
// 2. value of tag ncp/shared_resource is "true"
func IsNcpSharedResource(resource *models.ManagedResource) bool {
	c := EvaluateTag(resource, NcpTagKeySharedResource)
	return c != nil && *c == "true"
}
