# ModifyDBProxyTargetGroup

Modifies the properties of a `DBProxyTargetGroup`.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyName**

The name of the proxy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**TargetGroupName**

The name of the target group to modify.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**ConnectionPoolConfig**

The settings that determine the size and behavior of the connection pool for the target group.

Type: [ConnectionPoolConfiguration](api-connectionpoolconfiguration.md) object

Required: No

**NewName**

The new name for the modified `DBProxyTarget`. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.

You can't rename the `default` target group.

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**DBProxyTargetGroup**

The settings of the modified `DBProxyTarget`.

Type: [DBProxyTargetGroup](api-dbproxytargetgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyNotFoundFault**

The specified proxy name doesn't correspond to a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**DBProxyTargetGroupNotFoundFault**

The specified target group isn't available for a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**InvalidDBProxyStateFault**

The requested operation can't be performed while the proxy is in this state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbproxytargetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbproxytargetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBProxyEndpoint

ModifyDBRecommendation

All content copied from https://docs.aws.amazon.com/.
