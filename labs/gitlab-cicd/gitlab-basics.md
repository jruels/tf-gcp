# GitLab Basics Lab

This lab will walk you through the fundamental features of GitLab and help you create your first CI/CD pipeline.

## Part 1: Getting Started with GitLab

### Set Up SSH Key in GitLab
1. Generate an SSH key (if you don't have one):
   ```bash
   ssh-keygen -t ed25519 -C "your.email@example.com"
   ```

2. Copy your public key to clipboard:
   ```bash
   # For macOS
   tr -d '\n' < ~/.ssh/id_ed25519.pub | pbcopy
   
   # For Linux (requires xclip)
   xclip -sel clip < ~/.ssh/id_ed25519.pub
   
   # For Windows Git Bash
   cat ~/.ssh/id_ed25519.pub | clip
   ```

3. Add the key to GitLab:
   - Sign in to GitLab
   - Click your avatar in the top right
   - Select "Edit profile"
   - On the left sidebar, select "SSH Keys"
   - Click "Add new key"
   - Paste your public key in the "Key" box
   - Add a title (e.g., "Work Laptop")
   - Click "Add key"

### Create a New Project
1. Log into GitLab
2. Click the "+" button in the top navigation bar
3. Select "New project/repository"
4. Choose "Create blank project"
5. Fill in the following:
   - Project name: `my-first-pipeline`
   - Project URL: Make sure your username is selected (NOT any group)
   - Visibility Level: Public
   - Initialize repository with a README: Yes
6. Click "Create project"

The project URL should look like: `https://gitlab.com/YOUR-USERNAME/my-first-pipeline`

## Part 2: Basic Repository Management

### Clone Your Repository
1. Configure git (if you haven't already):
   ```bash
   git config --global user.name "Your Name"
   git config --global user.email "your.email@example.com"
   ```

2. Set up SSH key (if you haven't already):
   ```bash
   # Generate SSH key
   ssh-keygen -t ed25519 -C "your.email@example.com"
   
   # Start the SSH agent
   eval "$(ssh-agent -s)"
   
   # Add your SSH key to the agent
   ssh-add ~/.ssh/id_ed25519
   
   # Copy your public key
   cat ~/.ssh/id_ed25519.pub
   ```
   Then add the copied key to GitLab: Settings → SSH Keys

3. Click the "Clone" button
4. Copy the SSH URL (not HTTPS)
5. Open your terminal and run:
   ```bash
   git clone <your-repo-url>
   cd my-first-pipeline
   ```

### Create a Simple Application
1. Create a new file called `index.html`:
   Create a new file called `index.html` and add the following content using your preferred code editor:

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
1. 
   Create a new file called `.gitlab-ci.yml` and add the following content using your preferred code editor:

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
```bash
git add index.html .gitlab-ci.yml
git commit -m "Add website and pipeline configuration"

# Debug git push issues
git remote -v  # Check remote URL
git config --get remote.origin.url  # Verify remote URL
ssh -T git@gitlab.com  # Test SSH connection to GitLab

# If everything looks good, push
git push -u origin main
```

## Part 3: Understanding CI/CD Features

### Pipeline Visualization
1. In Gitlab, go to Build → Pipelines
2. You should see your pipeline running
3. Click on the pipeline to see the stages:
   - Test stage: Verifies the HTML file exists
   - Build stage: Creates a public directory
   - Deploy stage: Deploys to GitLab Pages

### Job Logs
1. Click on any job in the pipeline
2. Observe the real-time log output
3. Notice the job status indicators:
   - Blue: Running
   - Green: Passed
   - Red: Failed

### Pipeline Settings
1. Go to Settings → CI/CD
2. Explore key settings:
   - Runners
   - Variables
   - Artifacts
   - Secure files

## Part 4: Working with Branches

### Create a Feature Branch
1. Create a new branch:
   ```bash
   git checkout -b feature/add-styling
   ```

2. Modify `index.html`:
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
3. Commit and push:
   ```bash
   git add index.html
   git commit -m "Add CSS styling"
   git push origin feature/add-styling
   ```
### Create a Merge Request
1. Go to Merge Requests → New merge request
2. Select:
   - Source branch: `feature/add-styling`
   - Target branch: `main`
3. Click "Compare branches and continue"
4. Fill in:
   - Title: "Add CSS styling to website"
   - Description: "Added basic CSS to improve the page appearance"
5. Click "Create merge request"

### Review Pipeline and Merge
1. Observe that the pipeline automatically runs for your merge request
2. Wait for all jobs to pass
3. Click "Merge" when ready

## Part 5: GitLab Pages

### View Your Deployed Site
1. Go to Deploy → Pages
2. You should see your site's URL (usually `https://<username>.gitlab.io/<project-name>`)
3. Click the URL to view your deployed website

## Part 6: Additional Features

### Issues
1. Go to Issues → New issue
2. Create an issue:
   - Title: "Add footer to website"
   - Description: "We should add a footer with contact information"
   - Labels: Select or create appropriate labels
3. Click "Create issue"

### Project Wiki
1. Go to Wiki
2. Click "Create your first page"
3. Add some documentation about your project
4. Use markdown formatting to structure your content

### Analytics
1. Explore Analyze → Repository Analytics
2. View commit history and contribution graphs
3. Check CI/CD analytics for pipeline performance

## Conclusion

Congratulations! You've learned the basics of:
- Creating a GitLab project
- Setting up a basic CI/CD pipeline
- Working with branches and merge requests
- Deploying to GitLab Pages
- Using issues and wiki features

Next Steps:
- Add more complex pipeline stages
- Explore environment deployments
- Set up code quality checks
- Configure automated testing
- Implement security scanning 


