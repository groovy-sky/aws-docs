# Displaying volume status for an Aurora PostgreSQL DB cluster

In Amazon Aurora, a DB cluster volume consists of a collection of logical blocks. Each of
these represents 10 gigabytes of allocated storage. These blocks are called
_protection groups_.

The data in each protection group is replicated across six physical storage devices,
called _storage nodes_. These storage nodes are allocated across
three Availability Zones (AZs) in the region where the DB cluster resides. In turn, each
storage node contains one or more logical blocks of data for the DB cluster volume. For
more information about protection groups and storage nodes, see [Introducing the Aurora storage\
engine](https://aws.amazon.com/blogs/database/introducing-the-aurora-storage-engine) on the AWS Database Blog. To learn more about Aurora cluster volumes
in general, see [Amazon Aurora storage](aurora-overview-storagereliability.md).

Use the `aurora_show_volume_status()` function to return the following
server status variables:

- `Disks` — The total number of logical blocks of data for the
DB cluster volume.

- `Nodes`— The total number of storage nodes for the
DB cluster volume.

You can use the `aurora_show_volume_status()` function to help avoid an
error when using the `aurora_inject_disk_failure()` fault injection function.
The `aurora_inject_disk_failure()` fault injection function simulates the
failure of an entire storage node, or a single logical block of data within a storage
node. In the function, you specify the index value of a specific logical block of data
or storage node. However, the statement returns an error if you specify an index value
greater than the number of logical blocks of data or storage nodes used by the DB
cluster volume. For more information about fault injection queries, see [Testing Amazon Aurora PostgreSQL by using fault injection queries](aurorapostgresql-managing-faultinjectionqueries.md).

###### Note

The `aurora_show_volume_status()` function is available for
Aurora PostgreSQL version 10.11. For more information about Aurora PostgreSQL versions, see
[Amazon Aurora PostgreSQL releases and engine versions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html).

**Syntax**

```nohighlight

SELECT * FROM aurora_show_volume_status();
```

**Example**

```sql

customer_database=> SELECT * FROM aurora_show_volume_status();
 disks | nodes
-------+-------
    96 |    45
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Testing Amazon Aurora PostgreSQL by using fault injection queries

Specifying the RAM disk for the stats\_temp\_directory
