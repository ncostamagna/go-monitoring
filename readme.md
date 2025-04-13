# AWS commands

```
# see instances
aws ec2 describe-instances --profile costamagna-admin --region us-east-1 --query "Reservations[*].Instances[*].[Tags[?Key=='Name']|[0].Value, InstanceId, InstanceType, State.Name, PublicIpAddress]" --output table 

# start instance 
aws ec2 start-instances --profile costamagna-admin --region us-east-1 --instance-ids i-02f74c58261cfed31

# stop instance
aws ec2 stop-instances --profile costamagna-admin --region us-east-1 --instance-ids {instance-id}

```

# Doc
notion page: https://www.notion.so/Prometheus-17630008732a80939106c2eded1cae01