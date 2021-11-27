# Kafka

## Start Kafka

```
IPADDR=192.168.122.10  # Your Host IP address
podman run -d --name kafka -p 9092:9092 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://${IPADDR}:9092 \
    docker.io/siji/kafka:3.0.0
```

## Publish Messages

```
kafka-producer -file example.yaml
```
