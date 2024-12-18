# © Copyright 2024 Hewlett Packard Enterprise Development LP
#!/bin/bash
# This script is responsible for creating the Kubernetes secret of the HPE COSI driver's admin user.
############################################################
# Help                                                     #
############################################################
help()
{
    # Display Help
    echo -e "\nCreate HPE COSI driver secret"
    echo -e "\nSyntax:"
    echo -e "     hpe_create_cosi_secret.sh [-h|--help] [-s|--s3_credentials S3_ACCESS_KEY,S3_SECRET_KEY] [-e|--s3_endpoint ENDPOINT] [-g|--glcp_credentials GLCP_USER_CLIENTID,GLCP_USER_SECRET_KEY] [-d|--dscc_zone DSCC_ZONE] [-c|--cluster_serial_number CLUSTER_SERIAL_NUMBER] [--cosi_secret_name COSI_SECRET_NAME] [--cosi_secret_namespace COSI_SECRET_NAMESPACE]\n"
    echo -e "Options:"
    echo -e "-h|--help                                                          Print this Help."
    echo -e "-s|--s3_credentials S3_ACCESS_KEY,S3_SECRET_KEY                    Specify the S3 admin user's access key and secret key separated by commas. (REQUIRED) E.g. -s \"testuser,testuser\""
    echo -e "-e|--s3_endpoint ENDPOINT                                          Specify the S3 endpoint of the HPE Alletra Storage MP X10000 object storage array, which is derived from its S3 DNS Subdomain name. (REQUIRED) E.g. -e \"http://s3-cluster.storage.com\""
    echo -e "-g|--glcp_credentials GLCP_USER_CLIENTID,GLCP_USER_SECRET_KEY      Specify the GLCP user's client ID and secret key separated using commas. (REQUIRED) E.g. -g \"ac40eac7-9a06-4ccb-8575-25a064af3a29,81656ff23d1f11efa0f6cabe03166b4d\""
    echo -e "-d|--dscc_zone DSCC_ZONE                                           Specify the DSCC zone. (REQUIRED) E.g. -d \"scdev03-app.qa.cds.hpe.com\""
    echo -e "-c|--cluster_serial_number CLUSTER_SERIAL_NUMBER                   Specify the serial number of the cluster. (REQUIRED) E.g. -c \"HFDT1Y87BC\""
    echo -e "--cosi_secret_name COSI_SECRET_NAME                                Specify the desired name of the secret. Default: 'cosi-user-secret-\$CLUSTER_SERIAL_NUMBER'. (OPTIONAL) E.g. --cosi_secret_name \"cosi-user-secret-hfdt1y87bc\""
    echo -e "--cosi_secret_namespace COSI_SECRET_NAMESPACE                      Specify the desired namespace of the secret. Default: 'default'. (OPTIONAL) E.g. --cosi_secret_namespace \"cosi-user\""
    exit 0
}

############################################################
# Create Kubernetes secret                                 #
############################################################
create_secret()
{
    # Check if namespace already exists. Doesn't print output and err to terminal.
    if ! kubectl get namespace/$cosi_secret_namespace > /dev/null 2>&1
    then
        echo "Namespace: '${cosi_secret_namespace}' does not exist. Creating namespace..."
        # Create the namespace.
        kubectl create namespace $cosi_secret_namespace
    fi
    # Check if secret already exists. Doesn't print output and err to terminal.
    if kubectl get secret/$cosi_secret_name -n $cosi_secret_namespace > /dev/null 2>&1
    then
        echo -e "Secret: '${cosi_secret_name}' already exists in namespace: '${cosi_secret_namespace}'. Please provide a different secret name in the arguments or delete the existing secret before running this script.\nWARNING: Deleting an existing secret may cause failures if COSI APIs are currently using it."
        exit 1
    fi
    # Create new secret. If kubectl returns error, print to terminal.
    echo "Creating new secret: '${cosi_secret_name}' in namespace: '${cosi_secret_namespace}'..."
    if kubectl create secret generic $cosi_secret_name -n $cosi_secret_namespace \
        --from-literal=accessKey=${s3_creds[0]} \
        --from-literal=secretKey=${s3_creds[1]} \
        --from-literal=endpoint=${s3_endpoint} \
        --from-literal=glcpUserClientId=${glcp_creds[0]} \
        --from-literal=glcpUserSecretKey=${glcp_creds[1]} \
        --from-literal=dsccZone=$dscc_zone \
        --from-literal=clusterSerialNumber=$cluster_serial_number > /dev/null
    then
        # If no error is found, print the following success message.
        echo -e "Created secret successfully. SecretName: '${cosi_secret_name}', Namespace: '${cosi_secret_namespace}'.\nThis secret should be referenced in the BucketClass and BucketAccessClass objects, using the keys 'cosiUserSecretName' and 'cosiUserSecretNamespace' under the 'parameters' field."
    fi
}

