package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// myServiceProvider is the main struct that represents the provider.
// It implements the Terraform Provider interface.
type myServiceProvider struct{}

// New is a factory pattern and creates a new
// instance of the provider.
// Terraform calls this when loading the provider binary.
func New() provider.Provider {
	return &myServiceProvider{}
}

// Metadata sets the provider's name that Terraform will use in configuration files.
func (p *myServiceProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "myservice"
}

// Schema defines the provider-level configuration schema.
// In this example, we allow setting an optional `api_base_url` in the Terraform config.
func (p *myServiceProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_base_url": schema.StringAttribute{
				Optional:    true,
				Description: "Base URL of the item service. Defaults to http://localhost:8080.",
			},
		},
	}
}

// Configure is called once when the provider is initialized.
// It reads the provider configuration and prepares data
// that can be passed to resources.
func (p *myServiceProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Define a temporary struct to hold the configuration values
	var config struct {
		ApiBaseUrl types.String `tfsdk:"api_base_url"`
	}

	// Populate the struct with values from Terraform config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set default if api_base_url was not provided
	apiBaseURL := "http://localhost:8080"
	if !config.ApiBaseUrl.IsNull() && !config.ApiBaseUrl.IsUnknown() {
		apiBaseURL = config.ApiBaseUrl.ValueString()
	}

	// Store this value so resources can access it later through req.ProviderData
	resp.DataSourceData = apiBaseURL
	resp.ResourceData = apiBaseURL
}

// Resources returns a list of resource types the provider supports.
// Here we register one resource type: myservice_item.
func (p *myServiceProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewItemResource, // This refers to your CRUD resource, implemented in resource.go
	}
}

// DataSources returns a list of data sources the provider supports.
// We have none in this example, so we return nil.
func (p *myServiceProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
