# DBShardGroup

Contains the details for an Amazon RDS DB shard group.

## Contents

###### Note

In the following list, the required parameters are described first.

**ComputeRedundancy**

Specifies whether to create standby DB shard groups for the DB shard group. Valid values are the following:

- 0 - Creates a DB shard group without a standby DB shard group. This is the default value.

- 1 - Creates a DB shard group with a standby DB shard group in a different Availability Zone (AZ).

- 2 - Creates a DB shard group with two standby DB shard groups in two different AZs.

Type: Integer

Required: No

**DBClusterIdentifier**

The name of the primary DB cluster for the DB shard group.

Type: String

Required: No

**DBShardGroupArn**

The Amazon Resource Name (ARN) for the DB shard group.

Type: String

Required: No

**DBShardGroupIdentifier**

The name of the DB shard group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

**DBShardGroupResourceId**

The AWS Region-unique, immutable identifier for the DB shard group.

Type: String

Required: No

**Endpoint**

The connection endpoint for the DB shard group.

Type: String

Required: No

**MaxACU**

The maximum capacity of the DB shard group in Aurora capacity units (ACUs).

Type: Double

Required: No

**MinACU**

The minimum capacity of the DB shard group in Aurora capacity units (ACUs).

Type: Double

Required: No

**PubliclyAccessible**

Indicates whether the DB shard group is publicly accessible.

When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint
resolves to the private IP address from within the DB shard group's virtual private cloud
(VPC). It resolves to the public IP address from outside of the DB shard group's VPC. Access
to the DB shard group is ultimately controlled by the security group it uses. That public
access isn't permitted if the security group assigned to the DB shard group doesn't permit
it.

When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.

For more information, see [CreateDBShardGroup](api-createdbshardgroup.md).

This setting is only for Aurora Limitless Database.

Type: Boolean

Required: No

**Status**

The status of the DB shard group.

Type: String

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

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbshardgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbshardgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbshardgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBSecurityGroupMembership

DBSnapshot

All content copied from https://docs.aws.amazon.com/.
