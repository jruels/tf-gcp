package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ItemResource defines the Terraform resource implementation.
// It wraps the CRUD operations Terraform will call.
type ItemResource struct{}

// This line ensures that ItemResource satisfies the resource.Resource interface.
var _ resource.Resource = &ItemResource{}

// NewItemResource creates a new instance of the resource.
func NewItemResource() resource.Resource {
	return &ItemResource{}
}

// itemModel maps the Terraform schema to a Go struct.
type itemModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Metadata provides the Terraform type name for this resource.
func (r *ItemResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "myservice_item"
}

// Schema defines the HCL attributes this resource supports.
func (r *ItemResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true, // Terraform sets this automatically after creation
				Description: "Unique identifier of the item",
			},
			"name": schema.StringAttribute{
				Required:    true, // Required in the HCL config
				Description: "Name of the item",
			},
		},
	}
}

// Create is called when Terraform applies a new resource.
// It POSTs to the service, saves the new ID, and updates the Terraform state.
func (r *ItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan itemModel
	diags := req.Plan.Get(ctx, &plan) // Get desired state from Terraform config
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build request payload
	payload := map[string]string{"name": plan.Name.ValueString()}
	body, _ := json.Marshal(payload)

	// Send POST request to the API
	httpResp, err := http.Post("http://localhost:8080/items", "application/json", bytes.NewBuffer(body))
	if err != nil {
		resp.Diagnostics.AddError("POST request failed", err.Error())
		return
	}
	defer httpResp.Body.Close()

	// Parse response
	respBody, _ := io.ReadAll(httpResp.Body)
	var created map[string]string
	json.Unmarshal(respBody, &created)

	// Set ID in the Terraform state
	plan.ID = types.StringValue(created["id"])

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read is called to refresh Terraform state from the actual remote resource.
func (r *ItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state itemModel
	diags := req.State.Get(ctx, &state) // Get current state from Terraform
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build GET request to fetch item from API
	url := fmt.Sprintf("http://localhost:8080/items/%s", state.ID.ValueString())
	httpResp, err := http.Get(url)
	if err != nil {
		resp.Diagnostics.AddError("GET request failed", err.Error())
		return
	}
	defer httpResp.Body.Close()

	// If item not found, remove it from Terraform state
	if httpResp.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}

	// Parse response
	respBody, _ := io.ReadAll(httpResp.Body)
	var item map[string]string
	json.Unmarshal(respBody, &item)

	// Update Terraform state
	state.Name = types.StringValue(item["name"])
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Update is called when Terraform detects a change to the resource configuration.
func (r *ItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan itemModel
	diags := req.Plan.Get(ctx, &plan) // Get new desired state
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build PUT request payload
	url := fmt.Sprintf("http://localhost:8080/items/%s", plan.ID.ValueString())
	payload := map[string]string{"name": plan.Name.ValueString()}
	body, _ := json.Marshal(payload)

	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		resp.Diagnostics.AddError("PUT request creation failed", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(request)
	if err != nil {
		resp.Diagnostics.AddError("PUT request failed", err.Error())
		return
	}
	defer httpResp.Body.Close()

	// Parse response
	respBody, _ := io.ReadAll(httpResp.Body)
	var updated map[string]string
	json.Unmarshal(respBody, &updated)

	// Update Terraform state
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete is called when Terraform destroys the resource.
func (r *ItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state itemModel
	diags := req.State.Get(ctx, &state) // Get current state
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build DELETE request
	url := fmt.Sprintf("http://localhost:8080/items/%s", state.ID.ValueString())
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		resp.Diagnostics.AddError("DELETE request creation failed", err.Error())
		return
	}

	client := &http.Client{}
	httpResp, err := client.Do(request)
	if err != nil {
		resp.Diagnostics.AddError("DELETE request failed", err.Error())
		return
	}
	defer httpResp.Body.Close()

	// Remove from Terraform state
	resp.State.RemoveResource(ctx)
}
