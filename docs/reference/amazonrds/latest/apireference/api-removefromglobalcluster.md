# RemoveFromGlobalCluster

Detaches an Aurora secondary cluster from an Aurora global database cluster. The cluster becomes a
standalone cluster with read-write capability instead of being read-only and receiving data from a
primary cluster in a different Region.

###### Note

This operation only applies to Aurora DB clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DbClusterIdentifier**

The Amazon Resource Name (ARN) identifying the cluster that was detached from the Aurora global database cluster.

Type: String

Required: Yes

**GlobalClusterIdentifier**

The cluster identifier to detach from the Aurora global database cluster.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/removefromglobalcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/removefromglobalcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegisterDBProxyTargets

RemoveRoleFromDBCluster

All content copied from https://docs.aws.amazon.com/.
