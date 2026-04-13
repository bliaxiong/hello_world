# CICD

Prerequisite: 
- Need a Kubernetes envrionment.
- Need to install ArgoCD.

I added a ci and release yaml file with Github Actions to run linting, unit tests, docker builds and push image to docker repository.

Then the release yaml creates the release. In my experience, this is usually done with human intervention to create a tag for release with manual blue/green deployment.

The CICD is kicked off on any push to the main branch or pull request.



### Notes:

- My initial release.yaml file had the immediate sync built-in and tried to make a POST call but failed due to the nature of the setup. I had to remove that part, update ArgoCD to poll every 20 seconds or so for faster relase instead of the default 3 min wait.

- Removed the ArgoCD callback in the Github Actions release.yaml file since the cluster is local to my WSL2. Using poller since its easier and in this scenario.

- There was a permissions issue with the Github repo CI pipeline and creating a Docker image and pushing it. I enabled write permissions in the settings.




# Real World Scnearios

- The devops team usually has a starter project in a git repo where teams that are starting new projects or microservices can just clone and it has most of the devops files setup and done. The devs just have to rename the project and git repo paths and some other simple configurations to get the code deployed to a lower environment.

- CICD are usually kicked of on a PR merge or release tag (i.e. v1.2.0). Not really a push to a dev branch - the dev branch serves as a place to save the work before a PR is created and merged.

- Deployments are usually done to multiple environments such as Dev, QA, Stage, Prod.

- During the CICD steps, there's usually security scans on Docker images that the app uses, i.e. if python, the python image would get scanned. Scans for code smell, scans for known vulnerbilities and exposures.

- If there were any steps that needed secrets, a github action can be used to fetch the secrets, whether from vault or AWS KMS.

- If using polling to a private repo, you'd need to configure it to use tokens, secrets to access the private repo.

- Kubernetes Mesh

- Istio

- Spinnaker (we used at Target)



