# StartDBInstanceAutomatedBackupsReplication

Enables replication of automated backups to a different AWS Region.

This command doesn't apply to RDS Custom.

For more information, see [Replicating Automated Backups to Another AWS Region](../../../../services/amazonrds/latest/userguide/user-replicatebackups.md) in the _Amazon RDS User Guide._

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceDBInstanceArn**

The Amazon Resource Name (ARN) of the source DB instance for the replicated automated backups, for example,
`arn:aws:rds:us-west-2:123456789012:db:mydatabase`.

Type: String

Required: Yes

**BackupRetentionPeriod**

The retention period for the replicated automated backups.

Type: Integer

Required: No

**KmsKeyId**

The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the
Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS Region, for example,
`arn:aws:kms:us-east-1:123456789012:key/AWS_ACCESS_KEY_ID_REDACTED`.

Type: String

Required: No

**PreSignedUrl**

In an AWS GovCloud (US) Region, an URL that contains a Signature Version 4 signed request
for the `StartDBInstanceAutomatedBackupsReplication` operation to call
in the AWS Region of the source DB instance. The presigned URL must be a valid request for the
`StartDBInstanceAutomatedBackupsReplication` API operation that can run in
the AWS Region that contains the source DB instance.

This setting applies only to AWS GovCloud (US) Regions. It's ignored in other
AWS Regions.

To learn how to generate a Signature Version 4 signed request, see
[Authenticating Requests: Using Query Parameters (AWS Signature Version 4)](../../../../services/s3/latest/api/sigv4-query-string-auth.md) and
[Signature Version 4 Signing Process](../../../../general/latest/gr/signature-version-4.md).

###### Note

If you are using an AWS SDK tool or the AWS CLI, you can specify
`SourceRegion` (or `--source-region` for the AWS CLI)
instead of specifying `PreSignedUrl` manually. Specifying
`SourceRegion` autogenerates a presigned URL that is a valid request
for the operation that can run in the source AWS Region.

Type: String

Required: No

**Tags.Tag.N**

A list of tags to associate with the replicated automated backups.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**DBInstanceAutomatedBackup**

An automated backup of a DB instance. It consists of system backups, transaction logs, and the database instance properties that
existed at the time you deleted the source instance.

Type: [DBInstanceAutomatedBackup](api-dbinstanceautomatedbackup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceAutomatedBackupQuotaExceeded**

The quota for retained automated backups was exceeded. This prevents you
from retaining any additional automated backups. The retained automated backups
quota is the same as your DB instance quota.

HTTP Status Code: 400

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**InvalidDBInstanceAutomatedBackupState**

The automated backup is in an invalid state.
For example, this automated backup is associated with an active instance.

HTTP Status Code: 400

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**KMSKeyNotAccessibleFault**

An error occurred accessing an AWS KMS key.

HTTP Status Code: 400

**StorageTypeNotSupported**

The specified `StorageType` can't be associated with the DB instance.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/startdbinstanceautomatedbackupsreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartDBInstance

StartExportTask

All content copied from https://docs.aws.amazon.com/.
