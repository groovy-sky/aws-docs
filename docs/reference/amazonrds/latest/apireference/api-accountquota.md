# AccountQuota

Describes a quota for an AWS account.

The following are account quotas:

- `AllocatedStorage` \- The total allocated storage per account, in GiB.
The used value is the total allocated storage in the account, in GiB.

- `AuthorizationsPerDBSecurityGroup` \- The number of ingress rules per DB security group.
The used value is the highest number of ingress rules in a DB security group in the account. Other
DB security groups in the account might have a lower number of ingress rules.

- `CustomEndpointsPerDBCluster` \- The number of custom endpoints per DB cluster.
The used value is the highest number of custom endpoints in a DB clusters in the account. Other
DB clusters in the account might have a lower number of custom endpoints.

- `DBClusterParameterGroups` \- The number of DB cluster parameter groups
per account, excluding default parameter groups. The used value is the count of
nondefault DB cluster parameter groups in the account.

- `DBClusterRoles` \- The number of associated AWS Identity and Access Management (IAM) roles per DB cluster.
The used value is the highest number of associated IAM roles for a DB cluster in the account. Other
DB clusters in the account might have a lower number of associated IAM roles.

- `DBClusters` \- The number of DB clusters per account.
The used value is the count of DB clusters in the account.

- `DBInstanceRoles` \- The number of associated IAM roles per DB instance.
The used value is the highest number of associated IAM roles for a DB instance in the account. Other
DB instances in the account might have a lower number of associated IAM roles.

- `DBInstances` \- The number of DB instances per account.
The used value is the count of the DB instances in the account.

Amazon RDS DB instances, Amazon Aurora DB instances, Amazon Neptune instances, and Amazon DocumentDB
instances apply to this quota.

- `DBParameterGroups` \- The number of DB parameter groups per account,
excluding default parameter groups. The used value is the count of nondefault DB
parameter groups in the account.

- `DBSecurityGroups` \- The number of DB security groups (not VPC
security groups) per account, excluding the default security group. The used
value is the count of nondefault DB security groups in the account.

- `DBSubnetGroups` \- The number of DB subnet groups per account.
The used value is the count of the DB subnet groups in the account.

- `EventSubscriptions` \- The number of event subscriptions per account.
The used value is the count of the event subscriptions in the account.

- `ManualClusterSnapshots` \- The number of manual DB cluster snapshots per account.
The used value is the count of the manual DB cluster snapshots in the account.

- `ManualSnapshots` \- The number of manual DB instance snapshots per account.
The used value is the count of the manual DB instance snapshots in the account.

- `OptionGroups` \- The number of DB option groups per account, excluding
default option groups. The used value is the count of nondefault DB option
groups in the account.

- `ReadReplicasPerMaster` \- The number of read replicas per DB
instance. The used value is the highest number of read replicas for a DB
instance in the account. Other DB instances in the account might have a lower
number of read replicas.

- `ReservedDBInstances` \- The number of reserved DB instances per account.
The used value is the count of the active reserved DB instances in the account.

- `SubnetsPerDBSubnetGroup` \- The number of subnets per DB subnet group.
The used value is highest number of subnets for a DB subnet group in the account. Other
DB subnet groups in the account might have a lower number of subnets.

For more information, see [Quotas for Amazon RDS](../../../../services/amazonrds/latest/userguide/chap-limits.md) in the
_Amazon RDS User Guide_ and [Quotas for Amazon Aurora](../../../../services/amazonrds/latest/aurorauserguide/chap-limits.md) in the
_Amazon Aurora User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**AccountQuotaName**

The name of the Amazon RDS quota for this AWS account.

Type: String

Required: No

**Max**

The maximum allowed value for the quota.

Type: Long

Required: No

**Used**

The amount currently used toward the quota maximum.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/accountquota.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/accountquota.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/accountquota.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

AdditionalStorageVolume

All content copied from https://docs.aws.amazon.com/.
