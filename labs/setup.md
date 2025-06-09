# Create a working directory in VS Code

### Step 1: Open the folder

1. Open Visual Studio Code
2. In Visual Studio Code, click **File** -> **Open Folder** and browse to your desired location. Create a new folder named `unlocking terraform` and enter it. Click **Open**.

# Configure GCP credentials

### **Step 1: Set up GCP Service Account**

1. In a browser, log into the [Google Cloud Console](https://console.cloud.google.com/) using the shared credentials in the spreadsheet below.
    * [Cloud credentials](https://docs.google.com/spreadsheets/d/1gjUd6gSoROwD7CUSqehBAtP4J1IUas5TPhu5zoGbVLY/edit?gid=2103659978#gid=2103659978)
2. Select your assigned project from the project dropdown at the top of the page
3. Search for "IAM & Admin" in the search bar.
4. Click **Service Accounts**
5. Click **CREATE SERVICE ACCOUNT**
6. Name it "terraform-admin-[YOUR-INITIALS]" (e.g., terraform-admin-jd for John Doe)
7. Click **CREATE AND CONTINUE**
8. Assign the "Editor" role
9. Click **DONE**
10. Click on the service account you just created
11. Click **KEYS** tab
12. Click **ADD KEY** -> **Create new key**
13. Choose **JSON** format
14. **IMPORTANT:** The key file will automatically download. Keep this file secure.

---

### **Step 2: Configure GCloud CLI**

1. In the Visual Studio Code Terminal run:

   ```
   gcloud auth activate-service-account --key-file=PATH_TO_YOUR_SERVICE_ACCOUNT_KEY.json
   ```

2. Set the default project:
   ```
   gcloud config set project PROJECT_ID
   ```
3. Set the default region:
   ```
   gcloud config set compute/region us-west1
   ```

---

### **Step 3: Test GCloud CLI**

1. Confirm that `gcloud` is properly authenticated:

   ```
   gcloud auth list
   ```

2. You should see your service account listed as active.

## Congratulations

Congratulations, you've successfully configured your machine.
