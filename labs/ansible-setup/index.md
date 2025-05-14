# Lab Setup 
Open a Terminal session in Visual Studio Code and `cd` to the extracted lab directory. Inside the directory, you will see a `keys` directory. Enter it using `cd` and run the following commands.   
* [Ansible VM information](https://docs.google.com/spreadsheets/d/1gTV6btPeIyyXylRkDn2_LNbWkf9BGU6wsi5eIb-ynLY/edit?gid=1973346361#gid=1973346361)

```
ssh -i lab.pem ansible@<Tower VM IP from the spreadsheet> 
```
After logging into the Tower server, test the connection to Managed Node1 and Managed Node2

## Set up a remote SSH session in Visual Studio Code.   
### Create the SSH configuration file.

On the left sidebar, click the icon that looks like a computer with a connection icon.

In the Remote Explorer, hover your mouse cursor over **SSH**, click on the gear icon (⚙️) in the top right corner, and select the top option: `C:\Users\tekstudent\.ssh\config` This will open the SSH configuration file in a new editor tab.
### Add the SSH configuration for the lab servers.
Add the following lines to the SSH configuration file, replacing `<IP of Tower server from the spreadsheet>` with the actual IP addresses of your Tower and `<Path to the cloned lab directory/keys/lab.pem>` with the correct path to the `lab.pem` file in your lab directory.

**PROTIP**: Right-click the `lab.pem` in the Visual Studio Code Explorer and click `Copy Path`. Paste it below as the value for `IdentifyFile`

```plaintext
Host tower
  HostName <IP of Tower server from the spreadsheet>
  IdentityFile <Path to the cloned lab directory/keys/lab.pem>
  User ansible
```

### Save the SSH configuration file.
Save the changes to the SSH configuration file and close it.
### Connect to the lab servers.
1. In the Remote Explorer, you should now see the entry for the Tower server under "SSH Targets."
2. Click on the entry to connect to the Tower server.
3. Visual Studio Code will open a new window connected to the Tower server.
4. You can now open a terminal in this new window and run commands on the Tower server.

### Create a working directory
In Visual Studio Code, you can create a new folder or file as if it was on your local machine.
Click **Open Folder** and select `/home/ansible`.
In future labs, you will create a directory for each lab.

## Congratulations!
You have successfully set up your lab environment and are ready to start working on the labs.
