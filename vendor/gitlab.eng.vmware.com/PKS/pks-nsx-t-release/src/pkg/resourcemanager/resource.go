package resourcemanager

import (
	"errors"
	"os"

	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/printer"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

var AdaptedIsNcpResource ResourceFilterFunc = ResourceFilterFunc(func(i interface{}, args ...string) (bool, error) {
	if len(args) != 1 {
		return false, errors.New("AdaptedIsNcpResource(): requires only one parameter")
	}
	m, err := nsx.GetManagedResource(i)
	if err != nil {
		return false, err
	}
	return nsx.IsNcpResource(m, args[0]), nil
})

var AdaptedIsNcpSharedResource ResourceFilterFunc = ResourceFilterFunc(func(i interface{}, args ...string) (bool, error) {
	m, err := nsx.GetManagedResource(i)
	if err != nil {
		return false, err
	}
	return nsx.IsNcpSharedResource(m), nil
})

var NotNcpExternalResource ResourceFilterFunc = ResourceFilterFunc(func(i interface{}, args ...string) (bool, error) {
	m, err := nsx.GetManagedResource(i)
	if err != nil {
		return true, err
	}
	external := nsx.EvaluateTag(m, nsx.NcpTagKeyExternal)
	if external != nil && util.StringVal(external) == "true" {
		return false, nil
	}
	return true, nil
})

var NameEquals ResourceFilterFunc = ResourceFilterFunc(func(i interface{}, args ...string) (bool, error) {
	if len(args) != 1 {
		return false, errors.New("NameEquals(): requires only one parameter")
	}
	m, err := nsx.GetManagedResource(i)
	if err != nil {
		return true, err
	}
	if m.DisplayName != args[0] {
		return false, nil
	}
	return true, nil
})

type cleanupResource struct {
	*printer.Printer
	nsx.Client
	resourceType string
	readOnly     bool
	collect      nsx.ResourceCollectFunc
	filter       []ResourceFilterFunc
	delete       nsx.ResourceDeleteFunc
	predelete    ResourcePreDeleteFunc
	afterdelete  ResourceAfterDeleteFunc
	prev         *cleanupResource
}

// Resource is used to do all kinds of operations on one specific resource
// Resource needs to be bound to one nsxclient as its runner
func NewResource(nsxclient nsx.Client) *cleanupResource {
	res := &cleanupResource{
		Client: nsxclient,
	}
	// by default, print to stderr
	res.Printer = printer.New(os.Stderr)
	return res
}

func (c *cleanupResource) GetResourceType() string {
	return c.resourceType
}

func (c *cleanupResource) SetResourceType(resourceType string) Resource {
	if c.resourceType != "" {
		cur := &cleanupResource{
			Client:   c.Client,
			readOnly: c.readOnly,
			Printer:  c.Printer,
		}
		cur.prev = c
		c = cur
	}
	c.resourceType = resourceType
	// if collect func is not defined, try to look it up in default registry
	if c.collect == nil {
		c.collect = c.ResourceCollectFunc(resourceType)
	}
	// if delete func is not defined, try to look it up in default registry
	if c.delete == nil {
		c.delete = c.ResourceDeleteFunc(resourceType)
	}
	return c
}

func (c *cleanupResource) SetReadOnly(readOnly bool) Resource {
	c.readOnly = readOnly
	return c
}

func (c *cleanupResource) CollectBy(f nsx.ResourceCollectFunc) Resource {
	c.collect = f
	return c
}

func (c *cleanupResource) PreDeleteBy(f ResourcePreDeleteFunc) Resource {
	c.predelete = f
	return c
}

func (c *cleanupResource) DeleteBy(f nsx.ResourceDeleteFunc) Resource {
	c.delete = f
	return c
}

func (c *cleanupResource) AfterDeleteBy(f ResourceAfterDeleteFunc) Resource {
	c.afterdelete = f
	return c
}

func (c *cleanupResource) SetPrinter(p *printer.Printer) Resource {
	c.Printer = p
	return c
}

func (c *cleanupResource) FilterBy(f ResourceFilterFunc, args ...string) Resource {
	c.filter = append(c.filter, ResourceFilterFunc(func(i interface{}, a ...string) (bool, error) {
		return f(i, args...)
	}))
	return c
}

func (c *cleanupResource) GetCollection() ResourceCollection {
	out := make(chan interface{})
	go func() {
		defer close(out)
		if c.collect == nil {
			c.Warn("no default collect function is defined for %s, skip cleanup\n", c.resourceType)
			return
		}
		rs, err := c.collect()
		if err != nil {
			c.Error("failed to collect %s, skip cleanup\n", c.resourceType)
			return
		}
		for _, r := range rs {
			pass := true
			for _, filter := range c.filter {
				pass, err = filter(r)
				if err != nil {
					c.Error("error: %s\n", err.Error())
					pass = false
				}
				if !pass {
					break
				}
			}
			if !pass {
				continue
			}
			out <- r
		}
		return
	}()
	return ResourceCollection(out)
}

func (c *cleanupResource) Cleanup() error {
	for r := range c.GetCollection() {
		m, err := nsx.GetManagedResource(r)
		if err != nil {
			c.Error("error: %s\n", err.Error())
			continue
		}
		if c.predelete != nil {
			err = c.predelete(r)
			if err != nil {
				c.Error("error: %s\n", err.Error())
				continue
			}
		}
		if c.delete != nil {
			c.VerboseInfo("%s with ID:%s to be deleted\n", c.resourceType, m.ID)
			if !c.readOnly {
				err = c.delete(m.ID)
				if err != nil {
					c.Warn("failed to delete %s with ID(%s): %s", c.resourceType, m.ID, err.Error())
				} else {
					c.VerboseInfo("%s with ID:%s is deleted successfully\n", c.resourceType, m.ID)
				}
			}
		}
		if c.afterdelete != nil {
			err = c.afterdelete(r)
			if err != nil {
				c.Error("error: %s\n", err.Error())
				continue
			}
		}
	}
	if c.prev != nil {
		c = c.prev
	}
	return nil
}
