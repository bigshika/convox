package k8s

import (
	"fmt"

	"github.com/convox/convox/pkg/structs"
	am "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p *Provider) ResourceGet(app, name string) (*structs.Resource, error) {
	d, err := p.Cluster.AppsV1().Deployments(p.AppNamespace(app)).Get(fmt.Sprintf("resource-%s", name), am.GetOptions{})
	if err != nil {
		return nil, err
	}

	cm, err := p.Cluster.CoreV1().ConfigMaps(p.AppNamespace(app)).Get(fmt.Sprintf("resource-%s", name), am.GetOptions{})
	if err != nil {
		return nil, err
	}

	status := "running"

	if d.Status.ReadyReplicas < 1 {
		status = "pending"
	}

	r := &structs.Resource{
		Name:   name,
		Status: status,
		Type:   d.ObjectMeta.Labels["kind"],
		Url:    cm.Data["URL"],
	}

	return r, nil
}

func (p *Provider) ResourceList(app string) (structs.Resources, error) {
	lopts := am.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s,type=resource", app),
	}

	ds, err := p.Cluster.AppsV1().Deployments(p.AppNamespace(app)).List(lopts)
	if err != nil {
		return nil, err
	}

	rs := structs.Resources{}

	for _, d := range ds.Items {
		r, err := p.ResourceGet(app, d.ObjectMeta.Labels["resource"])
		if err != nil {
			return nil, err
		}

		rs = append(rs, *r)
	}

	return rs, nil
}

func (p *Provider) SystemResourceCreate(kind string, opts structs.ResourceCreateOptions) (*structs.Resource, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) SystemResourceDelete(name string) error {
	return fmt.Errorf("unimplemented")
}

func (p *Provider) SystemResourceGet(name string) (*structs.Resource, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) SystemResourceLink(name, app string) (*structs.Resource, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) SystemResourceList() (structs.Resources, error) {
	return structs.Resources{}, nil
}

func (p *Provider) SystemResourceTypes() (structs.ResourceTypes, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) SystemResourceUnlink(name, app string) (*structs.Resource, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (p *Provider) SystemResourceUpdate(name string, opts structs.ResourceUpdateOptions) (*structs.Resource, error) {
	return nil, fmt.Errorf("unimplemented")
}
