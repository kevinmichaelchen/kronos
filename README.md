# kronos

This project provides a gRPC API for analytics.

## FAQ

### Where does the name come from?
Its name is a cooler misspelling of the [Greek god of time](https://en.wikipedia.org/wiki/Chronos).  

### Why might we use this?
- Segment is for customer data.
- Bigtable is probably cheaper than a 3rd party offering.

### Why Bigtable?
- [Bigtable](https://en.wikipedia.org/wiki/Bigtable) is a sparsely populated table 
that can scale to billions of rows, supports high read and write throughput at low latency, 
and is an ideal source for MapReduce operations.
- Using [Cloud Bigtable](https://cloud.google.com/bigtable/docs/overview) means Google handles
upgrades, restarts, durability, auto-scaling, and cross-regional replication for us. 

## Running
We use the [gcloud emulator](https://cloud.google.com/sdk/gcloud/reference/beta/emulators/bigtable/)
to run locally.
```
# In one tab
gcloud beta emulators bigtable start

# In a separate tab
make
```

### Creating tables (run each time emulator restarts)
```
go get -u cloud.google.com/go/bigtable/cmd/cbt

# Locate the cbt command.
# It should install under $GOPATH/bin.
# If you don't have GOPATH set, it defaults to $HOME/go.

env BIGTABLE_EMULATOR_HOST=localhost:8086 \
  ~/go/bin/cbt \
    -project=my-project \
    -instance=my-instance \
    createtable mytable

env BIGTABLE_EMULATOR_HOST=localhost:8086 \
  ~/go/bin/cbt \
    -project=my-project \
    -instance=my-instance \
    createfamily mytable event
```

### Hitting the gRPC API
#### Sending an event
```
grpcurl -v -plaintext \
  -d '{"event": "login", "userID": "f78002f4-873d-4e79-bf13-0453c4951312", "properties": {"a": "b"}}' \
  :8080 proto.EventService/SendEvent
```

#### Reading all events (debugging only)
```
grpcurl -v -plaintext -d '{}' :8080 proto.EventService/ReadEvents
```