# Stopping a database activity stream

You can stop an activity stream using the console or AWS CLI.

If you delete your Amazon RDS
database instance, the activity stream is stopped and the underlying Amazon Kinesis stream is deleted automatically.

###### To turn off an activity stream

1. Open the Amazon RDS console at [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Databases**.

3. Choose a database
    that you want to stop the database activity stream for.

4. For **Actions**, choose **Stop activity stream**.
    The **Database Activity Stream** window
    appears.
1. Choose **Immediately**.

      When you choose **Immediately**, the RDS instance restarts right away. If you choose
       **During the next maintenance window**, the RDS instance doesn't restart right away. In
       this case, the database activity stream doesn't stop until the next maintenance window.

2. Choose **Continue**.

To stop database activity streams for your database, configure the DB instance using the AWS CLI command [stop-activity-stream](../../../cli/latest/reference/rds/stop-activity-stream.md). Identify the AWS Region for the
DB instance using the
`--region` parameter. The `--apply-immediately` parameter is optional.

For Linux, macOS, or Unix:

```nohighlight

aws rds --region MY_REGION \
    stop-activity-stream \
    --resource-arn MY_DB_ARN \
    --apply-immediately
```

For Windows:

```nohighlight

aws rds --region MY_REGION ^
    stop-activity-stream ^
    --resource-arn MY_DB_ARN ^
    --apply-immediately
```

To stop database activity streams for your database, configure the DB instance using the [StopActivityStream](../../../../reference/amazonrds/latest/apireference/api-stopactivitystream.md) operation. Identify the AWS Region for the DB instance using the `Region` parameter. The
`ApplyImmediately` parameter is optional.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Getting the activity stream status

Monitoring activity streams

All content copied from https://docs.aws.amazon.com/.
