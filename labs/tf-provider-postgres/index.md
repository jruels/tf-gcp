## Creating a Terraform Provider for PostgreSQL

### Objective

In this lab, you will:

* Install PostgreSQL and pgAdmin using Chocolatey in the **Visual Studio Code terminal**.
* Configure PostgreSQL to accept local connections.
* Use **pgAdmin** to connect to the PostgreSQL server.
* Create a database named **bankingdb** and a **customer_accounts** table.
* Verify the database and table creation.

* Build a custom provider to connect to Postgres.
* Configure Terraform to use the local provider.
* Test your new provider.



------



### Step 1: Open Visual Studio Code as Administrator

1. Click the **Start Menu** and search for **Visual Studio Code**.
2. Right-click **Visual Studio Code** and select **Run as Administrator**.
3. Open the **terminal** by selecting **View > Terminal** or pressing Ctrl + ~.



**Note:** Running Visual Studio Code as Administrator is required for **Chocolatey** to install software.



------

### Step 2: Install PostgreSQL and pgAdmin



Run the following command in the **VS Code terminal** to install PostgreSQL and pgAdmin using **Chocolatey**:

```
choco install postgresql pgadmin4 -y
```

**Note:** The installation may take a few minutes.



------



### **Step 3: Start PostgreSQL Service**



1. In the **Visual Studio Code terminal**, run the following command to ensure PostgreSQL is running:

```
pg_ctl status
```

2. If the service is not running, start it with:

```
pg_ctl start
```

**Note:** PostgreSQL runs as a background service and listens on **port 5432** by default.



------



### **Step 4: Open pgAdmin and Connect to PostgreSQL**

1. **Launch pgAdmin** from the Start menu.
2. In the **pgAdmin dashboard**, click **“Add New Server”**.
3. In the **General tab**, set the following:
4. **Name:** PostgreSQL_Local
5. Switch to the **Connection tab** and configure:
   * **Host name/address:** localhost
   * **Port:** 5432
   * **Username:** postgres
   * **Password:** Post1260
   * Click **Save**, then double-click the server to establish the connection.



------



### **Step 5: Create the Banking Database**

1. In pgAdmin, **right-click** on **“Servers > PostgreSQL_Local”**.
2. Select **Create > Database**.
3. Set:
   * **Database Name:** bankingdb
   * **Owner:** postgres
   * Click **Save**.



------



### **Step 6: Create the Customer Accounts Table**

1. In pgAdmin, expand **Databases > bankingdb > Schemas > public**.

2. Right-click **Tables** and select **Query Tool**.

3. Enter the following SQL command to create the customer_accounts table:

```postgresql
CREATE TABLE customer_accounts (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    account_type VARCHAR(20) NOT NULL,
    balance DECIMAL(12,2) DEFAULT 0.00 NOT NULL
);
```

4. Click **Execute (F5)**.



------



### **Step 7: Verify the Table Creation**



Run the following SQL query in the **Query Tool** to confirm the table was created successfully:

```
SELECT * FROM customer_accounts;
```

If the table exists, you will see an empty result set with column headers.



------



### **Step 8: Verify Database Access with Credentials**

1. **Disconnect from PostgreSQL** in pgAdmin.
2. **Reconnect** using:
   * **Host:** localhost
   * **Username:** postgres
   * **Password:** Post1260
   * Navigate to **Databases > bankingdb** and expand **Tables**.
   * Ensure **customer_accounts** is listed.



------





### **Congratulations!**

You have successfully installed PostgreSQL, configured pgAdmin, created a database, and verified access to your table.



---



# Develop a custom provider to connect to PostgreSQL 





## Overview



You will create a custom Terraform provider in this lab that manages a PostgreSQL instance. You'll learn how Terraform providers work and implement basic CRUD (Create, Read, Update, Delete) operations using the Go programming language.

## Environment Setup



### Windows Setup



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

