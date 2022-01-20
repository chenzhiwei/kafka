# Kafka

## Kafka Server

Suppose Zookeeper cluster is ready.

### Create Kafka admin and inter-broker-admin users

```
export KAFKA_OPTS="-Djava.security.auth.login.conf=/etc/kafka/zookeeper-client.jaas.conf"

bin/kafka-configs.sh --zookeeper zk1:2182 --zk-tls-config-file zookeeper-client.properties \
    --alter --add-config 'SCRAM-SHA-256=[password=admin-password]' --entity-type users \
    --entity-name admin

bin/kafka-configs.sh --zookeeper zk1:2182 --zk-tls-config-file zookeeper-client.properties \
    --alter --add-config 'SCRAM-SHA-256=[password=inter-broker-password]' --entity-type users \
    --entity-name inter-broker-admin
```

### Start Kafka

```
export KAFKA_OPTS="-Djava.security.auth.login.conf=/etc/kafka/kafka-server.jaas.conf"
bin/kafka-server-start.sh /etc/kafka/server.properties
```

## Kafka Client

```
bin/kafka-topics.sh --command-config kafka-client.properties \
    --bootstrap-server kafka1:9093 \
    --create \
    --config min.insync.replicas=2 \
    --partition 3 \
    --replication-factor 3 \
    --topic topicName

bin/kafka-configs.sh --command-config kafka-client.properties \
    --bootstrap-server kafka1:9093 \
    --alter --add-config 'SCRAM-SHA-256=[password=password]' --entity-type users \
    --entity-name alice

bin/kafka-acls.sh --command-config kafka-client.properties \
    --bootstrap-server kafka1:9093 \
    --add --allow-principal User:alice --operation All --topic topicName
```
