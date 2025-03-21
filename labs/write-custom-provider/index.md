# Custom Terraform Provider Development Lab

## Overview
You will create a custom Terraform provider in this lab that manages AWS S3 buckets. You'll learn how Terraform providers work and implement basic CRUD (Create, Read, Update, Delete) operations using the Go programming language.

## Prerequisites
- Go 1.21+ installed and configured
- Terraform v1.15+
- AWS credentials (access key and secret key)
- Basic understanding of Terraform (no Go experience required)

## Environment Setup

### Windows Setup
1. Install Go:
   
   Use chocolatey to install Go
   
   ```powershell
   choco install -y golang
   ```

   Confirm it was installed successfully.
   
   ```powershell
   go version
   ```
   
   

## Lab Setup

### 1. Create Development Directory

Windows:
```powershell
mkdir custom-tf-provider
cd custom-tf-provider
```



### 2. Configure Terraform Development Overrides

Create a Terraform configuration file:

Windows:
```powershell
New-Item -Path "$env:APPDATA\terraform.rc" -ItemType File
```

Add the following content:
```hcl
provider_installation {
  dev_overrides {
      "hashicorp.com/edu/custom-s3" = "C:/Users/TEKstudent/go/bin"  # Windows path
  }
  direct {}
}
```



### 3. Initialize Go Module

```sh
go mod init terraform-provider-custom-s3
```

This creates a `go.mod` file to manage dependencies.

## Provider Structure

Create the following directory structure:

Windows:
```powershell
mkdir internal
mkdir internal\provider
```

The provider will have this structure:
```
custom-tf-provider/
├── go.mod
├── main.go
└── internal/
    └── provider/
        ├── provider.go
        ├── bucket_data_source.go
        └── bucket_resource.go
```

### 1. Create the Main Entry Point

Create `main.go` in the root directory with the following content:

```go
package main

import (
    "context"
    "flag"
    "log"
    "terraform-provider-custom-s3/internal/provider"
    "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var version string = "dev"

func main() {
    var debug bool

    flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers")
    flag.Parse()

    opts := providerserver.ServeOpts{
        Address: "hashicorp.com/edu/custom-s3",
        Debug:   debug,
    }

    err := providerserver.Serve(context.Background(), provider.New(version), opts)
    if err != nil {
        log.Fatal(err.Error())
    }
}
```

This file serves several purposes:
- Entry point for the provider plugin
- Sets up the provider server process
- Configures debugging options
- Defines the provider's registry address
- Initializes error handling

### 2. Install Dependencies

Run:
```sh
go mod tidy
```

This will download required dependencies and update `go.mod`.

### 3. Implement the Provider

Create `internal/provider/provider.go`:

```go
package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ provider.Provider = &cs3Provider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &cs3Provider{
			version: version,
		}
	}
}

// cs3Provider is the provider implementation.
type cs3Provider struct {
	version string
}

// cs3ProviderModel maps provider schema data to a Go type.
type cs3ProviderModel struct {
	Region    types.String `tfsdk:"region"`
	AccessKey types.String `tfsdk:"access_key"`
	SecretKey types.String `tfsdk:"secret_key"`
}

// Metadata returns the provider type name.
func (p *cs3Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "customs3"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *cs3Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"region": schema.StringAttribute{
				Optional: true,
			},
			"access_key": schema.StringAttribute{
				Optional: true,
			},
			"secret_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *cs3Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// Retrieve provider data from configuration
	var config cs3ProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Region.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("region"),
			"Unknown Region",
			"The provider cannot create the Custom S3 client as there is an unknown configuration value for the AWS Region. ",
		)
	}
	if config.AccessKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_key"),
			"Unknown Access Key value",
			"The provider cannot create the Custom S3 client as there is an unknown configuration value for the AWS Access Key. ",
		)
	}
	if config.SecretKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("secret_key"),
			"Unknown Secret Key value",
			"The provider cannot create the Custom S3 client as there is an unknown configuration value for the AWS Secret Key. ",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	region := os.Getenv("AWS_REGION")
	access_key := os.Getenv("AWS_ACCESS_KEY_ID")
	secret_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if !config.Region.IsNull() {
		region = config.Region.ValueString()
	}

	if !config.AccessKey.IsNull() {
		access_key = config.AccessKey.ValueString()
	}

	if !config.SecretKey.IsNull() {
		secret_key = config.SecretKey.ValueString()
	}

	if region == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("region"),
			"Missing Region",
			"The provider cannot create the AWS client as there is a missing or empty value for the Region. ",
		)
	}

	if access_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_key"),
			"Missing Access Key",
			"The provider cannot create the AWS client as there is a missing or empty value for the Access Key. ",
		)
	}

	if secret_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("secret_key"),
			"Missing Secret Key",
			"The provider cannot create the AWS client as there is a missing or empty value for the Secret Key. ",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	// Create AWS client
	client, err := session.NewSession(&aws.Config{
		Region:      aws.String(region), // Specify the AWS region
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *cs3Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewBucketDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *cs3Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrderResource,
	}
}

```

