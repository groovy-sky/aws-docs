# Turning Performance Insights on and off for Amazon RDS

###### Important

AWS has announced the end-of-life date for Performance Insights: June 30, 2026. After this date, Amazon RDS will no longer support the Performance Insights console experience,
flexible retention periods (1-24 months), and their associated pricing. The Performance Insights API will continue to exist with no pricing changes. Costs for the
Performance Insights API will appear in your AWS bill with the cost of CloudWatch Database Insights.

We recommend that you upgrade any DB instances
using the paid tier of Performance Insights to the Advanced mode of Database Insights before June 30, 2026.
For information about upgrading to the Advanced mode of Database Insights, see
[Turning on the Advanced mode of Database Insights for Amazon RDS](user-databaseinsights-turningonadvanced.md).

If you take no action, DB instances using Performance Insights
will default to using the Standard mode of Database Insights. With Standard mode of Database Insights, you might lose access to performance data history beyond 7 days and might not be able to use execution plans
and on-demand analysis features in the Amazon RDS console. After June 30, 2026 only the Advanced mode of Database Insights will support execution plans and on-demand analysis.

With CloudWatch Database Insights, you can monitor database load for your fleet of databases and analyze and troubleshoot performance at scale.
For more information about Database Insights, see [Monitoring Amazon RDS databases with CloudWatch Database Insights](user-databaseinsights.md).
For pricing information, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

You can turn on Performance Insights for your DB instance or Multi-AZ DB cluster
when you create it. If needed, you can turn it off later by modifying your
DB instance
from the console. Turning Performance Insights on and off
doesn't cause downtime, a reboot, or a failover.

###### Note

Performance Schema is an optional performance tool used by Amazon RDS for MariaDB or MySQL. If you turn Performance Schema on or off, you need
to reboot. If you turn Performance Insights on or off, however, you don't need to reboot. For more information,
see [Overview of the Performance Schema for Performance Insights on Amazon RDS for MariaDB or MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.EnableMySQL.html).

The Performance Insights agent consumes limited CPU and memory on the DB host. When
the DB load is high, the agent limits the performance impact by collecting data less
frequently.

Console

In the console, you can turn Performance Insights on or off when you create or modify a DB instance
or Multi-AZ DB cluster.

Turning Performance Insights on or off when creating a DB instance or Multi-AZ DB cluster

After creating a new DB instance or Multi-AZ DB cluster,
Amazon RDS enables Performance Insights by default. To turn off Performance Insights, choose the option
**Database Insights – Standard** and deselect the option **Enable Performance Insights**.

For more information, see the following topics.

- To create a DB instance, follow the instructions for your DB engine in [Creating an Amazon RDS DB instance](user-createdbinstance.md).

- To create a Multi-AZ DB cluster, follow the instructions for your DB engine in [Creating a Multi-AZ DB cluster for Amazon RDS](create-multi-az-db-cluster.md).

The following screenshot shows the **Performance Insights** section.

![Turn on Performance Insights during DB instance or Multi-AZ DB cluster creation with console](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/perf_insights_enabling.png)

If you choose **Enable Performance Insights**, you have the following options:

- **Retention** (for the Standard mode of Database Insights only) – The amount of time to retain Performance Insights data. The retention setting is **Default (7 days)**. To retain your performance
data for longer, specify 1–24 months. For more information about retention periods, see
[Pricing and data retention for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.Overview.cost.html).

- **AWS KMS key** – Specify your AWS KMS key.
Performance Insights encrypts all potentially sensitive
data using your KMS key. Data is encrypted in flight and at rest.
For more information, see [Changing an AWS KMS policy for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.cmk-policy.html).

Turning Performance Insights on or off when modifying a DB instance or Multi-AZ DB cluster

In the console, you can modify
a DB instance or Multi-AZ DB cluster to manage Performance Insights.

###### To manage Performance Insights for a DB instance or Multi-AZ DB cluster using the console

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. Choose **Databases**.

3. Choose a DB instance or Multi-AZ DB cluster,
    and choose **Modify**.

4. To turn on Performance Insights, select **Enable Performance Insights**. To turn off Performance Insights, choose the option
    **Database Insights – Standard** and deselect the option **Enable Performance Insights**.

If you choose **Enable Performance Insights**, you have the following options:

