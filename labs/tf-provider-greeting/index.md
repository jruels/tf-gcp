# Create a Hello World provider 





## Overview



This lab walks through the steps for creating a simple `Hello World` provider. It does not require any additional services.

## Environment Setup



### Windows Setup

If Go is not already installed, run the following: 

1. Install Go:

   Use chocolatey to install Go

   ```
   choco install -y golang
   ```

   

   Confirm it was installed successfully.

   ```
   go version
   ```

   

## Lab Setup



### 1. Create Development Directory



Windows:

* Create a folder named `custom-tf-providers` and open it in Visual Studio Code
  * Inside the `custom-tf-providers` folder, create a folder for each custom provider, for example: 
    * ​	`terraform-provider-greeting`



### 2. Configure Terraform Development Overrides



Create a Terraform configuration file:

Run the following in Visual Studio Code (using a PowerShell terminal)

```
New-Item -Path "$env:APPDATA\terraform.rc" -ItemType File
```



Add the following content:

```
provider_installation {
  dev_overrides {
      "hashicorp.com/edu/greeting" = "C:/Users/TEKstudent/go/bin"
  }
  direct {}
}
```



### 3. Initialize Go Module

Inside the `custom-tf-providers\terraform-provider-greeting` folder run: 

```
go mod init terraform-provider-greeting
```



This creates a `go.mod` file to manage dependencies.

## Provider Structure



Create the following directory structure:

* Inside VS Code, create `internal` and inside that create a sub-folder `provider`
* You should now have a folder path of `custom-tf-providers\terraform-provider-greeting\internal\provider`



The provider will have this structure:

```
custom-tf-provider/
├── go.mod
├── main.go
└── internal/
    └── provider/
        ├── provider.go
        ├── client.go
        ├── data_source.go
        └── resource.go
```



### 1. Create the Main Entry Point



Create `main.go` in the root directory with the following content:

```go
package main

import (
	"context"
	"flag"
	"log"

	"hello_provider/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Run the provider in debug mode")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.NewProvider, providerserver.ServeOpts{
		Address: "hashicorp.com/edu/greeting",
		Debug:   debug,
	})

	if err != nil {
		log.Fatal(err)
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

```
go mod tidy
```



This will download required dependencies and update `go.mod`.

### 3. Implement the Provider



Create `internal\provider\provider.go` with the following contents:

```
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GreetingProvider struct{}

type GreetingProviderModel struct {
	Person types.String `tfsdk:"person"`
}

// type Client struct {
// 	Name string
// }

func (p *GreetingProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "greeting"
	resp.Version = "1.0.0"
}

func (p *GreetingProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"person": schema.StringAttribute{
				Required:    true,
				Description: "The name of the person to greet.",
			},
		},
	}
}

func (p *GreetingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config GreetingProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	person := config.Person.ValueString()
	if person == "" {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"The 'person' attribute must be provided in the provider configuration.",
		)
		return
	}

	client := &Client{Name: person}
	resp.ResourceData = client
	resp.DataSourceData = client

	fmt.Printf("Provider configured for: Hello, %s!\n", person)
}

// Resources registers the resources provided by the provider.
func (p *GreetingProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGreetingResource,
	}
}

// DataSources registers the data sources provided by the provider.
func (p *GreetingProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGreetingDataSource,
	}
}

func NewProvider() provider.Provider {
	return &GreetingProvider{}
}
```



This file handles:

- Provider configuration schema
- Displaying resource data
- Error handling and diagnostics

### 4. Build and Install



Windows:

```
go build -o terraform-provider-greeting.exe
move terraform-provider-greeting.exe %USERPROFILE%\go\bin\
```



## Testing the Provider



### 1. Create Test Directory



Inside the `custom-tf-providers\terraform-greeting-provider` folder, create a `test` folder.



### 2. Create Test Configuration



Create `main.tf` inside the `test` folder with the following: 

```json
terraform {
  required_providers {
    greeting = {
      source = "hashicorp.com/edu/greeting"
      version = "1.0.0"
    }
  }
}

provider "greeting" {
  person = "John Doe" # This sets the default person for the provider
}

resource "greeting_resource" "example" {
  person = "Jim Halpert" # Input to the resource
}
```



### 3. Plan and Apply



```
terraform plan
terraform apply
```



### Challenge

Now that you've confirmed the plugin works, update the `main.tf` to include outputs that display: 

* Name 
* ID 



### 5. Clean Up



```
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