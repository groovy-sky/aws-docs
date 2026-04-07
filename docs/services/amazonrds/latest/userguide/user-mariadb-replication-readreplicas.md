# Working with MariaDB read replicas

Following, you can find specific information about working with read replicas on
Amazon RDS for MariaDB. For general information about read replicas and instructions for using
them, see [Working with DB instance read replicas](user-readrepl.md).

- [Configuring replication filters with MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.ReplicationFilters.html)

- [Configuring delayed replication with MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.DelayReplication.html)

- [Updating read replicas with MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.Updates.html)

- [Working with Multi-AZ read replica deployments with MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.MultiAZ.html)

- [Using cascading read replicas with RDS for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.Cascading.html)

- [Monitoring MariaDB read replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.Monitor.html)

- [Starting and stopping replication with MariaDB read replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.StartStop.html)

- [Troubleshooting a MariaDB read replica problem](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.Troubleshooting.MariaDB.html)

## Configuring read replicas with MariaDB

Before a MariaDB DB instance can serve as a replication source, make sure to turn on
automatic backups on the source DB instance by setting the backup retention period
to a value other than 0. This requirement also applies to a read replica that is the
source DB instance for another read replica.

You can create up to 15 read replicas from one DB instance within the same Region. For replication to operate
effectively, each read replica should have as the same amount of compute and storage
resources as the source DB instance. If you scale the source DB instance, also scale
the read replicas.

RDS for MariaDB supports cascading read replicas. To learn how to configure cascading read replicas,
see [Using cascading read replicas with RDS for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.ReadReplicas.Cascading.html).

You can run multiple read replica create and delete actions at the same time that reference
the same source DB instance. When you perform these actions, stay within the limit of 15 read replicas
for each source instance.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MariaDB replication

Configuring replication filters
