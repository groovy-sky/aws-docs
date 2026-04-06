# Logging with AWS CloudTrail for S3 Tables

Amazon S3 is integrated with AWS CloudTrail, a service that provides a record of actions taken by a
user, role, or an AWS service. CloudTrail captures all API calls for Amazon S3 as events. Using
the information collected by CloudTrail, you can determine the request that was made to Amazon S3,
the IP address from which the request was made, when it was made, and additional details.
When a supported event activity occurs in Amazon S3, that activity is recorded in a CloudTrail
event. You can use AWS CloudTrail trail to log management events and data events for S3 Tables.
For more information, see [Amazon S3 CloudTrail events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html)
and [What is\
AWS CloudTrail?](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html) in the _AWS CloudTrailUser Guide_.

## CloudTrail management events for S3 Tables

Management events provide information about management operations that are performed on resources in your AWS account.

By default, CloudTrail logs management events for S3 Tables. The `eventsource` for CloudTrail management events for S3 Tables is ` s3tables.amazonaws.com`.
When you set up your AWS account, CloudTrail management events are enabled by default. The following API actions are tracked by CloudTrail and logged as management events.

- [`CreateNamespace`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_CreateNamespace.html)

- [`CreateTable`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_CreateTable.html)

- [`CreateTableBucket`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_CreateTableBucket.html)

- [`DeleteNamespace`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_DeleteNamespace.html)

- [`DeleteTable`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_DeleteTable.html)

- [`DeleteTableBucket`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_DeleteTableBucket.html)

- [`DeleteTableBucketPolicy`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_DeleteTableBucketPolicy.html)

- [`DeleteTablePolicy`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_DeleteTablePolicy.html)

- [`GetNamespace`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetNamespace.html)

- [`GetTable`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTable.html)

- [`GetTableBucket`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTableBucket.html)

- [`GetTableBucketMaintenanceConfiguration`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTableBucketMaintenanceConfiguration.html)

- [`GetTableBucketPolicy`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTableBucketPolicy.html)

- [`GetTableMaintenanceConfiguration`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTableMaintenanceConfiguration.html)

- [`GetTableMaintenanceJobStatus`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTableMaintenanceJobStatus.html)

- [`GetTableMetadataLocation`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTableMetadataLocation.html)

- [`GetTablePolicy`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_GetTablePolicy.html)

- [`ListNamespaces`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_ListNamespaces.html)

- [`ListTableBuckets`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_ListTableBuckets.html)

- [`ListTables`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_ListTables.html)

- [`PutTableBucketMaintenanceConfiguration`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_PutTableBucketMaintenanceConfiguration.html)

- [`PutTableMaintenanceConfiguration`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_PutTableMaintenanceConfiguration.html)

- [`PutBucketPolicy`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_PutBucketPolicy.html)

- [`PutTablePolicy`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_PutTablePolicy.html)

- [`RenameTable`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_RenameTable.html)

- [`UpdateTableMetadataLocation`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3TableBuckets_UpdateTableMetadataLocation.html)

For more information on CloudTrail management events, see [Logging management events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html) in the _AWS CloudTrail User_
_Guide_.

## CloudTrail management events for S3 Tables maintenance

S3 logs automatic maintenance operations as `TablesMaintenanceEvent` management events in CloudTrail. These events occur during operations like compaction and snapshot expiration. For more information about S3 table maintenance, see [Maintenance for tables](s3-tables-maintenance.md).

### How to identify maintenance events

You can identify S3 Tables maintenance events in CloudTrail logs by these attribute values:

- `eventSource: s3tables.amazonaws.com`

- `eventType: AwsServiceEvent`

- `eventName: TablesMaintenanceEvent`

- `userAgent: maintenance.s3tables.amazonaws.com`

- `activityType:`

- `IcebergCompaction` (for compaction)

- `IcebergSnapshotManagement` (for snapshot expiration)

For an example of a compaction maintenance event, see [Example – CloudTrail log file for a table maintenance management event](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-log-files.html#example-ct-log-s3tables-3).

## CloudTrail data events for S3 Tables

Data events provide information about the resource operations performed on or in a
resource.By default, CloudTrail trails don't log data events, but you can configure trails to log data events.

When you log data events for a trail in CloudTrail, you will choose or specify the
resource type. S3 Tables has two resources types, `AWS::S3Tables::Table` and
`AWS::S3Tables::TableBucket`.

The following data events are logged to CloudTrail.

- [`AbortMultipartUpload`](../api/api-abortmultipartupload.md)

- [`CompleteMultipartUpload`](../api/api-completemultipartupload.md)

- [`CreateMultipartUpload`](../api/api-createmultipartupload.md)

- [`GetObject`](../api/api-getobject.md)

- [`HeadObject`](../api/api-headobject.md)

- [`ListParts`](../api/api-listparts.md)

- [`PutObject`](../api/api-putobject.md)

- [`UploadPart`](../api/api-uploadpart.md)

For more information on CloudTrail data events, see [Logging data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) in the _AWS CloudTrail User_
_Guide_.

For additional information about CloudTrail events for S3 Tables, see the following
topics:

###### Topics

- [AWS CloudTrail data event log file examples for S3 Tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-log-files.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Logging and monitoring

CloudTrail log examples
