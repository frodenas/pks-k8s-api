/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"fmt"

	"github.com/vmware/govmomi/object"
)

// GetFolder gets the folder object given a folder path.
func (vc *client) GetFolder(folderPath string) (*object.Folder, error) {
	log.Info(fmt.Sprintf("Getting Folder `%s`", folderPath))

	folder, err := vc.finder.Folder(vc.context, folderPath)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Folder `%s`", folderPath))
		return nil, err
	}

	return folder, nil
}