############################################################
############################################################
# Main program                                             #
############################################################
############################################################

# Set variables
cosi_secret_namespace="default"
cosi_secret_name=""
default_secret_name=true
s3_creds=()
glcp_creds=()

############################################################
# Process the input options.                               #
############################################################
# Get the options
# Options s, g, d and c require arguments.
opt_short="hs:e:g:d:c:"
opt_long="help,s3_credentials:,s3_endpoint:,glcp_credentials:,dscc_zone:,cluster_serial_number:,cosi_secret_name:,cosi_secret_namespace:"
options=$(getopt -o "$opt_short" --long "$opt_long" -- "$@")

# If invalid or no options have been received, print Help
if [[ $? -gt 0 || $# -eq 0 ]]
then
    help
fi

eval set -- ${options}
while [ $# -gt 0 ]
do
option="$1"
    case $option in
        -h|--help) # display Help
            help
            exit;;
        -s|--s3_credentials) # Enter the s3 access key and secret key
            # Split the argument string using the comma (,) delimiter
            IFS=',' read -ra S3_CREDS <<< "$2"
            for i in "${S3_CREDS[@]}"; do
                # append the string to the array
                s3_creds+=($i)
            done;
            shift
            ;;
        -e|--s3_endpoint) # Enter the s3 endpoint
            s3_endpoint=$2;
            shift
            ;;
        -g|--glcp_credentials) # Enter the GLCP user client ID and secret key
            # Split the argument string using the comma (,) delimiter
            IFS=',' read -ra GLCP_CREDS <<< "$2"
            for i in "${GLCP_CREDS[@]}"; do
                # append the string to the array
                glcp_creds+=($i)
            done;
            shift
            ;;
        -d|--dscc_zone) # Enter the DSCC zone
            dscc_zone=$2;
            shift
            ;;
        -c|--cluster_serial_number) # Enter the cluster serial number
            cluster_serial_number=$2;
            shift
            ;;
        --cosi_secret_name) # Enter the COSI secret name
            default_secret_name=false
            cosi_secret_name=$2;
            shift
            ;;
        --cosi_secret_namespace) # Enter the COSI secret namespace
            cosi_secret_namespace=$2;
            shift
            ;;
        --) # do nothing
            ;;
        *) # Invalid option
            echo "$0: unexpected parameter $1" >&2;
            exit 1;;
    esac
    shift
done

# If all mandatory options are not present, return error.
if [[ "${#s3_creds[@]}" -ne 2 || ! "$s3_endpoint" || "${#glcp_creds[@]}" -ne 2 || ! "$dscc_zone" || ! "$cluster_serial_number" ]]
then
    echo "Options -s, -e, -g, -d and -c are required. Please refer Help for more details."
    help
fi

if [[ $default_secret_name = true ]]
then
    # Kubernetes secret name does not support upper case letters, the cluster serial number is converted to lower case before appending.
    cosi_secret_name="cosi-user-secret-$(echo $cluster_serial_number | tr '[:upper:]' '[:lower:]')"
    echo "Secret name has not been provided in the arguments. Using auto-generated name: '$cosi_secret_name'"
fi

############################################################
# Prechecks.                                               #
############################################################

# Check if the kubectl executable exists
if ! command -v kubectl &> /dev/null
then
    echo "ERROR: kubectl could not be found. Please install kubectl and try again."
    exit 1
fi
# Check if the user has permissions to execute kubectl and if $PATH has been set properly.
if ! [ -x "$(command -v kubectl)" ]
then
    echo "ERROR: Could not execute kubectl. Please check if the binary is executable and is added to \$PATH."
    exit 1
fi
# Check if the kubectl commands work.
if ! err=$(kubectl get ns 2>&1)
then
    echo -e "ERROR: kubectl failed to connect to the cluster. See error message below for more details.\n$err"
    exit 1
fi

create_secret
