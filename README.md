# Docker Compose CD Workflow
This project intends to demonstrate a Continuous Deployment (CD) workflow for an application deployed using docker compose to a machine with SSH access. The workflow is implemented using GitHub Actions and is located in [`.github/workflows/release.yaml`](https://github.com/Adibov/docker-compose-cd-workflow/blob/master/.github/workflows/release.yaml).

## Features
- **Automated Deployment**: The workflow is triggered on every push to the `master` branch.
- **Docker Compose**: The application is deployed using docker compose.
- **Near-Zero Downtime**: The application is updated with near-zero downtime.
- **Rollback With One Click**: The workflow can be easily rolled back to the previous version with one click (just run the desired deployment job).


## Licence
This project is licensed under the MIT License - see the [LICENSE](https://github.com/Adibov/docker-compose-cd-workflow/blob/master/LICENSE).