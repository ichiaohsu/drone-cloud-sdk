FROM google/cloud-sdk:181.0.0-alpine
ADD drone-gsr /bin/drone-gsr

ENTRYPOINT ["/bin/drone-gsr"]