# Turning off logical replication

After completing your replication tasks, you should stop the replication process, drop
replication slots, and turn off logical replication. Before dropping slots, make sure
that they're no longer needed. Active replication slots can't be dropped.

###### To turn off logical replication

1. Drop all replication slots.

To drop all of the replication slots, connect to the publisher and run the
    following SQL command.

```nohighlight

SELECT pg_drop_replication_slot(slot_name)
     FROM pg_replication_slots
    WHERE slot_name IN (SELECT slot_name FROM pg_replication_slots);
```

The replication slots can't be active when you run this command.

2. Modify the custom DB cluster parameter group associated with the publisher as
    detailed in [Setting up logical replication for your Aurora PostgreSQL DB cluster](aurorapostgresql-replication-logical-configure.md), but set
    the `rds.logical_replication` parameter to 0.

For more information about custom parameter groups, see [Modifying parameters in a DB cluster parameter groupin Amazon Aurora](user-workingwithparamgroups-modifyingcluster.md).

3. Restart the publisher Aurora PostgreSQL DB cluster for the change to the
    `rds.logical_replication` parameter to take effect.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting up
logical replication

Monitoring
the write-through cache and logical slots

All content copied from https://docs.aws.amazon.com/.
