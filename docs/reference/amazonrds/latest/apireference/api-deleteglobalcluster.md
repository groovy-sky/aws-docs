# DeleteGlobalCluster

Deletes a global database cluster. The primary and secondary clusters must already be detached or
destroyed first.

###### Note

This action only applies to Aurora DB clusters.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**GlobalClusterIdentifier**

The cluster identifier of the global database cluster being deleted.

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

**GlobalClusterNotFoundFault**

The `GlobalClusterIdentifier` doesn't refer to an existing global database cluster.

HTTP Status Code: 404

**InvalidGlobalClusterStateFault**

The global cluster is in an invalid state and can't perform the requested operation.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deleteglobalcluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deleteglobalcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteEventSubscription

DeleteIntegration

All content copied from https://docs.aws.amazon.com/.
