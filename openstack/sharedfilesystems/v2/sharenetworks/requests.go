package sharenetworks

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToShareNetworkCreateMap() (map[string]any, error)
}

// CreateOpts contains options for creating a ShareNetwork. This object is
// passed to the sharenetworks.Create function. For more information about
// these parameters, see the ShareNetwork object.
type CreateOpts struct {
	// The UUID of the Neutron network to set up for share servers
	NeutronNetID string `json:"neutron_net_id,omitempty"`
	// The UUID of the Neutron subnet to set up for share servers
	NeutronSubnetID string `json:"neutron_subnet_id,omitempty"`
	// The UUID of the nova network to set up for share servers
	NovaNetID string `json:"nova_net_id,omitempty"`
	// The share network name
	Name string `json:"name"`
	// The share network description
	Description string `json:"description"`
}

// ToShareNetworkCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToShareNetworkCreateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "share_network")
}

// Create will create a new ShareNetwork based on the values in CreateOpts. To
// extract the ShareNetwork object from the response, call the Extract method
// on the CreateResult.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToShareNetworkCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(ctx, createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete will delete the existing ShareNetwork with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	resp, err := client.Delete(ctx, deleteURL(client, id), nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToShareNetworkListQuery() (string, error)
}

// ListOpts holds options for listing ShareNetworks. It is passed to the
// sharenetworks.List function.
type ListOpts struct {
	// admin-only option. Set it to true to see all tenant share networks.
	AllTenants bool `q:"all_tenants"`
	// The UUID of the project where the share network was created
	ProjectID string `q:"project_id"`
	// The neutron network ID
	NeutronNetID string `q:"neutron_net_id"`
	// The neutron subnet ID
	NeutronSubnetID string `q:"neutron_subnet_id"`
	// The nova network ID
	NovaNetID string `q:"nova_net_id"`
	// The network type. A valid value is VLAN, VXLAN, GRE or flat
	NetworkType string `q:"network_type"`
	// The Share Network name
	Name string `q:"name"`
	// The Share Network description
	Description string `q:"description"`
	// The Share Network IP version
	IPVersion gophercloud.IPVersion `q:"ip_version"`
	// The Share Network segmentation ID
	SegmentationID int `q:"segmentation_id"`
	// List all share networks created after the given date
	CreatedSince string `q:"created_since"`
	// List all share networks created before the given date
	CreatedBefore string `q:"created_before"`
	// Limit specifies the page size.
	Limit int `q:"limit"`
	// Limit specifies the page number.
	Offset int `q:"offset"`
}

// ToShareNetworkListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToShareNetworkListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListDetail returns ShareNetworks optionally limited by the conditions provided in ListOpts.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailURL(client)
	if opts != nil {
		query, err := opts.ToShareNetworkListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		p := ShareNetworkPage{pagination.MarkerPageBase{PageResult: r}}
		p.Owner = p
		return p
	})
}

// Get retrieves the ShareNetwork with the provided ID. To extract the ShareNetwork
// object from the response, call the Extract method on the GetResult.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(ctx, getURL(client, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToShareNetworkUpdateMap() (map[string]any, error)
}

// UpdateOpts contain options for updating an existing ShareNetwork. This object is passed
// to the sharenetworks.Update function. For more information about the parameters, see
// the ShareNetwork object.
type UpdateOpts struct {
	// The share network name
	Name *string `json:"name,omitempty"`
	// The share network description
	Description *string `json:"description,omitempty"`
	// The UUID of the Neutron network to set up for share servers
	NeutronNetID string `json:"neutron_net_id,omitempty"`
	// The UUID of the Neutron subnet to set up for share servers
	NeutronSubnetID string `json:"neutron_subnet_id,omitempty"`
	// The UUID of the nova network to set up for share servers
	NovaNetID string `json:"nova_net_id,omitempty"`
}

// ToShareNetworkUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToShareNetworkUpdateMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "share_network")
}

// Update will update the ShareNetwork with provided information. To extract the updated
// ShareNetwork from the response, call the Extract method on the UpdateResult.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToShareNetworkUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(ctx, updateURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// AddSecurityServiceOptsBuilder allows extensions to add additional parameters to the
// AddSecurityService request.
type AddSecurityServiceOptsBuilder interface {
	ToShareNetworkAddSecurityServiceMap() (map[string]any, error)
}

// AddSecurityServiceOpts contain options for adding a security service to an
// existing ShareNetwork. This object is passed to the sharenetworks.AddSecurityService
// function. For more information about the parameters, see the ShareNetwork object.
type AddSecurityServiceOpts struct {
	SecurityServiceID string `json:"security_service_id"`
}

// ToShareNetworkAddSecurityServiceMap assembles a request body based on the contents of an
// AddSecurityServiceOpts.
func (opts AddSecurityServiceOpts) ToShareNetworkAddSecurityServiceMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "add_security_service")
}

// AddSecurityService will add the security service to a ShareNetwork. To extract the updated
// ShareNetwork from the response, call the Extract method on the UpdateResult.
func AddSecurityService(ctx context.Context, client *gophercloud.ServiceClient, id string, opts AddSecurityServiceOptsBuilder) (r UpdateResult) {
	b, err := opts.ToShareNetworkAddSecurityServiceMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(ctx, addSecurityServiceURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// RemoveSecurityServiceOptsBuilder allows extensions to add additional parameters to the
// RemoveSecurityService request.
type RemoveSecurityServiceOptsBuilder interface {
	ToShareNetworkRemoveSecurityServiceMap() (map[string]any, error)
}

// RemoveSecurityServiceOpts contain options for removing a security service from an
// existing ShareNetwork. This object is passed to the sharenetworks.RemoveSecurityService
// function. For more information about the parameters, see the ShareNetwork object.
type RemoveSecurityServiceOpts struct {
	SecurityServiceID string `json:"security_service_id"`
}

// ToShareNetworkRemoveSecurityServiceMap assembles a request body based on the contents of an
// RemoveSecurityServiceOpts.
func (opts RemoveSecurityServiceOpts) ToShareNetworkRemoveSecurityServiceMap() (map[string]any, error) {
	return gophercloud.BuildRequestBody(opts, "remove_security_service")
}

// RemoveSecurityService will remove the security service from a ShareNetwork. To extract the updated
// ShareNetwork from the response, call the Extract method on the UpdateResult.
func RemoveSecurityService(ctx context.Context, client *gophercloud.ServiceClient, id string, opts RemoveSecurityServiceOptsBuilder) (r UpdateResult) {
	b, err := opts.ToShareNetworkRemoveSecurityServiceMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Post(ctx, removeSecurityServiceURL(client, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
