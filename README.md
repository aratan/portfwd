
# GitDock 0.0.6
Automate your Docker deployments with GitDock! This command-line interface (CLI) tool monitors your GitHub repositories for new commits and automatically builds and pushes Docker images to Docker Hub. With GitDock, you can streamline the deployment process and easily manage your containerized applications.


<img src="https://miro.medium.com/v2/resize:fit:640/format:webp/1*YpIbJwWxQgNl2BEozZ4FZQ.png" alt="gitdock for cicd" width="500" height="500">

# Getting Started
## To get started with GitDock, you'll need to have a few things set up:

A GitHub account and repository with a Dockerfile and Docker Compose file.
A Docker Hub account and a repository with the same name as your GitHub repository.
A personal access token from GitHub with repo and read:user scopes.

## Your Docker Hub login credentials.
Once you have everything set up, you can install GitDock and start deploying your containers! Here's how to use GitDock:


> gitdock  <user> <repo> <branch> <email-alert>

<user>: Your Docker Hub username.
<repo>: The name of your GitHub repository.
<branch>: The name of the branch you want to monitor for changes (usually main or master).
<e-mail>: The e-mail Alert.

## FILE: .env


### Github and E-Mail ####

TOKEN_GITHUB=

USER_EMAIL=

USER_PASS=

### Discord ####

DISCORD_BOT_TOKEN=

ID_SERVER=

ID_CHANNEL=
  
## For example:

conda activate mlops  "pytest"

To not run gitdock with root permissions, you can do the following:

sudo groupadd docker

sudo usermod -aG docker $USER

docker login

![image](https://user-images.githubusercontent.com/4398830/233807866-225dd44c-cc16-4140-b4cb-d590ae79ac6b.png)
![image](https://user-images.githubusercontent.com/4398830/233482186-cbb173f7-635b-459a-a2d7-396cf46c5500.png)
![image](https://user-images.githubusercontent.com/4398830/233482379-d484e385-bc7d-455a-b8c1-ac506f0a1990.png)



## Features
GitDock offers the following features:

Automated Docker deployments based on GitHub commits.

Support for Dockerfiles and Docker Compose files.

Streamlined deployment process.

Easy management of containerized applications.

## Conclusion
If you're looking for a way to automate your Docker deployments, GitDock is the tool for you! With its easy setup process and streamlined 

deployment process, you can focus on developing your applications and let GitDock handle the rest. Try it out today!


