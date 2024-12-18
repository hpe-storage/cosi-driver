# © Copyright 2024 Hewlett Packard Enterprise Development LP

#!/bin/bash
#
# This script is responsible for collecting logs in the HPE COSI driver.
# It gathers relevant log files and stores them in a specified location.
# The collected logs can be used for troubleshooting and debugging purposes.
#
# Collects HPE COSI controller sidecar container logs.
provisioner_log_collection() {

pod_list=$(kubectl get pods -n $namespace --selector=app.kubernetes.io/part-of=container-object-storage-interface -o jsonpath='{.items..metadata.name}')
timestamp=`date '+%Y%m%d_%H%M%S'`
tmp_log_dir="/var/log/hpe-cosi-provisioner-logs-$timestamp"
temp_cr_log_dir="$temp_log_dir/hpe-cosi-cr-logs"
dest_log_dir="/var/log"
hostname=$(cat /etc/hostname)
tar_file_name="hpe-cosi-provisioner-logs-$hostname-$timestamp.tar.gz"

mkdir -p $tmp_log_dir
if [[ ! -z "$pod_list" ]]
then
	for pod_name in $pod_list
	do
		container_list=$(kubectl get pod $pod_name -n $namespace -o jsonpath='{.spec.containers[*].name}')
		if [[ ! -z "$container_list" ]]
		then
			for container_name in $container_list
			do
				timeout 30 kubectl logs $pod_name -n $namespace -c $container_name &> $tmp_log_dir/$pod_name.$container_name.log
				# Get the restart count for the container in the pod
				restart_count=$(kubectl get pod $pod_name -n $namespace -o json | jq -r ".status.containerStatuses[] | select(.name==\"$container_name\") | .restartCount")
				if [[ $restart_count > 0 ]]
				then
					timeout 30 kubectl logs $pod_name -n $namespace -c $container_name --previous &> $tmp_log_dir/$pod_name.$container_name.before_restart.log
				fi		
			done
		fi
	done
fi

# Collect the events log
collect_events_logs $tmp_log_dir

# collect cr logs
collect_cr_logs $tmp_log_dir

if [[ ! -z $(ls $tmp_log_dir) ]]
then
	tar -czf $tar_file_name -C $tmp_log_dir . &> /dev/null
	mv $tar_file_name $dest_log_dir &> /dev/null
fi
	
rm -rf $tmp_log_dir

if [[ -f "$dest_log_dir/$tar_file_name" ]]
then
	echo "HPE COSI provisioner logs were collected into $dest_log_dir/$tar_file_name on host $hostname."
else
	echo "Unable to collect HPE COSI provisioner log files."
fi

}

collect_cr_logs() {

	tmp_log_dir=$1
	temp_cr_log_dir="$temp_log_dir/hpe-cosi-cr-logs"

	get_bucket_list=$(kubectl get bucket -n $namespace -o jsonpath='{.items..metadata.name}')
	get_bucket_class_list=$(kubectl get bucketclass -n $namespace -o jsonpath='{.items..metadata.name}')
	get_bucket_claim_list=$(kubectl get bucketclaim -n $namespace -o jsonpath='{.items..metadata.name}')
	get_bucket_access_list=$(kubectl get bucketaccess -n $namespace -o jsonpath='{.items..metadata.name}')
	get_bucket_access_class_list=$(kubectl get bucketaccessclass -n $namespace -o jsonpath='{.items..metadata.name}')

	if [[ ! -z "$get_bucket_list" || ! -z "$get_bucket_class_list" || ! -z "$get_bucket_claim_list" || ! -z "$get_bucket_access_list" || ! -z "$get_bucket_access_class_list" ]]
	then
		mkdir -p $temp_cr_log_dir
	fi

	for bucket in $get_bucket_list
	do
		timeout 30 kubectl describe bucket $bucket -n $namespace &> /$temp_cr_log_dir/hpe-cosi-bucket-$bucket.log
	done

	for bucket_class in $get_bucket_class_list
	do
		timeout 30 kubectl describe bucketclass $bucket_class -n $namespace &> /$temp_cr_log_dir/hpe-cosi-bucketclass-$bucket_class.log
	done

	for bucket_claim in $get_bucket_claim_list
	do
		timeout 30 kubectl describe bucketclaim $bucket_claim -n $namespace &> /$temp_cr_log_dir/hpe-cosi-bucketclaim-$bucket_claim.log
	done

	for bucket_access in $get_bucket_access_list
	do
		timeout 30 kubectl describe bucketaccess $bucket_access -n $namespace &> /$temp_cr_log_dir/hpe-cosi-bucketaccess-$bucket_access.log
	done

    for bucket_access_class in $get_bucket_access_class_list
	do
		timeout 30 kubectl describe bucketaccessclass $bucket_access_class -n $namespace &> /$temp_cr_log_dir/hpe-cosi-bucketaccessclass-$bucket_access_class.log
	done
}

collect_events_logs() {
	temp_log_dir=$1
	timeout 30 kubectl get events -n $namespace --sort-by='.metadata.creationTimestamp' -o json &> /$temp_log_dir/hpe-cosi-events.log
}

display_usage() {
echo "Collect HPE storage diagnostic logs using kubectl."
echo -e "\nUsage:"
echo -e "     hpe-logcollector.sh [-h|--help] [-n|--namespace NAMESPACE]"
echo -e "Options:"
echo -e "-h|--help                  Print this usage text"
echo -e "-n|--namespace NAMESPACE   Collect logs from HPE COSI deployment in namespace"
echo -e "                           NAMESPACE (default: default)\n"
exit 0

}
namespace="default"
#Main Function
if ! options=$(getopt -o han: -l help,namespace: -- "$@")
then
    exit 1
fi

eval set -- $options
while [ $# -gt 0 ]
do
key="$1"
    case $key in
    -h|--help) display_usage; break;;
    -n|--namespace) namespace=$2; shift;;
    --) ;;
    *) echo "$0: unexpected parameter $1" >&2; exit 1;;
    esac
    shift
done

provisioner_log_collection