This file handles:
- Provider configuration schema
- AWS authentication
- Session management
- Error handling and diagnostics


### Create the resources files
`internal\provider\bucket_data_source.go`

```go
package provider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &bucketDataSource{}
	_ datasource.DataSourceWithConfigure = &bucketDataSource{}
)

// NewBucketDataSource is a helper function to simplify the provider implementation.
func NewBucketDataSource() datasource.DataSource {
	return &bucketDataSource{}
}

// bucketDataSource is the data source implementation.
type bucketDataSource struct {
	client *session.Session
}

// bucketDataSourceModel maps the data source schema data.
type bucketDataSourceModel struct {
	Buckets []bucketModel `tfsdk:"buckets"`
}

// bucketModel maps coffees schema data.
type bucketModel struct {
	Date tftypes.String `tfsdk:"date"`
	Name tftypes.String `tfsdk:"name"`
	Tags tftypes.String `tfsdk:"tags"`
}

func (d *bucketDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_buckets"
}

func (d *bucketDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"buckets": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"date": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"tags": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (d *bucketDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state bucketDataSourceModel

	svc := s3.New(d.client)

	buckets, err := svc.ListBuckets(nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Bucket data",
			err.Error(),
		)
		return
	}

	for _, bucket := range buckets.Buckets {
		bucketState := bucketModel{
			Date: types.StringValue(bucket.CreationDate.Format("2006-01-02 15:04:05")),
			Name: types.StringValue(*bucket.Name),
		}
		state.Buckets = append(state.Buckets, bucketState)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *bucketDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*session.Session)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *session.Session, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

```

`internal\provider\bucket_resource.go`

