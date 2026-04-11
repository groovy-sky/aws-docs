# Attaching to the remote RDS for Db2 DB instance

Use the following steps to attach to your remote RDS for Db2 DB instance and run `get snapshot` operations.

###### To attach to the remote RDS for Db2 DB instance

1. Run a client-side IBM Db2 CLP session. For information about cataloging your RDS for Db2
    DB instance and database, see [Connecting to your Amazon RDS for Db2 DB instance with IBM Db2 CLP](db2-connecting-with-clp-client.md). Make a note of the master
    username and master password for your RDS for Db2 DB instance.

2. Attach to the RDS for Db2 DB instance. In the following example, replace
    `node_name`,
    `master_username`, and
    `master_password` with the TCPIP node name that you
    catalogued and the master username and master password for your RDS for Db2 DB
    instance.

```nohighlight

db2 attach to node_name user master_username using master_password
```

After attaching to the remote RDS for Db2 DB instance, you can run the following commands and
other `get snapshot` commands. For more information, see [GET\
SNAPSHOT command](https://www.ibm.com/docs/en/db2/11.5?topic=commands-get-snapshot) in the IBM Db2 documentation.

```nohighlight

db2 list applications
db2 get snapshot for all databases
db2 get snapshot for database manager
db2 get snapshot for all applications
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Granting and revoking privileges

Database
tasks

All content copied from https://docs.aws.amazon.com/.