- **Retention** (for the Standard mode of Database Insights only) – The amount of time to retain Performance Insights data. The retention setting is **Default (7 days)**. To retain your performance
data for longer, specify 1–24 months. For more information about retention periods, see
[Pricing and data retention for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.Overview.cost.html).

- **AWS KMS key** – Specify your
KMS key. Performance Insights encrypts all
potentially sensitive data using your KMS key. Data is
encrypted in flight and at rest. For more information, see
[Encrypting Amazon RDS resources](overview-encryption.md).

5. Choose **Continue**.

6. For **Scheduling of Modifications**, choose Apply immediately. If you
    choose Apply during the next scheduled maintenance window, your instance
    ignores this setting and turns on Performance Insights
    immediately.

7. Choose **Modify instance**.

AWS CLI

When you use the [create-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html) AWS CLI command, turn on Performance Insights by
specifying `--enable-performance-insights` and set `--database-insights-mode` to either `advanced` or `standard`.
To turn off Performance Insights, specify `--no-enable-performance-insights` and set `database-insights-mode` to `standard`.

You can also specify these values using the following AWS CLI commands:

- [create-db-cluster](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-cluster.html)

- [modify-db-cluster](https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-cluster.html)

- [create-db-instance-read-replica](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance-read-replica.html)

- [modify-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-instance.html)

- [restore-db-instance-from-s3](https://docs.aws.amazon.com/cli/latest/reference/rds/restore-db-instance-from-s3.html)

When you turn on Performance Insights in the CLI, you can optionally specify the number of days to retain Performance Insights data with the
`--performance-insights-retention-period` option. You can specify `7`, `month` \\* 31 (where `month` is a number from 1–23),
or `731`. For example, if you want to retain your performance data for 3 months, specify `93`, which is 3 \* 31. The default
is `7` days. For more information about retention periods, see [Pricing and data retention for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.Overview.cost.html).

The following example turns on Performance Insights for `sample-db-cluster` and specifies that Performance Insights data is
retained for 93 days (3 months).

For Linux, macOS, or Unix:

```nohighlight

aws rds modify-db-cluster \
	--database-insights-mode standard \
    --db-cluster-identifier sample-db-instance \
    --enable-performance-insights \
    --performance-insights-retention-period 93
```

For Windows:

```nohighlight

aws rds modify-db-cluster ^
	--database-insights-mode standard ^
    --db-cluster-identifier sample-db-instance ^
    --enable-performance-insights ^
    --performance-insights-retention-period 93
```

If you specify a retention period such as 94 days, which isn't a valid value, RDS issues an error.

```

An error occurred (InvalidParameterValue) when calling the CreateDBInstance operation:
Invalid Performance Insights retention period. Valid values are: [7, 31, 62, 93, 124, 155, 186, 217,
248, 279, 310, 341, 372, 403, 434, 465, 496, 527, 558, 589, 620, 651, 682, 713, 731]
```

###### Note

You can only toggle Performance Insights for an instance in a DB cluster where Performance Insights is not managed at the cluster level.

RDS API

When you create a new DB instance using the [CreateDBInstance](../../../../reference/amazonrds/latest/apireference/api-createdbinstance.md) operation Amazon RDS API operation, turn on Performance Insights
by setting `EnablePerformanceInsights` to `True`. To turn off Performance Insights, set
`EnablePerformanceInsights` to `False` and set `DatabaseInsightsMode` to `standard`.

You can also specify the `EnablePerformanceInsights` value using
the following API operations:

- [CreateDBCluster](../../../../reference/amazonrds/latest/apireference/api-createdbcluster.md) (Multi-AZ DB cluster)

- [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) (Multi-AZ DB cluster)

- [ModifyDBInstance](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md)

- [CreateDBInstanceReadReplica](../../../../reference/amazonrds/latest/apireference/api-createdbinstancereadreplica.md)

- [RestoreDBInstanceFromS3](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_RestoreDBInstanceFromS3.html)

When you turn on Performance Insights, you can optionally specify the amount of time, in days, to retain Performance Insights data with the
`PerformanceInsightsRetentionPeriod` parameter. You can specify `7`, `month` \\* 31 (where `month` is a number from 1–23),
or `731`. For example, if you want to retain your performance data for 3 months, specify `93`, which is 3 \* 31. The default
is `7` days. For more information about retention periods, see [Pricing and data retention for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.Overview.cost.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Pricing and data retention for Performance Insights

Performance Schema for MariaDB or
MySQL
