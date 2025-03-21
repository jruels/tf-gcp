package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure resource implementation satisfies required interfaces
var _ resource.Resource = &GreetingResource{}

type GreetingResource struct{}

// NewGreetingResource creates a new GreetingResource
func NewGreetingResource() resource.Resource {
	return &GreetingResource{}
}

// Metadata returns the resource type name
func (r *GreetingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "greeting_message"
}

// Schema defines the structure of the resource
func (r *GreetingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A simple resource that displays 'Hello, world!'",
		Attributes: map[string]schema.Attribute{
			"message": schema.StringAttribute{
				Description: "The greeting message",
				Computed:    true,
			},
		},
	}
}

// Create sets the value of the greeting message
func (r *GreetingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.State.Set(ctx, &struct {
		Message types.String `tfsdk:"message"`
	}{
		Message: types.StringValue("Hello, world!"),
	})
}

// Read does nothing since the message never changes
func (r *GreetingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update does nothing since the message never changes
func (r *GreetingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete does nothing since there is nothing to clean up
func (r *GreetingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