* Create a folder named `custom-tf-provider` and open it in Visual Studio Code



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
      "registry.terraform.io/example/banking" = "C:/Users/TEKstudent/go/bin"
  }
  direct {}
}
```



### 3. Initialize Go Module



```
go mod init terraform-provider-banking
```



This creates a `go.mod` file to manage dependencies.

## Provider Structure



Create the following directory structure:

* Inside VS Code, create `internal` and inside that create a sub-folder `provider`
* You should now have a folder path of `custom-tf-provider\internal\provider`



The provider will have this structure:

```
custom-tf-provider/
├── go.mod
├── main.go
└── internal/
    └── provider/
        ├── provider.go
        ├── database_client.go
        └── resource_customer_account.go
```



### 1. Create the Main Entry Point



Create `main.go` in the root directory with the following content:

```go
package main

import (
	"context"
	"log"

	bankingprovider "github.com/donis/terraform-provider-banking/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	// ✅ Start the Terraform provider with the correct configuration
	err := providerserver.Serve(context.Background(), bankingprovider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/example/banking",
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
	"database/sql"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// ✅ Rename DbClient to BankingDBClient
// type BankingDBClient struct {
// 	DB *sql.DB
// }

// ✅ Implement function to create a new database client
func NewBankingDBClient(host string, port int64, user, password, dbname string) (*BankingDBClient, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &BankingDBClient{DB: db}, nil
}

// ✅ Define bankingProvider struct
type bankingProvider struct {
	dbClient *BankingDBClient
}

// ✅ Implement Metadata function
func (p *bankingProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "banking"
}

// ✅ Implement Schema function
func (p *bankingProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"db_host":     schema.StringAttribute{Required: true},
			"db_port":     schema.Int64Attribute{Required: true},
			"db_user":     schema.StringAttribute{Required: true},
			"db_password": schema.StringAttribute{Required: true, Sensitive: true},
			"db_name":     schema.StringAttribute{Required: true},
		},
	}
}

// ✅ Implement Configure function
func (p *bankingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config struct {
		DBHost     string `tfsdk:"db_host"`
		DBPort     int64  `tfsdk:"db_port"`
		DBUser     string `tfsdk:"db_user"`
		DBPassword string `tfsdk:"db_password"`
		DBName     string `tfsdk:"db_name"`
	}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := NewBankingDBClient(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		resp.Diagnostics.AddError("Database Connection Error", fmt.Sprintf("Could not connect to database: %s", err.Error()))
		return
	}
	p.dbClient = client
	resp.ResourceData = client
}

// ✅ Implement DataSources function (required by Terraform Plugin Framework)
func (p *bankingProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil // No data sources implemented yet
}

// ✅ Register Resources (Terraform Resources like customer accounts)
func (p *bankingProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource { return &customerAccountResource{client: p.dbClient} },
	}
}

// ✅ Fix New() function to return the provider correctly
func New() provider.Provider {
	return &bankingProvider{}
}

```



This file handles:

- Provider configuration schema
- Database connection info
- Data structure
- Error handling and diagnostics

### 4. Build and Install



Windows:

```
go build -o terraform-provider-banking.exe
move terraform-provider-custom-s3.exe %USERPROFILE%\go\bin\
```



## Testing the Provider



### 1. Create Test Directory



Inside the `custom-tf-provider` folder, create a `test` folder.



### 2. Create Test Configuration



Create `main.tf` inside the `test` folder with the following: 

```json
terraform {
  required_providers {
    banking = {
      source = "registry.terraform.io/example/banking"
    }
  }
}

provider "banking" {
  db_host     = "localhost"
  db_port     = 5432
  db_user     = "postgres"
  db_password = "Post1260"
  db_name     = "bankingdb"
}

resource "banking_customer_account" "customer1" {
  first_name   = "Alice"
  last_name    = "Doe"
  email        = "alice@example.com"
  account_type = "savings"
  balance      = 1500.75
}
```



### 3. Plan and Apply



```
terraform plan
terraform apply
```



### 4. Verify

You can check pgAdmin and confirm the table was populated to verify it worked. 



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