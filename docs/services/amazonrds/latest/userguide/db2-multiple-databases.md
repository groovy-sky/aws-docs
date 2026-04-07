# Multiple databases on an Amazon RDS for Db2 DB instance

You can create multiple databases on a single RDS for Db2 DB instance by calling the [rdsadmin.create\_database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-sp-managing-databases.html#db2-sp-create-database) stored procedure. A
single RDS for Db2 DB instance is limited to 50 databases. This number includes databases in
both activated and deactivated states.

###### Note

If you create multiple databases on an RDS for Db2 DB instance that was created before
November 15, 2024, then you must reboot the DB instance to enable support for multiple
databases.

By default, Amazon RDS activates databases when you create them. To optimize memory resources,
you can deactivate databases that you use infrequently and then activate them later when
needed. For more information, see [Deactivating a database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-managing-databases.html#db2-deactivating-database) and [Activating a database](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-managing-databases.html#db2-activating-database).

The number of activated databases on a DB instance depends on the available memory
resources on the server. Memory resources differ based on the DB instance class and the
amount of memory configured for the database. For information about DB instance classes, see
[DB instance classes](concepts-dbinstanceclass.md). For information about how to update the
memory for an RDS for Db2 database, see [rdsadmin.update\_db\_param](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-sp-managing-databases.html#db2-sp-update-db-param).

We recommend that you choose a DB instance class that has 2 GB of memory for common
database tasks, operating system requirements, and other Amazon RDS automation tasks such as
backups. For more information about changing the DB instance class, see [Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

In addition, IBM recommends a minimum of 1 GB of memory for each active database. For more
information, see [Disk\
and memory requirements](https://www.ibm.com/docs/en/db2/11.5?topic=servers-disk-memory-requirements) in the IBM documentation.

You can calculate the maximum number of active databases a DB instance can have with the
following formula:

```nohighlight

Active database limit = (total server memory - 2 GB) / 1 GB
```

The following example shows the maximum number of active databases for a DB instance with
a db.m6i.xlarge DB instance class:

```nohighlight

Active database limit = (total server memory - 2 GB) / 1 GB
                      = (16 GB - 2 GB) / 1 GB
                      = 14 databases
```

When Amazon RDS recovers a database after a crash, it activates the database if it was
previously active. In certain cases, such as when you modify a DB instance class to a lower
memory configuration, there might be insufficient memory to activate all databases on the DB
instance. In those cases, Amazon RDS activates databases in the order in which they were
created.

###### Note

Any databases that Amazon RDS can't activate because of insufficient memory remain in a
deactivated state.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DB instance
prerequisites

Connecting to your Db2 DB
instance
