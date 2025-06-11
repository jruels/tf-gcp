# GitLab Basics Lab 

This lab uses Visual Studio Code to walk you through fundamental GitLab features and sets up your first CICD pipeline.

## Part 1: Getting Started with GitLab

### Set Up SSH Key in GitLab

1. **Open Visual Studio Code.**

2. **If you're using Windows, open a Git Bash terminal:**

   - Click the **+▼** button in the terminal panel
   - Select **Git Bash** from the list

3. **In the VS Code Explorer sidebar**, right-click the `unlocking terraform` folder and select **New Folder**. Name it `gl-ssh-key`.

4. **Generate the SSH key and save it as** `**gitlab_key**` **inside the** `**gl-ssh-key**` **folder:**

   - In the Git Bash terminal, run:

     ```bash
     ssh-keygen -t ed25519 -C "your.email@example.com" -f "<path-to-unlocking-terraform>/gl-ssh-key/gitlab_key"
     ```

     Replace `<path-to-unlocking-terraform>` with the path to your `unlocking terraform` folder

5. **Create and configure the** `SSH config` **file:**

   - In VS Code, go to **File → Open File...**

   - Navigate to your `.ssh` folder:

     - On Windows, type `%USERPROFILE%\.ssh` into the file dialog and press Enter
     - On macOS/Linux, use `~/.ssh`

   - If the `config` file exists, select it. If not, create a new file named `config`.

   - Add the following content:

   - **NOTE**: Update the `IdentifyFile` field with the correct path to your `gitlab_key`.

     ```
     Host gitlab.com
       HostName gitlab.com
       User git
       IdentityFile ~/unlocking\ terraform/gl-ssh-key/gitlab_key
       IdentitiesOnly yes
     ```

6. **Copy your public key:**

   - In the VS Code Explorer, open `gitlab_key.pub` from the `gl-ssh-key` folder
   - Highlight the contents and copy them (`Ctrl+C` / `Cmd+C`)

