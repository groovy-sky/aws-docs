# Renaming a DB instance in an active-active cluster

You can change the name of a DB instance in an active-active cluster. To rename more than one DB instance in an
active-active cluster, do so one DB instance at a time. So, rename one DB instance and rejoin it to the cluster
before you rename the next DB instance.

###### To rename a DB instance in an active-active cluster

1. Connect to the DB instance in a SQL client, and call the [mysql.rds\_group\_replication\_stop](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-stored-proc-active-active-clusters.html#mysql_rds_group_replication_stop) stored procedure:

```nohighlight

call mysql.rds_group_replication_stop();
```

2. Rename the DB instance by following the instructions in [Renaming a DB instance](user-renameinstance.md).

3. Modify the `group_replication_group_seeds` parameter in each DB parameter group associated with
    a DB instance in the active-active cluster.

In the parameter setting, replace the old DB instance endpoint with the new DB instance endpoint.
    For more information about setting parameters, see [Modifying parameters in a DB parameter group in Amazon RDS](user-workingwithparamgroups-modifying.md).

4. Connect to the DB instance in a SQL client, and call the [mysql.rds\_group\_replication\_start](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-stored-proc-active-active-clusters.html#mysql_rds_group_replication_start) stored procedure:

```nohighlight

call mysql.rds_group_replication_start(0);
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Stopping Group
Replication in an active-active cluster

Removing a DB instance from an active-active cluster
