# ModifyDBProxyEndpoint

Changes the settings for an existing DB proxy endpoint.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyEndpointName**

The name of the DB proxy sociated with the DB proxy endpoint that you want to modify.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**NewDBProxyEndpointName**

The new identifier for the `DBProxyEndpoint`. An identifier must
begin with a letter and must contain only ASCII letters, digits, and hyphens; it
can't end with a hyphen or contain two consecutive hyphens.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

**VpcSecurityGroupIds.member.N**

The VPC security group IDs for the DB proxy endpoint. When the DB proxy endpoint
uses a different VPC than the original proxy, you also specify a different
set of security group IDs than for the original proxy.

Type: Array of strings

Required: No

## Response Elements

The following element is returned by the service.

**DBProxyEndpoint**

The `DBProxyEndpoint` object representing the new settings for the DB proxy endpoint.

Type: [DBProxyEndpoint](api-dbproxyendpoint.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyEndpointAlreadyExistsFault**

The specified DB proxy endpoint name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.

HTTP Status Code: 400

**DBProxyEndpointNotFoundFault**

The DB proxy endpoint doesn't exist.

HTTP Status Code: 404

**InvalidDBProxyEndpointStateFault**

You can't perform this operation while the DB proxy endpoint is in a particular state.

HTTP Status Code: 400

**InvalidDBProxyStateFault**

The requested operation can't be performed while the proxy is in this state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbproxyendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbproxyendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBProxy

ModifyDBProxyTargetGroup

All content copied from https://docs.aws.amazon.com/.
