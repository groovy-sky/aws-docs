# DeregisterDBProxyTargets

Remove the association between one or more `DBProxyTarget` data structures and a `DBProxyTargetGroup`.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyName**

The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**DBClusterIdentifiers.member.N**

One or more DB cluster identifiers.

Type: Array of strings

Required: No

**DBInstanceIdentifiers.member.N**

One or more DB instance identifiers.

Type: Array of strings

Required: No

**TargetGroupName**

The identifier of the `DBProxyTargetGroup`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyNotFoundFault**

The specified proxy name doesn't correspond to a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**DBProxyTargetGroupNotFoundFault**

The specified target group isn't available for a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**DBProxyTargetNotFoundFault**

The specified RDS DB instance or Aurora DB cluster isn't available for a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**InvalidDBProxyStateFault**

The requested operation can't be performed while the proxy is in this state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deregisterdbproxytargets.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deregisterdbproxytargets.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteTenantDatabase

DescribeAccountAttributes

All content copied from https://docs.aws.amazon.com/.
