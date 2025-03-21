package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider implementation satisfies required interfaces
var _ provider.Provider = &GreetingProvider{}

type GreetingProvider struct{}

// New returns a new instance of the provider
func New() provider.Provider {
	return &GreetingProvider{}
}

// Metadata returns the provider metadata
func (p *GreetingProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "greeting"
}

// Schema defines the provider's schema (not needed for this simple provider)
func (p *GreetingProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Greeting provider that returns 'Hello, world!'",
	}
}

// Configure does nothing as there is no configuration needed
func (p *GreetingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources method is required but not needed for this provider
func (p *GreetingProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

// Resources registers the Greeting resource
func (p *GreetingProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGreetingResource,
	}
}
