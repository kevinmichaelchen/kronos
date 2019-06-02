## Bigtable Overview
From this [blog post](https://syslog.ravelin.com/the-joy-and-pain-of-using-google-bigtable-4210604c75be):
> First, a quick primer on Bigtable:
>
> Bigtable is essentially a giant, sorted, 3 dimensional map. 
> The first dimension is the row key.
> One can look up any row given a row key very quickly.
> You can also scan rows in alphabetical order quickly.
> You can start and end the scan at any given place.
> One caveat is you can only scan one way. Which is annoying. But ho hum.
>
> The second dimension are columns within a row. 
> Columns are keyed by a string and can be segmented into column families.
> From what we have seen with our use of column families, they are not much more than a prefix on your column key.
> You can filter columns individually or via column families to return only those columns you are interested in.
>
> Finally, every column contains a set of cells.
> Cells hold your values which are just bytes.
> This is the third dimension of our 3D map.
> Cells are indexed by a timestamp in milliseconds.
> You can filter your request to ask for the latest cell or a custom range of cells.
> So any particular value in Bigtable is identified by its row key, column key (including column family) and cell timestamp.
>
> Congratulations, you are now a Bigtable expert!

## Choosing a row key
The only way to query Cloud Bigtable efficiently is by row key.
> When Cloud Bigtable stores rows, it sorts them by row key in lexicographic order. 
> There is effectively a single index per table, which is the row key. 
> Queries that access a single row, or a contiguous range of rows, execute quickly and efficiently. 
> All other queries result in a full table scan, which will be far, far slower.
>
> Choosing a row key that facilitates common queries is of paramount importance to the overall performance of the system. 
> Enumerate your queries, put them in order of importance, and then design row keys that work for those queries.

So [choosing a row key is very important](https://cloud.google.com/bigtable/docs/schema-design#row-keys).
It's important for row keys to distribute, otherwise most writes will overload a single node.
Since our userIDs are UUIDs, and UUIDs are pretty regularly distributed and avoid "hotspotting", it makes sense
to start the row key with the userID.
```
userID#epochMS
```
If we usually retrieve most recent records first, we can use a reverse timestamp (`math.MaxInt64 - timestamp`).
Note: we shouldn't use a hyphen to delimit values in the row key, since UUIDs already contain hyphens.

Using `userID#epochMS` is great for querying all events related to a user in March, for example.

Sometimes, we have to the same data more than once to facilitate different queries.
> Use two tables, each with a row key appropriate to one of the queries.
> This is a good solution, because it results in a robust, scalable system.

## Common queries
Here are some ideas for row key, and what query they'd be useful for.

| Query (find ...)                                      | Row key                             |
| ----------------------------------------------------- |:-----------------------------------:|
| all events for Alice in March                         | `userID#epochMS`                    |
| all events for a device in March                      | `deviceID#epochMS`                  |
| all of Alice's logins in March                        | `userID#eventType#epochMS`          |
| all of Alice's logins on a particular device in March | `userID#eventType#deviceID#epochMS` |
| all times file X was launched in March                | `fileID#eventType#epochMS`          |
| number of logins in last day/week/month               | `eventType#epochMS`                 |
| how long Alice spent in VR in last day/week           | `userID#eventType#epochMS`          |
| how long all users spent in VR in last day/week       | `userID#eventType#epochMS`          |

We only have a few event types, e.g., `login`, `logout`, `file-launched`, `file-closed`.

To figure out how long Alice spent in VR in last day/week, we can maybe use a column family of the event type,
and maybe use a reverse timestamp (so most recent events come first). For this to work, logout would need to be
disabled for connectivity...

Or we can have clients send a heartbeat every 5 seconds. That way, getting a duration is boiled down
to summing a count of heartbeats and multiplying by 5. This only works if client device is connected though...

The heartbeat approach is probably on the right track though.
Whether connected or disconnected, the client periodically writes heartbeats to a heartbeat file.
The client then sends the heartbeat events when connected.
Thus, we can use the row key `userID#eventType#epochMS` to find Alice's time in VR.

Not only do heartbeat events make answering this query easier,
but if we didn't use heartbeats, then disconnected logout wouldn't be possible,
and the logout network call would have to be retried until successful.