```go
package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &orderResource{}
	_ resource.ResourceWithConfigure = &orderResource{}
)

// NewOrderResource is a helper function to simplify the provider implementation.
func NewOrderResource() resource.Resource {
	return &orderResource{}
}

// orderResource is the resource implementation.
type orderResource struct {
	client *session.Session
}

// orderResourceModel maps the resource schema data.
type orderResourceModel struct {
	ID          tftypes.String   `tfsdk:"id"`
	Buckets     []orderItemModel `tfsdk:"buckets"`
	LastUpdated tftypes.String   `tfsdk:"last_updated"`
}

// orderItemModel maps order item data.
type orderItemModel struct {
	Date tftypes.String `tfsdk:"date"`
	Name tftypes.String `tfsdk:"name"`
	Tags tftypes.String `tfsdk:"tags"`
}

// Metadata returns the resource type name.
func (r *orderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_s3_bucket"
}

// Schema defines the schema for the resource.
func (r *orderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"buckets": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"date": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
						"tags": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

// Create a new resource.
func (r *orderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan orderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = types.StringValue(strconv.Itoa(1))
	for index, item := range plan.Buckets {
		// Create an S3 service client
		svc := s3.New(r.client)
		awsStringBucket := strings.Replace(item.Name.String(), "\"", "", -1)

		// Create input parameters for the CreateBucket operation
		input := &s3.CreateBucketInput{
			Bucket: aws.String(awsStringBucket),
		}

		// Execute the CreateBucket operation
		_, err := svc.CreateBucket(input)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating order",
				"Could not create order, unexpected error: "+err.Error(),
			)
			return
		}

		// Add tags
		var tags []*s3.Tag
		tagValue := strings.Replace(item.Tags.String(), "\"", "", -1)
		tags = append(tags, &s3.Tag{
			Key:   aws.String("tfkey"),
			Value: aws.String(tagValue),
		})

		_, err = svc.PutBucketTagging(&s3.PutBucketTaggingInput{
			Bucket: aws.String(awsStringBucket),
			Tagging: &s3.Tagging{
				TagSet: tags,
			},
		})
		if err != nil {
			fmt.Println("Error adding tags to the bucket:", err)
			return
		}

		fmt.Printf("Bucket %s created successfully\n", item.Name)

		plan.Buckets[index] = orderItemModel{
			Name: types.StringValue(awsStringBucket),
			Date: types.StringValue(time.Now().Format(time.RFC850)),
			Tags: types.StringValue(tagValue),
		}
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *orderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state orderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, item := range state.Buckets {
		awsStringBucket := strings.Replace(item.Name.String(), "\"", "", -1)

		svc := s3.New(r.client)
		params := &s3.HeadBucketInput{
			Bucket: aws.String(awsStringBucket),
		}

		_, err := svc.HeadBucket(params)
		if err != nil {
			fmt.Println("Error getting bucket information:", err)
			os.Exit(1)
		}
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan orderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.ID = types.StringValue(strconv.Itoa(1))

	for index, item := range plan.Buckets {
		// Create an S3 service client
		svc := s3.New(r.client)
		awsStringBucket := strings.Replace(item.Name.String(), "\"", "", -1)

		// Add tags
		var tags []*s3.Tag
		tagValue := strings.Replace(item.Tags.String(), "\"", "", -1)
		tags = append(tags, &s3.Tag{
			Key:   aws.String("tfkey"),
			Value: aws.String(tagValue),
		})

		_, err := svc.PutBucketTagging(&s3.PutBucketTaggingInput{
			Bucket: aws.String(awsStringBucket),
			Tagging: &s3.Tagging{
				TagSet: tags,
			},
		})
		if err != nil {
			fmt.Println("Error adding tags to the bucket:", err)
			return
		}

		plan.Buckets[index] = orderItemModel{
			Name: types.StringValue(strings.Replace(awsStringBucket, "\"", "", -1)),
			Date: types.StringValue(time.Now().Format(time.RFC850)),
			Tags: types.StringValue(strings.Replace(tagValue, "\"", "", -1)),
		}
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state orderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, item := range state.Buckets {
		svc := s3.New(r.client)

		input := &s3.DeleteBucketInput{
			Bucket: aws.String(strings.Replace(item.Name.String(), "\"", "", -1)),
		}

		_, err := svc.DeleteBucket(input)
		if err != nil {
			log.Fatalf("failed to delete bucket, %v", err)
		}

	}
}

// Configure adds the provider configured client to the resource.
func (r *orderResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*session.Session)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *session.Session, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

```

### 4. Build and Install

Windows:
```powershell
go build -o terraform-provider-custom-s3.exe
move terraform-provider-custom-s3.exe %USERPROFILE%\go\bin\
```

## Testing the Provider

### 1. Create Test Directory

Create a new directory for testing:

Windows:
```powershell
mkdir test
cd test
```



### 2. Create Test Configuration

Create `main.tf`:
```hcl
terraform {
  required_providers {
    customs3 = {
      source = "hashicorp.com/edu/custom-s3"
    }
  }
}

provider "customs3" {
  region     = "us-west-1"  # or your preferred region
  access_key = "<YOUR_ACCESS_KEY>"
  secret_key = "<YOUR_SECRET_KEY>"
}

resource "customs3_s3_bucket" "test" {
  buckets = [{
    name = "test-bucket-2398756"
    tags = "yourbucket"
  }]
}
```

### 3. Plan and Apply

```sh
terraform plan
terraform apply
```

### 4. Verify

1. Check AWS Console to verify bucket creation
2. List buckets using AWS CLI:
   ```sh
   aws s3 ls
   ```

### 5. Clean Up

```sh
terraform destroy
```

## Troubleshooting

### Common Windows Issues

1. Path Issues:
   - Ensure `%USERPROFILE%\go\bin` is in your PATH
   - Use `echo %PATH%` to verify

2. Permission Issues:
   - Run PowerShell as Administrator when needed
   - Check file permissions in `%USERPROFILE%\go\bin`

3. Go Build Errors:
   - Clear Go build cache: `go clean -cache`
   - Ensure all dependencies are installed: `go mod tidy`



## Congrats!
