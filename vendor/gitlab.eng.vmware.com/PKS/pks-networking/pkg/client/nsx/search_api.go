/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/search"
)

const (
	resourceSpecificQuery    = "resource_type:%s AND tags.scope:\"%s\" AND tags.tag:\"%s\""
	nonResourceSpecificQuery = "tags.scope:\"%s\" AND tags.tag:\"%s\""
)

// SearchByTag takes an optional resourceType and a Tag and returns a matching list of resources.
func (nc *client) SearchByTag(resourceType *string, tag models.Tag) (*models.SearchResults, error) {
	query := buildQuery(resourceType, tag)

	params := search.NewSearchByTagParams().WithQuery(query)

	res, err := nc.client.Search.SearchByTag(params, nc.auth)

	if err != nil {
		return nil, err
	}

	numResultsLeft := res.Payload.ResultCount - int64(len(res.Payload.Results))
	cursor := res.Payload.Cursor

	for numResultsLeft > 0 {
		params = search.NewSearchByTagParams().WithQuery(query).WithCursor(&cursor)
		nextpage, err := nc.client.Search.SearchByTag(params, nc.auth)
		if err != nil {
			return nil, err
		}
		res.Payload.Results = append(res.Payload.Results, nextpage.Payload.Results...)
		cursor = nextpage.Payload.Cursor
		numResultsLeft -= int64(len(nextpage.Payload.Results))
	}

	return res.Payload, nil
}

// buildQuery assembles a search query
func buildQuery(resourceType *string, tag models.Tag) string {
	if resourceType == nil {
		return fmt.Sprintf(nonResourceSpecificQuery, tag.Scope, tag.Tag)
	}
	return fmt.Sprintf(resourceSpecificQuery, *resourceType, tag.Scope, tag.Tag)
}
