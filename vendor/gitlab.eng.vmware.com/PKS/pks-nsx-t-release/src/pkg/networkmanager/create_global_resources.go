package networkmanager

import (
	"github.com/Sirupsen/logrus"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/workflow"
)

func (n *networkManager) CreateGlobalResources() error {
	if err := n.createNCPGlobalSwitchingProfile(); err != nil {
		return err
	}
	return nil
}

func (n *networkManager) createNCPGlobalSwitchingProfile() error {
	switchingProfile, err := n.np.GetSwitchingProfileByTag(models.SwitchingProfileTypeIDEntryKeySpoofGuardSwitchingProfile, nsx.NcpTagKeyOwner, nsx.NcpTagValueOwner)
	if err != nil {
		return err
	}

	if switchingProfile == nil {
		createGlobalSwitchingProfileResp := &CreateSwitchingProfileResp{}
		createGlobalSwitchingProfileRequest := &CreateSwitchingProfileRequest{
			Name:        "Container Spoofguard",
			Description: "Container Spoofguard",
			Tags: []*models.Tag{
				{
					Scope: nsx.NcpTagKeyOwner,
					Tag:   nsx.NcpTagValueOwner,
				},
			},
			WhiteListProviders: []string{nsx.SpoofGuardPortBindings},
		}

		createGlobalSwitchingProfile := n.NewCreateSwitchingProfile(createGlobalSwitchingProfileRequest, createGlobalSwitchingProfileResp, "CreateSwitchingProfile")
		if err := createGlobalSwitchingProfile.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (n *networkManager) NewCreateSwitchingProfile(req *CreateSwitchingProfileRequest, resp *CreateSwitchingProfileResp, logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.createSwitchingProfile(req, resp, n.log.WithField(projName, logField))
	})
}

func (n *networkManager) createSwitchingProfile(req *CreateSwitchingProfileRequest, resp *CreateSwitchingProfileResp, log logrus.FieldLogger) error {
	switchingProfileID, err := n.np.CreateSpoofGuardSwitchingProfile(req.Name, req.Description, req.Tags, req.WhiteListProviders)
	if err != nil {
		return err
	}
	resp.SwitchingProfileID = switchingProfileID

	return nil
}

type CreateSwitchingProfileRequest struct {
	Name               string
	Description        string
	Tags               []*models.Tag
	WhiteListProviders []string
}

type CreateSwitchingProfileResp struct {
	SwitchingProfileID string
}