7. **Create a GitLab account:**

   - Go to [https://gitlab.com](https://gitlab.com/) and click **Register**
   - Use your email and create a new GitLab account
   - During sign-up, select the **trial** option (free trial is acceptable)
   - GitLab may prompt you to verify your identity — choose **phone number** verification to ensure CI/CD pipelines can run
   - Once signed up and verified, log in

8. **Add the SSH key to GitLab:**

   - In GitLab, go to your avatar > **Edit profile** > **SSH Keys**
   - Paste the copied key into the **Key** field
   - Add a title such as `Work Laptop`
   - Click **Add key**

9. **Create a new project in GitLab:**

   - At the top of the page, click "**+**" > **New project/repository**
   - Select **Create blank project**
   - Fill in:
     - Project name: `my-first-pipeline`
     - Project URL: Make sure your username is selected (**NOT any group**)
     - Visibility Level: Public
     - Initialize repository with a README: Yes
   - Click **Create project**

10. **Verify pipelines are working:**

    - After your project is created, go to **Build → Pipelines** from the left-hand menu
    - Confirm that GitLab does **not prompt you again for identity verification**
    - If it does, provide your phone number and verify with the code received.
    - This ensures pipelines will run properly before you clone or push any code

## Part 2: Basic Repository Management

### Clone Your Repository

1. In GitLab:
   - Click the **Clone** button
   - Copy the **SSH URL**
2. In VS Code:
   - Open **Command Palette** (`Ctrl+Shift+P` or `Cmd+Shift+P`)
   - Type `Git: Clone` and select it
   - Paste the SSH URL
   - When prompted, select the parent folder where you want to save the project
   - VS Code will ask: "Would you like to open the cloned repository?" Click **Open**

### Create a Simple Application

1. In the new VS Code windows, select Explorer (left panel):

   - Right-click on the project root > **New File** > Name it `index.html`
   - Paste the following content:

   ```html
   <!DOCTYPE html>
   <html>
   <head>
       <title>My First Pipeline</title>
   </head>
   <body>
       <h1>Hello, GitLab!</h1>
       <p>This page was deployed using GitLab CI/CD.</p>
   </body>
   </html>
   ```

### Create Your First Pipeline

1. Right-click on the project root > **New File** > Name it `.gitlab-ci.yml`

2. Paste the following content:

   ```yaml
   stages:
     - test
     - build
     - deploy
   
   test-job:
     stage: test
     image: alpine
     script:
       - echo "Running tests..."
       - test -f index.html
       - grep "GitLab CI/CD" index.html
   
   build-job:
     stage: build
     image: alpine
     script:
       - echo "Building..."
       - mkdir public
       - cp index.html public/
     artifacts:
       paths:
         - public
   
   pages:
     stage: deploy
     script:
       - echo "Deploying to GitLab Pages..."
     artifacts:
       paths:
         - public
     only:
       - main
   ```

### Commit and Push Changes

1. In the Source Control view (left sidebar, or `Ctrl+Shift+G`):
   - You’ll see your two new files listed
   - Click the **+** (plus sign) next to each file to stage it
   - Enter a commit message: `Add website and pipeline configuration`
   - Click **Commit** to commit the changes. 
   - Click **Sync** and select yes, when prompted to confirm.

## Part 3: Understanding CI/CD Features

### View the Pipeline

1. In GitLab:
   - Go to **Build → Pipelines**
   - Click your running pipeline to view stages

### View Job Logs

1. Click any job to view logs in real-time
2. Notice:
   - Blue: Running
   - Green: Passed
   - Red: Failed

### Review Pipeline Settings

1. In GitLab:
   - Go to **Settings → CI/CD**
   - Browse through Runners, Variables, Artifacts, Secure Files

## Part 4: Working with Branches

### Create a Feature Branch

1. In the bottom-left corner of VS Code, click the branch name (e.g., `main`) in the status bar to open the branch menu

2. Select **Create new branch** and name it: `feature/add-styling`

3. Open `index.html` and update it to include the **<style>** block:

   ```html
   <!DOCTYPE html>
   <html>
   <head>
       <title>My First Pipeline</title>
       <style>
           body {
               font-family: Arial, sans-serif;
               margin: 40px;
               line-height: 1.6;
               color: #333;
           }
           h1 {
               color: #2084E2;
           }
       </style>
   </head>
   <body>
       <h1>Hello, GitLab!</h1>
       <p>This page was deployed using GitLab CI/CD.</p>
   </body>
   </html>
   ```

4. Save the file.

5. Commit and push:

   - Go to Source Control
   - Add commit message: `Add CSS styling`
   - Click ✓ to commit
   - In the pop-up window, select **Yes** to stage all changed files.
   - Click **Publish Branch**

### Create a Merge Request

1. In GitLab:
   - Go to **Merge Requests → New merge request**
   - Select:
     - Source: `feature/add-styling`
     - Target: `main`
   - Click **Compare branches and continue**
   - Fill in title/description
   - Click **Create merge request**
2. Wait for pipeline to pass and click **Merge**

## Part 5: GitLab Pages

### View Your Deployed Site

1. In GitLab:
   - Navigate to **Deploy → Pages**
   - Click the site URL to view it live

## Part 6: Additional Features

### Create an Issue

1. In GitLab:
   - Go to **Issues → New issue**
   - Add:
     - Title: `Add footer to website`
     - Description: `We should add a footer with contact information.`
   - Click **Create issue**

### Create Wiki Documentation

1. In GitLab:
   - Go to **Wiki**
   - Click **Create your first page**
   - Write content in Markdown
   - Save page

### Explore Analytics

1. In GitLab:
   - Go to **Analyze → Repository Analytics**
   - View commit history, pipeline duration, and contributions

## Conclusion

You’ve successfully:

- Created a GitLab project
- Used VS Code to manage your repo and files
- Set up a pipeline and deployed to GitLab Pages
- Explored collaboration and CI/CD features

**Next Steps:**

- Expand your `.gitlab-ci.yml`
- Use environments
- Add linting and security scans
- Experiment with GitLab Auto DevOps