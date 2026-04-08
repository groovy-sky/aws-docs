# DeleteDBClusterAutomatedBackup

Deletes automated backups using the `DbClusterResourceId` value of the source DB cluster or the Amazon
Resource Name (ARN) of the automated backups.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DbClusterResourceId**

The identifier for the source DB cluster, which can't be changed and which is unique to an AWS Region.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**DBClusterAutomatedBackup**

An automated backup of a DB cluster. It consists of system backups, transaction logs, and the database cluster
properties that existed at the time you deleted the source cluster.

Type: [DBClusterAutomatedBackup](api-dbclusterautomatedbackup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterAutomatedBackupNotFoundFault**

No automated backup for this DB cluster was found.

HTTP Status Code: 404

**InvalidDBClusterAutomatedBackupStateFault**

The automated backup is in an invalid state.
For example, this automated backup is associated with an active cluster.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deletedbclusterautomatedbackup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deletedbclusterautomatedbackup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBCluster

DeleteDBClusterEndpoint

All content copied from https://docs.aws.amazon.com/.
