# DynamoDB API operations supported by resource-based policies

This topic lists the API operations that are supported by resource-based policies.
However, for cross-account access, you can only use a certain set of DynamoDB APIs through
resource-based policies. You can't attach resource-based policies to resource types, such as
backups and imports. The IAM actions, which correspond with the APIs operating on these
resource types, are excluded from the supported IAM actions in resource-based policies.
Because table administrators configure internal table settings within the same account, APIs,
such as [UpdateTimeToLive](../../../../reference/amazondynamodb/latest/apireference/api-updatetimetolive.md)
and [DisableKinesisStreamingDestination](../../../../reference/amazondynamodb/latest/apireference/api-disablekinesisstreamingdestination.md), don't support cross-account access through
resource-based policies.

The DynamoDB data plane and control plane APIs that support cross-account access also support
table name overloading, which lets you specify the table ARN instead of the table name. You
can specify table ARN in the `TableName` parameter of these APIs. However, not all
of these APIs support cross-account access.

###### Topics

- [Data plane API operations](#rbac-data-plane-actions)

- [PartiQL API operations](#rbac-partiql-actions)

- [Control plane API operations](#rbac-control-plane-actions)

- [Version 2019.11.21 (Current) global tables API operations](#rbac-current-global-table-actions)

- [Version 2017.11.29 (Legacy) global tables API operations](#rbac-legacy-global-table-actions)

- [Tags API operations](#rbac-tags-actions)

- [Backup and Restore API operations](#rbac-backup-restore-actions)

- [Continuous Backup/Restore (PITR) API operations](#rbac-continuous-backup-restore-actions)

- [Contributor Insights API operations](#rbac-contributor-insights-actions)

- [Export API operations](#rbac-export-actions)

- [Import API operations](#rbac-import-actions)

- [Amazon Kinesis Data Streams API operations](#rbac-kinesis-actions)

- [Resource-based policy API operations](#rbac-rbp-actions)

- [Time-to-Live API operations](#rbac-ttl-actions)

- [Other API operations](#rbac-other-actions)

- [DynamoDB Streams API operations](#rbac-ds-actions)

## Data plane API operations

The following table lists the API-level support provided by [data plane](howitworks-api.md#HowItWorks.API.DataPlane) API operations for resource-based
policies and cross-account access.

Data Plane - Tables/indexes APIsResource-based policy supportCross-account support

[DeleteItem](../../../../reference/amazondynamodb/latest/apireference/api-deleteitem.md)

YesYes

[GetItem](../../../../reference/amazondynamodb/latest/apireference/api-getitem.md)

YesYes

[PutItem](../../../../reference/amazondynamodb/latest/apireference/api-putitem.md)

YesYes

[Query](../../../../reference/amazondynamodb/latest/apireference/api-query.md)

YesYes

[Scan](../../../../reference/amazondynamodb/latest/apireference/api-scan.md)

YesYes

[UpdateItem](../../../../reference/amazondynamodb/latest/apireference/api-updateitem.md)

YesYes

[TransactGetItems](../../../../reference/amazondynamodb/latest/apireference/api-transactgetitems.md)

YesYes

[TransactWriteItems](../../../../reference/amazondynamodb/latest/apireference/api-transactwriteitems.md)

YesYes

[BatchGetItem](../../../../reference/amazondynamodb/latest/apireference/api-batchgetitem.md)

YesYes

[BatchWriteItem](../../../../reference/amazondynamodb/latest/apireference/api-batchwriteitem.md)

YesYes

## PartiQL API operations

The following table lists the API-level support provided by [PartiQL](howitworks-api.md#HowItWorks.API.DataPlane.partiql) API operations for
resource-based policies and cross-account access.

PartiQL APIsResource-based policy supportCross-account support

[BatchExecuteStatement](../../../../reference/amazondynamodb/latest/apireference/api-batchexecutestatement.md)

YesNo

[ExecuteStatement](../../../../reference/amazondynamodb/latest/apireference/api-executestatement.md)

YesNo

[ExecuteTransaction](../../../../reference/amazondynamodb/latest/apireference/api-executetransaction.md)

YesNo

## Control plane API operations

The following table lists the API-level support provided by [control plane](howitworks-api.md#HowItWorks.API.ControlPlane) API operations for
resource-based policies and cross-account access.

Control Plane - Tables APIsResource-based policy supportCross-account support

[CreateTable](../../../../reference/amazondynamodb/latest/apireference/api-createtable.md)

NoNo

[DeleteTable](../../../../reference/amazondynamodb/latest/apireference/api-deletetable.md)

YesYes

[DescribeTable](../../../../reference/amazondynamodb/latest/apireference/api-describetable.md)

YesYes

[UpdateTable](../../../../reference/amazondynamodb/latest/apireference/api-updatetable.md)

YesYes

## Version 2019.11.21 (Current) global tables API operations

The following table lists the API-level support provided by [Version 2019.11.21 (Current) global tables](globaltables.md) API operations for resource-based
policies and cross-account access.

Version 2019.11.21 (Current) global tables APIsResource-based policy supportCross-account support

[DescribeTableReplicaAutoScaling](../../../../reference/amazondynamodb/latest/apireference/api-describetablereplicaautoscaling.md)

YesNo

[UpdateTableReplicaAutoScaling](../../../../reference/amazondynamodb/latest/apireference/api-updatetablereplicaautoscaling.md)

YesNo

## Version 2017.11.29 (Legacy) global tables API operations

The following table lists the API-level support provided by [Version 2017.11.29 (Legacy) global tables](globaltables-v1.md) API operations
for resource-based policies and cross-account access.

Version 2017.11.29 (Legacy) global tables APIsResource-based policy supportCross-account support

[CreateGlobalTable](../../../../reference/amazondynamodb/latest/apireference/api-createglobaltable.md)

NoNo

[DescribeGlobalTable](../../../../reference/amazondynamodb/latest/apireference/api-describeglobaltable.md)

NoNo

[DescribeGlobalTableSettings](../../../../reference/amazondynamodb/latest/apireference/api-describeglobaltablesettings.md)

NoNo

[ListGlobalTables](../../../../reference/amazondynamodb/latest/apireference/api-listglobaltables.md)

NoNo

[UpdateGlobalTable](../../../../reference/amazondynamodb/latest/apireference/api-updateglobaltable.md)

NoNo

[UpdateGlobalTableSettings](../../../../reference/amazondynamodb/latest/apireference/api-updateglobaltablesettings.md)

NoNo

## Tags API operations

The following table lists the API-level support provided by API operations related to
[tags](tagging-operations.md) for resource-based policies and
cross-account access.

Tags APIsResource-based policy supportCross-account support

[ListTagsOfResource](../../../../reference/amazondynamodb/latest/apireference/api-listtagsofresource.md)

YesYes

[TagResource](../../../../reference/amazondynamodb/latest/apireference/api-tagresource.md)

YesYes

[UntagResource](../../../../reference/amazondynamodb/latest/apireference/api-untagresource.md)

YesYes

## Backup and Restore API operations

The following table lists the API-level support provided by API operations related to
[backup and restore](backup-and-restore.md) for resource-based policies
and cross-account access.

Backup and Restore APIsResource-based policy supportCross-account support

[CreateBackup](../../../../reference/amazondynamodb/latest/apireference/api-createbackup.md)

YesNo

[DescribeBackup](../../../../reference/amazondynamodb/latest/apireference/api-describebackup.md)

NoNo

[DeleteBackup](../../../../reference/amazondynamodb/latest/apireference/api-deletebackup.md)

NoNo

[RestoreTableFromBackup](../../../../reference/amazondynamodb/latest/apireference/api-restoretablefrombackup.md)

NoNo

## Continuous Backup/Restore (PITR) API operations

The following table lists the API-level support provided by API operations related to
[Continuous Backup/Restore (PITR)](point-in-time-recovery.md) for
resource-based policies and cross-account access.

Continuous Backup/Restore (PITR) APIsResource-based policy supportCross-account support

[DescribeContinuousBackups](../../../../reference/amazondynamodb/latest/apireference/api-describecontinuousbackups.md)

YesNo

[RestoreTableToPointInTime](../../../../reference/amazondynamodb/latest/apireference/api-restoretabletopointintime.md)

YesNo

[UpdateContinuousBackups](../../../../reference/amazondynamodb/latest/apireference/api-updatecontinuousbackups.md)

YesNo

## Contributor Insights API operations

The following table lists the API-level support provided by API operations related to
[Continuous Backup/Restore (PITR)](point-in-time-recovery.md) for
resource-based policies and cross-account access.

Contributor Insights APIsResource-based policy supportCross-account support

[DescribeContributorInsights](../../../../reference/amazondynamodb/latest/apireference/api-describecontributorinsights.md)

YesNo

[ListContributorInsights](../../../../reference/amazondynamodb/latest/apireference/api-listcontributorinsights.md)

NoNo

[UpdateContributorInsights](../../../../reference/amazondynamodb/latest/apireference/api-updatecontributorinsights.md)

YesNo

## Export API operations

The following table lists the API-level support provided by Export API operations for
resource-based policies and cross-account access.

Export APIsResource-based policy supportCross-account support

[DescribeExport](../../../../reference/amazondynamodb/latest/apireference/api-describeexport.md)

NoNo

[ExportTableToPointInTime](../../../../reference/amazondynamodb/latest/apireference/api-exporttabletopointintime.md)

YesNo

[ListExports](../../../../reference/amazondynamodb/latest/apireference/api-listexports.md)

NoNo

## Import API operations

The following table lists the API-level support provided by Import API operations for
resource-based policies and cross-account access.

Import APIsResource-based policy supportCross-account support

[DescribeImport](../../../../reference/amazondynamodb/latest/apireference/api-describeimport.md)

NoNo

[ImportTable](../../../../reference/amazondynamodb/latest/apireference/api-importtable.md)

NoNo

[ListImports](../../../../reference/amazondynamodb/latest/apireference/api-listimports.md)

NoNo

## Amazon Kinesis Data Streams API operations

The following table lists the API-level support provided by Kinesis Data Streams API operations for
resource-based policies and cross-account access.

Kinesis APIsResource-based policy supportCross-account support

[DescribeKinesisStreamingDestination](../../../../reference/amazondynamodb/latest/apireference/api-describekinesisstreamingdestination.md)

YesNo

[DisableKinesisStreamingDestination](../../../../reference/amazondynamodb/latest/apireference/api-disablekinesisstreamingdestination.md)

YesNo

[EnableKinesisStreamingDestination](../../../../reference/amazondynamodb/latest/apireference/api-enablekinesisstreamingdestination.md)

YesNo

[UpdateKinesisStreamingDestination](../../../../reference/amazondynamodb/latest/apireference/api-updatekinesisstreamingdestination.md)

YesNo

## Resource-based policy API operations

The following table lists the API-level support provided by resource-based policy API
operations for resource-based policies and cross-account access.

Resource-based policy APIsResource-based policy supportCross-account support

[GetResourcePolicy](../../../../reference/amazondynamodb/latest/apireference/api-getresourcepolicy.md)

YesNo

[PutResourcePolicy](../../../../reference/amazondynamodb/latest/apireference/api-putresourcepolicy.md)

YesNo

[DeleteResourcePolicy](../../../../reference/amazondynamodb/latest/apireference/api-deleteresourcepolicy.md)

YesNo

## Time-to-Live API operations

The following table lists the API-level support provided by [time to\
live](ttl.md) (TTL) API operations for resource-based policies and cross-account
access.

TTL APIsResource-based policy supportCross-account support

[DescribeTimeToLive](../../../../reference/amazondynamodb/latest/apireference/api-describetimetolive.md)

YesNo

[UpdateTimeToLive](../../../../reference/amazondynamodb/latest/apireference/api-updatetimetolive.md)

YesNo

## Other API operations

The following table lists the API-level support provided by other miscellaneous API
operations for resource-based policies and cross-account access.

Other APIsResource-based policy supportCross-account support

[DescribeLimits](../../../../reference/amazondynamodb/latest/apireference/api-describelimits.md)

NoNo

[DescribeEndpoints](../../../../reference/amazondynamodb/latest/apireference/api-describeendpoints.md)

NoNo

[ListBackups](../../../../reference/amazondynamodb/latest/apireference/api-listbackups.md)

NoNo

[ListTables](../../../../reference/amazondynamodb/latest/apireference/api-listtables.md)

NoNo

## DynamoDB Streams API operations

The following table lists the API-level support of DynamoDB Streams APIs for resource-based policies
and cross-account access.

DynamoDB Streams APIsResource-based policy supportCross-account support

[DescribeStream](../../../../reference/amazondynamodb/latest/apireference/api-streams-describestream.md)

YesYes

[GetRecords](../../../../reference/amazondynamodb/latest/apireference/api-streams-getrecords.md)

YesYes

[GetShardIterator](../../../../reference/amazondynamodb/latest/apireference/api-streams-getsharditerator.md)

YesYes

[ListStreams](../../../../reference/amazondynamodb/latest/apireference/api-streams-liststreams.md)

NoNo

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Blocking public access

IAM
authorization

All content copied from https://docs.aws.amazon.com/.
