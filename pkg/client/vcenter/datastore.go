/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"fmt"

	"github.com/vmware/govmomi/object"
)

// GetDatastore gets the datastore object given a datastore path.
func (vc *client) GetDatastore(datastorePath string) (*object.Datastore, error) {
	log.Info(fmt.Sprintf("Getting Datastore `%s`", datastorePath))

	datastore, err := vc.finder.Datastore(vc.context, datastorePath)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Datastore `%s`", datastorePath))
		return nil, err
	}

	return datastore, nil
}
