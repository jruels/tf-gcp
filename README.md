# Terraform Developer

This site includes the information for the Terraform Developer class.

## Windows VM credentials 

| Name                 | Username    | Virtual Machine Password | Virtual Machine Portal                  |
|----------------------|------------|--------------------------|-----------------------------------------|
| Angelina Amato       | WNS2022-01  | tekWNS2022!              | [https://my.ablazedesktop.com](https://my.ablazedesktop.com) |
| Arman Avetisyan      | WNS2022-02  | tekWNS2022!              |                                         |
| Daniel Dunn          | WNS2022-03  | tekWNS2022!              |                                         |
| Wogayehu Gebremariam | WNS2022-04  | tekWNS2022!              |                                         |
| Murugan Gnanasekaran | WNS2022-05  | tekWNS2022!              |                                         |
| Shaaheen Farahmand   | WNS2022-06  | tekWNS2022!              |                                         |
| Dave                 | WNS2022-07  | tekWNS2022!              |                                         |
| Kyle Margrave        | WNS2022-08  | tekWNS2022!              |                                         |
| Carilyn McCall       | WNS2022-09  | tekWNS2022!              |                                         |
| LaTonya McILwain     | WNS2022-10  | tekWNS2022!              |                                         |
| Andy Nadarewistsch   | WNS2022-11  | tekWNS2022!              |                                         |
| Mani Nagarajan       | WNS2022-12  | tekWNS2022!              |                                         |
| Reed O'Neal          | WNS2022-13  | tekWNS2022!              |                                         |
| Rosemary Otalor      | WNS2022-14  | tekWNS2022!              |                                         |
| Hardik Patel         | WNS2022-15  | tekWNS2022!              |                                         |
| Rob Sletten          | WNS2022-16  | tekWNS2022!              |                                         |
| Anupama Subramany    | WNS2022-17  | tekWNS2022!              |                                         |
| Son Thach           | WNS2022-18  | tekWNS2022!              |                                         |
| Arjun Thimmareddy   | WNS2022-19  | tekWNS2022!              |                                         |
| Lavanya Veluru      | WNS2022-20  | tekWNS2022!              |                                         |
| Instructor          | WNS2022-21  | tekWNS2022!              |                                         |

[Setup VM](labs/setup.md)

### Day 1 
[Courseware](https://github.com/jruels/tf-dev/raw/refs/heads/main/Go%20courseware%20and%20labs/Courseware%20Day%201.zip)

#### Day 1 labs 
To access the lab files, please open Visual Studio Code on the provided Windows VM and complete the following steps. 

1. Right click the PowerShell icon in the taskbar and select "Run as Administrator"
2. In the new PowerShell window install Git `choco install -y git`
3. After installing Git close PowerShell and Open Visual Studio Code
4. In the Visual Studio Code sidebar click the third icon down "Source Control"
5. Click "Clone repository" and enter the `https://github.com/jruels/tf-dev`
6. In the File Explorer window that pops up, create a new folder "repos" and select it. 
7. After the repository is done cloning, In Visual Studio Code click the top icon on the left sidebar "Explorer" and expand "Go Foundation" to access the Go lab files.

## Lab details 
[Cloud & Ansible machine info](https://docs.google.com/spreadsheets/d/1gTV6btPeIyyXylRkDn2_LNbWkf9BGU6wsi5eIb-ynLY/edit?gid=2103659978#gid=2103659978)


### Day 2
Lab 1: [Terraform - Create an instance](labs/tf-first-instance)    
Lab 2: [Variables and output](labs/tf-variables-and-output)   
Lab 3: [Multi resource deployment](labs/tf-more-variables)   
Lab 4: [Strings, bool, and numbers](labs/tf-even-more-variables)   
Lab 5: [Enable remote state](labs/tf-remote-state)   

### Day 3
Lab 6: [Import existing resources](labs/tf-import)   
Lab 7: [Provisioners](labs/tf-provisioner)   
Lab 8: [Use registry modules](labs/tf-module)   
Lab 9: [Write your own module](labs/tf-write-module)   
Lab 10: [Refactor monolithic codebase](labs/tf-refactor)   

### Day 4
[Ansible VM Info](ansible-vms.md)   
Lab 11: [Access Ansible VMs](labs/ssh-setup)   
Lab 12: [Install Ansible](labs/setup-ansible)   
Lab 13: [Ansible inventory](labs/inventory)   
Lab 14: [Ansible ad-hoc](labs/ad-hoc)   
Lab 15: [Ansible playbook fundamentals](labs/playbook-fun)   

### Day 5
Lab 16: [Ansible handling](labs/error-handling)   
Lab 17: [Write a custom Terraform provider](labs/write-custom-provider)   
