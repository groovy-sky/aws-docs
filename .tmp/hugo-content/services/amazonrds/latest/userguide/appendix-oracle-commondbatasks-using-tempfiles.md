# Working with tempfiles in RDS for Oracle

## Adding a tempfile to the instance store on a read replica

When you create a temporary tablespace on a primary DB instance, the read
replica doesn't create tempfiles. Assume that an empty temporary tablespace
exists on your read replica for either of the following reasons:

- You dropped a tempfile from the tablespace on your read replica.
For more information, see [Dropping tempfiles on a read replica](appendix-oracle-commondbatasks-dropping-tempfiles-replica.md).

- You created a new temporary tablespace on the primary DB instance. In
this case, RDS for Oracle synchronizes the metadata to the read
replica.

You can add a tempfile to the empty temporary tablespace, and store the
tempfile in the instance store. To create a tempfile in the instance store,
use the Amazon RDS procedure
`rdsadmin.rdsadmin_util.add_inst_store_tempfile`. You can use
this procedure only on a read replica. The procedure has the following
parameters.

Parameter nameData typeDefaultRequiredDescription

`p_tablespace_name`

varchar

—

Yes

The name of the temporary tablespace on your read
replica.

In the following example, the empty temporary tablespace
`temp01` exists on your read replica. Run the
following command to create a tempfile for this tablespace, and store it in
the instance store.

```sql

EXEC rdsadmin.rdsadmin_util.add_inst_store_tempfile(p_tablespace_name => 'temp01');
```

For more information, see [Storing temporary data in an RDS for Oracle instance store](chap-oracle-advanced-features-instance-store.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Working with
Oracle tablespaces

Dropping tempfiles on a read replica

All content copied from https://docs.aws.amazon.com/.
