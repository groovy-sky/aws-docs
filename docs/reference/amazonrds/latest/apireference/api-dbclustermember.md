# DBClusterMember

Contains information about an instance that is part of a DB cluster.

## Contents

###### Note

In the following list, the required parameters are described first.

**DBClusterParameterGroupStatus**

Specifies the status of the DB cluster parameter group for this member of the DB cluster.

Type: String

Required: No

**DBInstanceIdentifier**

Specifies the instance identifier for this member of the DB cluster.

Type: String

Required: No

**IsClusterWriter**

Indicates whether the cluster member is the primary DB instance for the DB cluster.

Type: Boolean

Required: No

**PromotionTier**

A value that specifies the order in which an Aurora Replica is promoted to the primary instance
after a failure of the existing primary instance. For more information,
see [Fault Tolerance for an Aurora DB Cluster](../../../../services/amazonrds/latest/aurorauserguide/aurora-managing-backups.md#Aurora.Managing.FaultTolerance) in the _Amazon Aurora User Guide_.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbclustermember.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbclustermember.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbclustermember.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBClusterEndpoint

DBClusterOptionGroupStatus

All content copied from https://docs.aws.amazon.com/.
