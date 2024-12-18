# Bash script to create a Kubernetes secret for the admin user of the HPE COSI driver

This script is responsible for creating the Kubernetes secret of the HPE COSI driver's admin user. It creates a namespace for the secret, checks if any older secret with the same name exists  and deletes it if found and creates a new secret with data provided as input arguments to the script. The secret's name and namespace will be displayed in the terminal upon successful execution of the script.

## Prerequisites

- Linux machine with bash shell
- Access to the Kubernetes cluster
- `kubectl` installed and configured

## How to use it

1. Clone the GitHub repository:
    ```console
    git clone https://github.com/hpe-storage/cosi-driver.git
    ```

2. Navigate to the directory containing the script:
    ```console
    cd cosi-driver/scripts/cosi_secret/
    ```

3. Run the script with the required options. For example, to create a Kubernetes secret for the admin user of the HPE COSI driver:
    ```console
    bash hpe_create_cosi_secret.sh --s3_credentials "testuser,testuser" --s3_endpoint "http://s3-cluster.instance.com" --glcp_credentials "glcp_user_client_id,secret_key" --dscc_zone "scdev03-xxx.xx.xxx.hpe.com" --cluster_serial_number "XXXXXXXXXX"
    ```

## Options

Option Name<br>(-short_form,--long_form argument)                                   | Description           | Required
:---------------------------------------------------------------------------------- |:--------------------- | :------
 `-h,--help`                             | Displays the help message.| Yes
 `-s,--s3_credentials S3_ACCESS_KEY,S3_SECRET_KEY` | The S3 admin user's access key and secret key, separated by commas.| Yes
 `-e,--s3_endpoint ENDPOINT`                      | The S3 endpoint of the HPE Alletra Storage MP X10000 object storage array which is derived from its S3 DNS Subdomain name.| Yes
 `-g,--glcp_credentials GLCP_USER_CLIENTID,GLCP_USER_SECRET_KEY` | The GLCP user's client ID and secret key, separated by commas.| Yes
 `-d,--dscc_zone DSCC_ZONE`              | The DSCC zone.| Yes
 `-c,--cluster_serial_number CLUSTER_SERIAL_NUMBER` | The serial number of the cluster.| Yes
 `--cosi_secret_name COSI_SECRET_NAME`    | The desired name of the secret. Default: `cosi-user-secret-$CLUSTER_SERIAL_NUMBER` | No
 `--cosi_secret_namespace COSI_SECRET_NAMESPACE` | The desired namespace of the secret. Default: `default` | No


## Output

The script will display the progress of the script on the terminal as it tries to create the namespace for the secret and the secret using `kubectl` commands. Any errors encountered during the execution will be displayed in the terminal.

## Troubleshooting

If you encounter any issues while running the script, please check the following:

- Ensure you have the necessary permissions to access the Kubernetes cluster.
- Check if `kubectl` is installed and properly configured.
- Make sure the script is executable. If not, run `chmod +x hpe-create_cosi_secret.sh` to make it executable.
