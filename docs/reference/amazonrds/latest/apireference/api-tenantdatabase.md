# TenantDatabase

A tenant database in the DB instance. This data type is an element in the response to
the `DescribeTenantDatabases` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**CharacterSetName**

The character set of the tenant database.

Type: String

Required: No

**DBInstanceIdentifier**

The ID of the DB instance that contains the tenant database.

Type: String

Required: No

**DbiResourceId**

The AWS Region-unique, immutable identifier for the DB instance.

Type: String

Required: No

**DeletionProtection**

Specifies whether deletion protection is enabled for the DB instance.

Type: Boolean

Required: No

**MasterUsername**

The master username of the tenant database.

Type: String

Required: No

**MasterUserSecret**

Contains the secret managed by RDS in AWS Secrets Manager for the master user password.

For more information, see [Password management with AWS Secrets Manager](../../../../services/amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide_ and [Password management with AWS Secrets Manager](../../../../services/amazonrds/latest/aurorauserguide/rds-secrets-manager.md)
in the _Amazon Aurora User Guide._

Type: [MasterUserSecret](api-masterusersecret.md) object

Required: No

**NcharCharacterSetName**

The `NCHAR` character set name of the tenant database.

Type: String

Required: No

**PendingModifiedValues**

Information about pending changes for a tenant database.

Type: [TenantDatabasePendingModifiedValues](api-tenantdatabasependingmodifiedvalues.md) object

Required: No

**Status**

The status of the tenant database.

Type: String

Required: No

**TagList.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

**TenantDatabaseARN**

The Amazon Resource Name (ARN) for the tenant database.

Type: String

Required: No

**TenantDatabaseCreateTime**

The creation time of the tenant database.

Type: Timestamp

Required: No

**TenantDatabaseResourceId**

The AWS Region-unique, immutable identifier for the tenant database.

Type: String

Required: No

**TenantDBName**

The database name of the tenant database.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/tenantdatabase.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/tenantdatabase.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/tenantdatabase.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetHealth

TenantDatabasePendingModifiedValues

All content copied from https://docs.aws.amazon.com/.
