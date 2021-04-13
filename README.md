# Introduction
This Password Manager is built on the top of Fortanix SDKMS. It uses the Fortanix SDKMS as backend storage and calls the appropriate APIs to perform different operations like create a secret, fetch a secret, log into the SDKMS.
It provides a set of user friendly CLIs to interact with the Password Manager

# Build docker image (from src)(optional step)
```
# git clone git@github.com:innovolt/password-manager.git
# bash password-manager/tools/build.sh
```

# Run the Password Manager
```
# docker run -it --network host navlok/innovolt-pm:latest
root@docker-desktop:/go#

Note: Replace navlok/innovolt-pm:latest with the local docker image built from source if it is required
```

# Log into Password Manager using Fortanix SDKMS user credentials
```
root@docker-desktop:/go# innovolt-pm login user --username <username> --password <password>
```

## Example
```
root@docker-desktop:/go# innovolt-pm login user --username abc@xyz.com --password mypwd
Logged in successfully
```


# Create a password for a domain
```
root@docker-desktop:/go# innovolt-pm secret create <secretName>
```

## Example
```
root@docker-desktop:/go# innovolt-pm secret create mysecret
Domain [https://amazon.in]: www.facebook.com
Username: demouser
Password: demopassword
+---------------+--------------------------------------+
| ACCOUNT NAME  |              ACCOUNT ID              |
+---------------+--------------------------------------+
| hello-curl    | 27783b15-d813-40dd-b397-c8ef45f3c7ac |
+---------------+--------------------------------------+
| AirB          | 6d81667d-f253-4ba4-96c1-61e978538b9b |
+---------------+--------------------------------------+
| Demo-App-Cert | 959f6e3b-786b-4e47-aef4-b65ef0087032 |
+---------------+--------------------------------------+
| Demo1         | dc4762dc-0cff-456d-8f29-4aadde0f1635 |
+---------------+--------------------------------------+
Select an Account [ID]: 27783b15-d813-40dd-b397-c8ef45f3c7ac
+------------+--------------------------------------+
| GROUP NAME |               GROUP ID               |
+------------+--------------------------------------+
| Demo-Group | 0cc9786e-5e05-458d-a81b-4df645442097 |
+------------+--------------------------------------+
Select a Group [ID]: 0cc9786e-5e05-458d-a81b-4df645442097
Secret is created successfully.
```

# Fetch a secret
```
root@docker-desktop:/go# innovolt-pm secret get <secretName>
```

## Example
```
root@docker-desktop:/go# innovolt-pm secret get mysecret
+---------------+--------------------------------------+
| ACCOUNT NAME  |              ACCOUNT ID              |
+---------------+--------------------------------------+
| hello-curl    | 27783b15-d813-40dd-b397-c8ef45f3c7ac |
+---------------+--------------------------------------+
| AirB          | 6d81667d-f253-4ba4-96c1-61e978538b9b |
+---------------+--------------------------------------+
| Demo-App-Cert | 959f6e3b-786b-4e47-aef4-b65ef0087032 |
+---------------+--------------------------------------+
| Demo1         | dc4762dc-0cff-456d-8f29-4aadde0f1635 |
+---------------+--------------------------------------+
Select an Account [ID]: 27783b15-d813-40dd-b397-c8ef45f3c7ac
+------------+--------------------------------------+
| GROUP NAME |               GROUP ID               |
+------------+--------------------------------------+
| Demo-Group | 0cc9786e-5e05-458d-a81b-4df645442097 |
+------------+--------------------------------------+
Select a Group [ID]: 0cc9786e-5e05-458d-a81b-4df645442097
+----------+------------------+----------+--------------+
|   NAME   |      DOMAIN      | USERNAME |   PASSWORD   |
+----------+------------------+----------+--------------+
| mysecret | www.facebook.com | demouser | demopassword |
+----------+------------------+----------+--------------+
```
