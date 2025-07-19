# Locally install
```sh
docker compose build
docker compose up -d
```
- go to http://localhost:3000/
- log in with admin/admin
- adding loki as a source in grafana (http://loki:3100)
- go to explore (loki) and run a query

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

# Promtail
Promtail is an open-source log collection agent created by Grafana Labs, designed to work seamlessly with Loki, Grafana’s log aggregation system. Promtail’s main role is to gather logs from various sources, process them, and send them to a Loki server, where they can be queried and visualized in Grafana.