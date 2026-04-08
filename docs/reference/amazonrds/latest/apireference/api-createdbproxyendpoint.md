# CreateDBProxyEndpoint

Creates a `DBProxyEndpoint`. Only applies to proxies that are associated with Aurora DB clusters.
You can use DB proxy endpoints to specify read/write or read-only access to the DB cluster. You can also use
DB proxy endpoints to access a DB proxy through a different VPC than the proxy's default VPC.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyEndpointName**

The name of the DB proxy endpoint to create.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**DBProxyName**

The name of the DB proxy associated with the DB proxy endpoint that you create.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**VpcSubnetIds.member.N**

The VPC subnet IDs for the DB proxy endpoint that you create. You can specify a
different set of subnet IDs than for the original DB proxy.

Type: Array of strings

Required: Yes

**EndpointNetworkType**

The network type of the DB proxy endpoint. The network type determines the IP version that the proxy endpoint supports.

Valid values:

- `IPV4` \- The proxy endpoint supports IPv4 only.

- `IPV6` \- The proxy endpoint supports IPv6 only.

- `DUAL` \- The proxy endpoint supports both IPv4 and IPv6.

Default: `IPV4`

Constraints:

- If you specify `IPV6` or `DUAL`, the VPC and all subnets must have an IPv6 CIDR block.

- If you specify `IPV6` or `DUAL`, the VPC tenancy cannot be `dedicated`.

Type: String

Valid Values: `IPV4 | IPV6 | DUAL`

Required: No

**Tags.Tag.N**

A list of tags.

For more information, see
[Tagging Amazon RDS resources](../../../../services/amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../../services/amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

Type: Array of [Tag](api-tag.md) objects

Required: No

**TargetRole**

The role of the DB proxy endpoint. The role determines whether the endpoint can be used for read/write
or only read operations. The default is `READ_WRITE`. The only role that proxies for RDS for Microsoft SQL Server
support is `READ_WRITE`.

Type: String

Valid Values: `READ_WRITE | READ_ONLY`

Required: No

**VpcSecurityGroupIds.member.N**

The VPC security group IDs for the DB proxy endpoint that you create. You can
specify a different set of security group IDs than for the original DB proxy.
The default is the default security group for the VPC.

Type: Array of strings

Required: No

## Response Elements

The following element is returned by the service.

**DBProxyEndpoint**

The `DBProxyEndpoint` object that is created by the API operation.
The DB proxy endpoint that you create might provide capabilities such as read/write
or read-only operations, or using a different VPC than the proxy's default VPC.

Type: [DBProxyEndpoint](api-dbproxyendpoint.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyEndpointAlreadyExistsFault**

The specified DB proxy endpoint name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.

HTTP Status Code: 400

**DBProxyEndpointQuotaExceededFault**

The DB proxy already has the maximum number of endpoints.

HTTP Status Code: 400

**DBProxyNotFoundFault**

The specified proxy name doesn't correspond to a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**InvalidDBProxyStateFault**

The requested operation can't be performed while the proxy is in this state.

HTTP Status Code: 400

**InvalidSubnet**

The requested subnet is invalid, or multiple subnets were requested that are not all in a common VPC.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createdbproxyendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createdbproxyendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDBProxy

CreateDBSecurityGroup

All content copied from https://docs.aws.amazon.com/.
