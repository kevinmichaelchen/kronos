# kronos
This project provides a gRPC API for analytics.

## FAQ

### Where does the name come from?
Its name is a cooler misspelling of the [Greek god of time](https://en.wikipedia.org/wiki/Chronos).  

### Why might we use this?
- Analytics are a big part of our workflow.
  - We have a [JIRA ticket](https://irisvr.atlassian.net/browse/PROS-441) for implementing events on the Quest.
- Segment is for customer data.
- Bigtable is probably cheaper than a 3rd party offering.

### Why Bigtable?
- [Bigtable](https://en.wikipedia.org/wiki/Bigtable) is a sparsely populated table 
that can scale to billions of rows, supports high read and write throughput at low latency, 
and is an ideal source for MapReduce operations.
- Using [Cloud Bigtable](https://cloud.google.com/bigtable/docs/overview) means Google handles
upgrades, restarts, durability, auto-scaling, and cross-regional replication for us. 
- Bigtable pairs well with Dataflow to form a [streaming analytics pipeline](https://cloud.google.com/solutions/big-data/stream-analytics/).

### What queries have been tested so far?


### Concerns
- How do we stop malicious actors from sending fabricated events? We'd probably use an API key similar to what Segment does.
- How do we migrate existing data from Segment into Cloud Bigtable?
- How do we port old Prospect builds to forward analytics events to this service instead of Segment?

## Running
### Start the emulator
We use the [gcloud emulator](https://cloud.google.com/sdk/gcloud/reference/beta/emulators/bigtable/)
to run locally.
```
# In one tab
gcloud beta emulators bigtable start
```

### Run tests
```
make testv
```

### Run the app
```
make
```

### Hitting the gRPC API
#### Sending a login event
```
grpcurl -v -plaintext \
  -d '{"userID": "f78002f4-873d-4e79-bf13-0453c4951312", "timeMs": 1559520445749, "properties": {"a": "b"}}' \
  :8080 proto.EventService/SendLoginEvent

grpcurl -v -plaintext \
  -d '{"userID": "f78002f4-873d-4e79-bf13-0453c4951312", "timeMs": 1559520445750, "properties": {"a": "b"}}' \
  :8080 proto.EventService/SendLoginEvent
```

#### Counting logins for user
```
grpcurl -v -plaintext \
  -d '{"userID": "f78002f4-873d-4e79-bf13-0453c4951312", "start": 1559520445749, "end": 1559520640198}' \
  :8080 proto.EventService/GetNumberOfLogins
```

#### Reading all events (debugging only)
```
grpcurl -v -plaintext -d '{}' :8080 proto.EventService/ReadEvents
```
