# SwitchoverGlobalCluster

Switches over the specified secondary DB cluster to be the new primary DB cluster in the global database cluster.
Switchover operations were previously called "managed planned failovers."

Aurora promotes the specified secondary cluster to assume full read/write capabilities and demotes the current primary cluster
to a secondary (read-only) cluster, maintaining the orginal replication topology. All secondary clusters are synchronized with the primary
at the beginning of the process so the new primary continues operations for the Aurora global database without losing any data. Your database
is unavailable for a short time while the primary and selected secondary clusters are assuming their new roles. For more information about
switching over an Aurora global database, see [Performing switchovers for Amazon Aurora global databases](../../../../services/amazonrds/latest/aurorauserguide/aurora-global-database-disaster-recovery.md#aurora-global-database-disaster-recovery.managed-failover) in the _Amazon Aurora User Guide_.

###### Note

This operation is intended for controlled environments, for operations such as "regional rotation" or to fall back to the original
primary after a global database failover.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**GlobalClusterIdentifier**

The identifier of the global database cluster to switch over. This parameter isn't case-sensitive.

Constraints:

- Must match the identifier of an existing global database cluster (Aurora global database).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][0-9A-Za-z-:._]*`

Required: Yes

**TargetDbClusterIdentifier**

The identifier of the secondary Aurora DB cluster to promote to the new primary for the global database cluster. Use the Amazon Resource Name (ARN) for the identifier so that
Aurora can locate the cluster in its AWS Region.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][0-9A-Za-z-:._]*`

Required: Yes

## Response Elements

The following element is returned by the service.

**GlobalCluster**

A data type representing an Aurora global database.

Type: [GlobalCluster](api-globalcluster.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBClusterNotFoundFault**

`DBClusterIdentifier` doesn't refer to an existing DB cluster.

HTTP Status Code: 404

**GlobalClusterNotFoundFault**

The `GlobalClusterIdentifier` doesn't refer to an existing global database cluster.

HTTP Status Code: 404

**InvalidDBClusterStateFault**

The requested operation can't be performed while the cluster is in this state.

HTTP Status Code: 400

**InvalidGlobalClusterStateFault**

The global cluster is in an invalid state and can't perform the requested operation.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/switchoverglobalcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/switchoverglobalcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SwitchoverBlueGreenDeployment

SwitchoverReadReplica

All content copied from https://docs.aws.amazon.com/.
