# CICD

The CICD is kicked off on any push to the main branch or pull request.






github action public library:

https://github.com/marketplace?type=actions




# Real World Scnearios

- The devops team usually has a starter project in a git repo where teams that are starting new projects or microservices can just clone and it has most of the devops files setup and done. The devs just have to rename the project and git repo paths and some other simple configurations to get the code deployed to a lower environment.

- CICD are usually kicked of on a PR merge or release tag (i.e. v1.2.0). Not really a push to a dev branch - the dev branch serves as a place to save the work before a PR is created and merged.

- Deployments are usually done to multiple environments such as Dev, QA, Stage, Prod.

- During the CICD steps, there's usually security scans on Docker images that the app uses, i.e. if python, the python image would get scanned. Scans for code smell, scans for known vulnerbilities and exposures.

- If there were any steps that needed secrets, a github action can be used to fetch the secrets, whether from vault or AWS KMS.



