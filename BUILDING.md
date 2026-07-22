# Building the HPE COSI Driver

- Clone the repo: git clone <https://github.com/hpe-storage/cosi-driver>
- cd to cosi-driver
- Turn on go modules support `export GO111MODULES=on`
- Set REGISTRY_NAME env to point to your image registry, if other than docker.io (default)
- Run `make all`. It will take HTTP_PROXY & HTTPS_PROXY env details, if set, and build the image.
- If different proxy configuration needs to be configured then use below command
   ```console
   make all HTTP_PROXY=http://custom-proxy:8080 HTTPS_PROXY=http://custom-proxy:8080
   ```
   > Note: Ensure Docker is installed and running before executing the make all command.
         To build multiarch docker images, we need a docker builder to be running.
         Refer to create docker builder: https://docs.docker.com/reference/cli/docker/buildx/create/
