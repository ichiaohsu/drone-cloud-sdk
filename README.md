# drone-cloud-sdk
This plugin was orginally implemented for [Drone.io](http://docs.drone.io/installation) to clone git repo from Google Source Repository. But it simply could be the basis to any plugin aim to retrive the service from Google Cloud Platform. It basically uses **gcloud** command to authenticate the docker.

## Usage
Use this plugin docker in pipeline. Pass command you want to use like:
```bash
gcloud source repos clone [REPO_NAME] [DIRECTORY_YOU_WANT]
```

The settings of yml might be like:
```bash
pipeline
  get_config:
    image: [DOCKER_IMAGE_REPO]
    secrets: [google_credentials]
    commands:
    - gcloud source repos clone [REPO] [DIRECTORY]
```

