# GlobalCluster

A data type representing an Aurora global database.

## Contents

###### Note

In the following list, the required parameters are described first.

**DatabaseName**

The default database name within the new global database cluster.

Type: String

Required: No

**DeletionProtection**

The deletion protection setting for the new global database cluster.

Type: Boolean

Required: No

**Endpoint**

The writer endpoint for the new global database cluster. This endpoint always
points to the writer DB instance in the current primary cluster.

Type: String

Required: No

**Engine**

The Aurora database engine used by the global database cluster.

Type: String

Required: No

**EngineLifecycleSupport**

The lifecycle type for the global cluster.

For more information, see CreateGlobalCluster.

Type: String

Required: No

**EngineVersion**

Indicates the database engine version.

Type: String

Required: No

**FailoverState**

A data object containing all properties for the current state of an in-process or pending switchover or failover process for this global cluster (Aurora global database).
This object is empty unless the `SwitchoverGlobalCluster` or `FailoverGlobalCluster` operation was called on this global cluster.

Type: [FailoverState](api-failoverstate.md) object

Required: No

**GlobalClusterArn**

The Amazon Resource Name (ARN) for the global database cluster.

Type: String

Required: No

**GlobalClusterIdentifier**

Contains a user-supplied global database cluster identifier. This identifier is the unique key that
identifies a global database cluster.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][0-9A-Za-z-:._]*`

Required: No

**GlobalClusterMembers.GlobalClusterMember.N**

The list of primary and secondary clusters within the global database cluster.

Type: Array of [GlobalClusterMember](api-globalclustermember.md) objects

Required: No

**GlobalClusterResourceId**

The AWS
[partition](../../../../services/glossary/latest/reference/glos-chap-id-docs-gateway.md#partition)-unique, immutable identifier for the global database cluster. This identifier is found in
AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed.

Type: String

Required: No

**Status**

Specifies the current state of this global database cluster.

Type: String

Required: No

**StorageEncrypted**

The storage encryption setting for the global database cluster.

Type: Boolean

Required: No

**StorageEncryptionType**

The type of encryption used to protect data at rest in the global database cluster. Possible values:

- `none` \- The global database cluster is not encrypted.

- `sse-rds` \- The global database cluster is encrypted using an AWS owned KMS key.

- `sse-kms` \- The global database cluster is encrypted using a customer managed KMS key or AWS managed KMS key.

Type: String

Valid Values: `none | sse-kms | sse-rds`

Required: No

**TagList.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/globalcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/globalcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/globalcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

GlobalClusterMember

All content copied from https://docs.aws.amazon.com/.
