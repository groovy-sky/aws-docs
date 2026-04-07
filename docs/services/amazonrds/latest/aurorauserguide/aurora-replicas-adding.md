# Adding Aurora Replicas to a DB cluster

An Aurora DB cluster with replication has one primary DB instance and up to 15 Aurora Replicas. The primary DB
instance supports read and write operations, and performs all data modifications to the cluster volume. Aurora Replicas connect to
the same storage volume as the primary DB instance, but support read operations only. You use Aurora Replicas to offload read
workloads from the primary DB instance. For more information, see [Aurora Replicas](aurora-replication.md#Aurora.Replication.Replicas).

Amazon Aurora Replicas have the following limitations:

- You can't create an Aurora Replica for an Aurora Serverless v1 DB cluster. Aurora Serverless v1 has a single DB instance that scales
up and down automatically to support all database read and write operations.

However, you can add reader instances to Aurora Serverless v2 DB clusters. For more information, see [Adding an Aurora Serverless v2 reader](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2-administration.html#aurora-serverless-v2-adding-reader).

We recommend that you distribute the primary instance and Aurora Replicas of your Aurora DB cluster
over multiple Availability Zones to improve the availability of your DB
cluster. For more information, see [Region availability](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html#Aurora.Overview.Availability).

To remove an Aurora Replica from an Aurora DB cluster, delete the Aurora Replica by following the
instructions in [Deleting a DB instance from an Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_DeleteCluster.html#USER_DeleteInstance).

###### Note

Amazon Aurora also supports replication with an external database, such as an RDS DB
instance. The RDS DB instance must be in the same AWS Region as Amazon Aurora. For more
information, see [Replication with Amazon Aurora](aurora-replication.md).

You can add Aurora Replicas to a DB cluster using the AWS Management Console, the AWS CLI, or the RDS API.

###### To add an Aurora replica to a DB cluster

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Databases**,
    and then select the DB cluster where you want to add the new DB instance.

3. Make sure that both the cluster and the primary instance are in the
    **Available** state. If the DB cluster or the primary instance
    are in a transitional state such as **Creating**, you can't
    add a replica.

    If the cluster doesn't have a primary instance, create one using the
    [create-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html) AWS CLI command.
    This situation can arise if you used the CLI to restore a DB cluster snapshot and then
    view the cluster in the AWS Management Console.

4. For **Actions**,
    choose **Add reader**.

The **Add reader** page appears.

5. On the **Add reader** page, specify options for your Aurora Replica.
    The following table shows settings for an Aurora Replica.

For this optionDo this

**Availability zone**

Determine if you want to specify a particular Availability Zone. The list includes only those
Availability Zones that are mapped to the DB subnet group that you chose when you created the DB
cluster. For more information about Availability Zones, see [Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html).

**Publicly accessible**

Select `Yes` to give the Aurora Replica a
public IP address; otherwise, select `No`.
For more information about hiding Aurora Replicas from
public access, see [Hiding a DB cluster in a VPC from the internet](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.Hiding).

**Encryption**

Select `Enable encryption` to enable encryption at rest
for this Aurora Replica. For more information, see [Encrypting Amazon Aurora resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Overview.Encryption.html).

**DB instance class**

Select a DB instance class that defines the processing
and memory requirements for the Aurora Replica. For
more information about DB instance class options, see
[Amazon AuroraDB instance classes](concepts-dbinstanceclass.md).

**Aurora replica source**

Select the identifier of the primary instance to
create an Aurora Replica for.

**DB instance identifier**

Enter a name for the instance that is unique for your
account in the AWS Region you selected. You might choose
to add some intelligence to the name such as including
the AWS Region and DB engine you selected, for example
`aurora-read-instance1`.

**Priority**

Choose a failover priority for the instance. If you
don't select a value, the default is
**tier-1**. This priority
determines the order in which Aurora Replicas are
promoted when recovering from a primary instance
failure. For more information, see [Fault tolerance for an Aurora DB cluster](concepts-aurorahighavailability.md#Aurora.Managing.FaultTolerance).

**Database port**

The port for an Aurora Replica is the same as the port
for the DB cluster.

**DB parameter group**

Select a parameter group. Aurora has a default
parameter group you can use, or you can create your own
parameter group. For more information about parameter
groups, see [Parameter groups for Amazon Aurora](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html).

**Performance Insights**

The **Turn on Performance Insights** check box is selected by default. The
value isn't inherited from the writer instance. For more information, see [Monitoring DB load with Performance Insights on Amazon Aurora](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.html).

**Enhanced monitoring**

Choose **Enable enhanced monitoring** to enable gathering
metrics in real time for the operating system that your
DB cluster runs on. For more information, see [Monitoring OS metrics with Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Monitoring.OS.html).

**Monitoring Role**

Only available if **Enhanced**
**Monitoring** is set to
**Enable enhanced monitoring**. Choose the IAM role that
you created to permit Amazon RDS to communicate with
Amazon CloudWatch Logs for you, or choose **Default** to
have RDS create a role for you named
`rds-monitoring-role`. For more
information, see [Monitoring OS metrics with Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Monitoring.OS.html).

**Granularity**

Only available if **Enhanced**
**Monitoring** is set to
**Enable enhanced monitoring**. Set the interval, in
seconds, between when metrics are collected for your DB
cluster.

**Auto minor version upgrade**

Select **Enable auto minor version upgrade** if you want to enable your
Aurora DB cluster to receive minor DB Engine
version upgrades automatically when they become
available.

The **Auto minor version upgrade** setting applies to both
Aurora PostgreSQL and Aurora MySQL DB clusters. For Aurora MySQL 2.x clusters, this setting upgrades the
clusters to a maximum version of 2.07.2.

For more information about engine updates for Aurora PostgreSQL, see
[Database engine updates for Amazon Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.html).

For more information about engine updates for Aurora MySQL, see
[Database engine updates for Amazon Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html).

6. Choose **Add reader** to create the Aurora
    Replica.

To create an Aurora Replica in your DB cluster, run the
[create-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html) AWS CLI command.
Include the name of the DB cluster as the `--db-cluster-identifier` option.
You can optionally specify an Availability Zone for the Aurora Replica using the
`--availability-zone` parameter, as shown in the following
examples.

For example, the following command creates a new MySQL 5.7–compatible Aurora Replica named
`sample-instance-us-west-2a`.

For Linux, macOS, or Unix:

```nohighlight

aws rds create-db-instance --db-instance-identifier sample-instance-us-west-2a \
    --db-cluster-identifier sample-cluster --engine aurora-mysql --db-instance-class db.r5.large \
    --availability-zone us-west-2a
```

For Windows:

```nohighlight

aws rds create-db-instance --db-instance-identifier sample-instance-us-west-2a ^
    --db-cluster-identifier sample-cluster --engine aurora-mysql --db-instance-class db.r5.large ^
    --availability-zone us-west-2a
```

The following command creates a new MySQL 5.7–compatible Aurora Replica named
`sample-instance-us-west-2a`.

For Linux, macOS, or Unix:

```nohighlight

aws rds create-db-instance --db-instance-identifier sample-instance-us-west-2a \
    --db-cluster-identifier sample-cluster --engine aurora-mysql --db-instance-class db.r5.large \
    --availability-zone us-west-2a
```

For Windows:

```nohighlight

aws rds create-db-instance --db-instance-identifier sample-instance-us-west-2a ^
    --db-cluster-identifier sample-cluster --engine aurora --db-instance-class db.r5.large ^
    --availability-zone us-west-2a
```

The following command creates a new PostgreSQL-compatible Aurora Replica named
`sample-instance-us-west-2a`.

For Linux, macOS, or Unix:

```nohighlight

aws rds create-db-instance --db-instance-identifier sample-instance-us-west-2a \
    --db-cluster-identifier sample-cluster --engine aurora-postgresql --db-instance-class db.r5.large \
    --availability-zone us-west-2a
```

For Windows:

```nohighlight

aws rds create-db-instance --db-instance-identifier sample-instance-us-west-2a ^
    --db-cluster-identifier sample-cluster --engine aurora-postgresql --db-instance-class db.r5.large ^
    --availability-zone us-west-2a
```

To create an Aurora Replica in your DB cluster, call the
[CreateDBInstance](../../../../reference/amazonrds/latest/apireference/api-createdbinstance.md) operation.
Include the name of the DB cluster as the `DBClusterIdentifier` parameter.
You can optionally specify an Availability Zone for the Aurora Replica using the
`AvailabilityZone` parameter.

For information about Auto Scaling Amazon Aurora with Aurora replicas, see the following sections.

###### Topics

- [Amazon Aurora Auto Scaling with Aurora Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Integrating.AutoScaling.html)

- [Adding an auto scaling policy to an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Integrating.AutoScaling.Add.html)

- [Editing an auto scaling policy for an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Integrating.AutoScaling.Edit.html)

- [Deleting an auto scaling policy from your Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Integrating.AutoScaling.Delete.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Modifying an Aurora DB cluster

Auto Scaling with Aurora Replicas
