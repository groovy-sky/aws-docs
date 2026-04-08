# ModifyDBProxy

Changes the settings for an existing DB proxy.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyName**

The identifier for the `DBProxy` to modify.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

**Auth.member.N**

The new authentication settings for the `DBProxy`.

Type: Array of [UserAuthConfig](api-userauthconfig.md) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

Required: No

**DebugLogging**

Specifies whether the proxy logs detailed connection and query information.
When you enable `DebugLogging`, the proxy captures connection details
and connection pool behavior from your queries. Debug logging increases CloudWatch costs
and can impact proxy performance. Enable this option only when you need
to troubleshoot connection or performance issues.

Type: Boolean

Required: No

**DefaultAuthScheme**

The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.
Valid values are `NONE` and `IAM_AUTH`.
When set to `IAM_AUTH`, the proxy uses end-to-end IAM authentication to connect to the database.

Type: String

Valid Values: `IAM_AUTH | NONE`

Required: No

**IdleClientTimeout**

The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this
value higher or lower than the connection timeout limit for the associated database.

Type: Integer

Required: No

**NewDBProxyName**

The new identifier for the `DBProxy`. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: No

**RequireTLS**

Whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
By enabling this setting, you can enforce encrypted TLS connections to the proxy, even if the associated database doesn't
use TLS.

Type: Boolean

Required: No

**RoleArn**

The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Required: No

**SecurityGroups.member.N**

The new list of security groups for the `DBProxy`.

Type: Array of strings

Required: No

## Response Elements

The following element is returned by the service.

**DBProxy**

The `DBProxy` object representing the new settings for the proxy.

Type: [DBProxy](api-dbproxy.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyAlreadyExistsFault**

The specified proxy name must be unique for all proxies owned by your AWS account in the specified AWS Region.

HTTP Status Code: 400

**DBProxyNotFoundFault**

The specified proxy name doesn't correspond to a proxy owned by your AWS account in the specified AWS Region.

HTTP Status Code: 404

**InvalidDBProxyStateFault**

The requested operation can't be performed while the proxy is in this state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbproxy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbproxy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBParameterGroup

ModifyDBProxyEndpoint

All content copied from https://docs.aws.amazon.com/.
