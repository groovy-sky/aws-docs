# Logging with AWS CloudTrail for S3 Tables

Amazon S3 is integrated with AWS CloudTrail, a service that provides a record of actions taken by a
user, role, or an AWS service. CloudTrail captures all API calls for Amazon S3 as events. Using
the information collected by CloudTrail, you can determine the request that was made to Amazon S3,
the IP address from which the request was made, when it was made, and additional details.
When a supported event activity occurs in Amazon S3, that activity is recorded in a CloudTrail
event. You can use AWS CloudTrail trail to log management events and data events for S3 Tables.
For more information, see [Amazon S3 CloudTrail events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md)
and [What is\
AWS CloudTrail?](../../../awscloudtrail/latest/userguide/cloudtrail-user-guide.md) in the _AWS CloudTrailUser Guide_.

## CloudTrail management events for S3 Tables

Management events provide information about management operations that are performed on resources in your AWS account.

By default, CloudTrail logs management events for S3 Tables. The `eventsource` for CloudTrail management events for S3 Tables is ` s3tables.amazonaws.com`.
When you set up your AWS account, CloudTrail management events are enabled by default. The following API actions are tracked by CloudTrail and logged as management events.

- [`CreateNamespace`](../api/api-s3tablebuckets-createnamespace.md)

- [`CreateTable`](../api/api-s3tablebuckets-createtable.md)

- [`CreateTableBucket`](../api/api-s3tablebuckets-createtablebucket.md)

- [`DeleteNamespace`](../api/api-s3tablebuckets-deletenamespace.md)

- [`DeleteTable`](../api/api-s3tablebuckets-deletetable.md)

- [`DeleteTableBucket`](../api/api-s3tablebuckets-deletetablebucket.md)

- [`DeleteTableBucketPolicy`](../api/api-s3tablebuckets-deletetablebucketpolicy.md)

- [`DeleteTablePolicy`](../api/api-s3tablebuckets-deletetablepolicy.md)

- [`GetNamespace`](../api/api-s3tablebuckets-getnamespace.md)

- [`GetTable`](../api/api-s3tablebuckets-gettable.md)

- [`GetTableBucket`](../api/api-s3tablebuckets-gettablebucket.md)

- [`GetTableBucketMaintenanceConfiguration`](../api/api-s3tablebuckets-gettablebucketmaintenanceconfiguration.md)

- [`GetTableBucketPolicy`](../api/api-s3tablebuckets-gettablebucketpolicy.md)

- [`GetTableMaintenanceConfiguration`](../api/api-s3tablebuckets-gettablemaintenanceconfiguration.md)

- [`GetTableMaintenanceJobStatus`](../api/api-s3tablebuckets-gettablemaintenancejobstatus.md)

- [`GetTableMetadataLocation`](../api/api-s3tablebuckets-gettablemetadatalocation.md)

- [`GetTablePolicy`](../api/api-s3tablebuckets-gettablepolicy.md)

- [`ListNamespaces`](../api/api-s3tablebuckets-listnamespaces.md)

- [`ListTableBuckets`](../api/api-s3tablebuckets-listtablebuckets.md)

- [`ListTables`](../api/api-s3tablebuckets-listtables.md)

- [`PutTableBucketMaintenanceConfiguration`](../api/api-s3tablebuckets-puttablebucketmaintenanceconfiguration.md)

- [`PutTableMaintenanceConfiguration`](../api/api-s3tablebuckets-puttablemaintenanceconfiguration.md)

- [`PutBucketPolicy`](../api/api-s3tablebuckets-putbucketpolicy.md)

- [`PutTablePolicy`](../api/api-s3tablebuckets-puttablepolicy.md)

- [`RenameTable`](../api/api-s3tablebuckets-renametable.md)

- [`UpdateTableMetadataLocation`](../api/api-s3tablebuckets-updatetablemetadatalocation.md)

For more information on CloudTrail management events, see [Logging management events](../../../awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.md) in the _AWS CloudTrail User_
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

For an example of a compaction maintenance event, see [Example – CloudTrail log file for a table maintenance management event](s3-tables-log-files.md#example-ct-log-s3tables-3).

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

For more information on CloudTrail data events, see [Logging data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md) in the _AWS CloudTrail User_
_Guide_.

For additional information about CloudTrail events for S3 Tables, see the following
topics:

###### Topics

- [AWS CloudTrail data event log file examples for S3 Tables](s3-tables-log-files.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging and monitoring

CloudTrail log examples

All content copied from https://docs.aws.amazon.com/.
