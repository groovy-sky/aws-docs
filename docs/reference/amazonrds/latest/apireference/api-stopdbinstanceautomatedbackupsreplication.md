# StopDBInstanceAutomatedBackupsReplication

Stops automated backup replication for a DB instance.

This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.

For more information, see [Replicating Automated Backups to Another AWS Region](../../../../services/amazonrds/latest/userguide/user-replicatebackups.md) in the _Amazon RDS User Guide._

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceDBInstanceArn**

The Amazon Resource Name (ARN) of the source DB instance for which to stop replicating
automate backups, for example,
`arn:aws:rds:us-west-2:123456789012:db:mydatabase`.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**DBInstanceAutomatedBackup**

An automated backup of a DB instance. It consists of system backups, transaction logs, and the database instance properties that
existed at the time you deleted the source instance.

Type: [DBInstanceAutomatedBackup](api-dbinstanceautomatedbackup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/stopdbinstanceautomatedbackupsreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopDBInstance

SwitchoverBlueGreenDeployment

All content copied from https://docs.aws.amazon.com/.
