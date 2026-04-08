# CloudwatchLogsExportConfiguration

The configuration setting for the log types to be enabled for export to CloudWatch
Logs for a specific DB instance or DB cluster.

The `EnableLogTypes` and `DisableLogTypes` arrays determine which logs will be exported
(or not exported) to CloudWatch Logs. The values within these arrays depend on the DB
engine being used.

For more information about exporting CloudWatch Logs for Amazon RDS DB instances, see [Publishing Database Logs to Amazon CloudWatch Logs](../../../../services/amazonrds/latest/userguide/user-logaccess.md#USER_LogAccess.Procedural.UploadtoCloudWatch) in the _Amazon RDS User Guide_.

For more information about exporting CloudWatch Logs for Amazon Aurora DB clusters, see [Publishing Database Logs to Amazon CloudWatch Logs](../../../../services/amazonrds/latest/aurorauserguide/user-logaccess.md#USER_LogAccess.Procedural.UploadtoCloudWatch) in the _Amazon Aurora User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**DisableLogTypes.member.N**

The list of log types to disable.

The following values are valid for each DB engine:

- Aurora MySQL - `audit | error | general | slowquery`

- Aurora PostgreSQL - `postgresql`

- RDS for MySQL - `error | general | slowquery`

- RDS for PostgreSQL - `postgresql | upgrade`

Type: Array of strings

Required: No

**EnableLogTypes.member.N**

The list of log types to enable.

The following values are valid for each DB engine:

- Aurora MySQL - `audit | error | general | slowquery`

- Aurora PostgreSQL - `postgresql`

- RDS for MySQL - `error | general | slowquery`

- RDS for PostgreSQL - `postgresql | upgrade`

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/cloudwatchlogsexportconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/cloudwatchlogsexportconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/cloudwatchlogsexportconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CharacterSet

ClusterPendingModifiedValues

All content copied from https://docs.aws.amazon.com/.
