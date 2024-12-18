# HPE COSI driver log collector script

This script is responsible for collecting logs from the HPE COSI driver. It gathers relevant log files and stores them in a specified location. The collected logs can be used for troubleshooting and debugging purposes.

## Prerequisites

- Linux machine with bash shell
- Access to the Kubernetes cluster where HPE COSI driver is deployed
- `kubectl` installed and configured

## How to use it

1. Clone the GitHub repository:
    ```console
    git clone https://github.com/hpe-storage/cosi-driver.git
    ```

2. Navigate to the directory containing the script:
    ```console
    cd cosi-driver/scripts/log-collection/
    ```

3. Run the script with the appropriate options. For example, to collect logs from the HPE COSI deployment in the 'default' namespace:
    ```console
    ./hpe-logcollector.sh -n default
    ```

## Options

- `-h, --help`     Display this help message and exit
- `-n, --namespace NAMESPACE` Collect logs from HPE COSI deployment in namespace

## Output

The script will collect the logs and store them in a .tar.gz file under the /var/log/ directory. The name of the file will be in the format `hpe-cosi-provisioner-logs-<hostname>-<timestamp>.tar.gz`.

## Troubleshooting

If you encounter any issues while running the script, please check the following:

- Ensure you have the necessary permissions to access the Kubernetes cluster and the namespace where the HPE COSI driver is deployed.
- Check if `kubectl` is installed and properly configured.
- Make sure the script is executable. If not, run `chmod +x hpe-logcollector.sh` to make it executable.

## Dependencies

While fetching provisioner/controller logs using kubectl command, we have used selector label "container-object-storage-interface". If this label is not attached to the object or changed then log collection script will not be able to fetch logs.
