# DBSnapshotTenantDatabase

Contains the details of a tenant database in a snapshot of a DB instance.

## Contents

###### Note

In the following list, the required parameters are described first.

**CharacterSetName**

The name of the character set of a tenant database.

Type: String

Required: No

**DBInstanceIdentifier**

The ID for the DB instance that contains the tenant databases.

Type: String

Required: No

**DbiResourceId**

The resource identifier of the source CDB instance. This identifier can't be changed
and is unique to an AWS Region.

Type: String

Required: No

**DBSnapshotIdentifier**

The identifier for the snapshot of the DB instance.

Type: String

Required: No

**DBSnapshotTenantDatabaseARN**

The Amazon Resource Name (ARN) for the snapshot tenant database.

Type: String

Required: No

**EngineName**

The name of the database engine.

Type: String

Required: No

**MasterUsername**

The master username of the tenant database.

Type: String

Required: No

**NcharCharacterSetName**

The `NCHAR` character set name of the tenant database.

Type: String

Required: No

**SnapshotType**

The type of DB snapshot.

Type: String

Required: No

**TagList.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

**TenantDatabaseCreateTime**

The time the DB snapshot was taken, specified in Coordinated Universal Time (UTC). If
you copy the snapshot, the creation time changes.

Type: Timestamp

Required: No

**TenantDatabaseResourceId**

The resource ID of the tenant database.

Type: String

Required: No

**TenantDBName**

The name of the tenant database.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbsnapshottenantdatabase.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbsnapshottenantdatabase.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbsnapshottenantdatabase.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBSnapshotAttributesResult

DBSubnetGroup

All content copied from https://docs.aws.amazon.com/.
