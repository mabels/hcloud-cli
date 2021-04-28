// Code generated by interfacer; DO NOT EDIT

package hcapi2

import (
	"context"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

// FirewallClientBase is an interface generated for "github.com/hetznercloud/hcloud-go/hcloud.FirewallClient".
type FirewallClientBase interface {
	All(context.Context) ([]*hcloud.Firewall, error)
	AllWithOpts(context.Context, hcloud.FirewallListOpts) ([]*hcloud.Firewall, error)
	ApplyResources(context.Context, *hcloud.Firewall, []hcloud.FirewallResource) ([]*hcloud.Action, *hcloud.Response, error)
	Create(context.Context, hcloud.FirewallCreateOpts) (hcloud.FirewallCreateResult, *hcloud.Response, error)
	Delete(context.Context, *hcloud.Firewall) (*hcloud.Response, error)
	Get(context.Context, string) (*hcloud.Firewall, *hcloud.Response, error)
	GetByID(context.Context, int) (*hcloud.Firewall, *hcloud.Response, error)
	GetByName(context.Context, string) (*hcloud.Firewall, *hcloud.Response, error)
	List(context.Context, hcloud.FirewallListOpts) ([]*hcloud.Firewall, *hcloud.Response, error)
	RemoveResources(context.Context, *hcloud.Firewall, []hcloud.FirewallResource) ([]*hcloud.Action, *hcloud.Response, error)
	SetRules(context.Context, *hcloud.Firewall, hcloud.FirewallSetRulesOpts) ([]*hcloud.Action, *hcloud.Response, error)
	Update(context.Context, *hcloud.Firewall, hcloud.FirewallUpdateOpts) (*hcloud.Firewall, *hcloud.Response, error)
}