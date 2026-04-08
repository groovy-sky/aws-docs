# DBProxy

The data structure representing a proxy managed by the RDS Proxy.

This data type is used as a response element in the `DescribeDBProxies` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**Auth.member.N**

One or more data structures specifying the authorization mechanism to connect to the associated RDS DB instance
or Aurora DB cluster.

Type: Array of [UserAuthConfigInfo](api-userauthconfiginfo.md) objects

Required: No

**CreatedDate**

The date and time when the proxy was first created.

Type: Timestamp

Required: No

**DBProxyArn**

The Amazon Resource Name (ARN) for the proxy.

Type: String

Required: No

**DBProxyName**

The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.

Type: String

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

Required: No

**Endpoint**

The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the
connection string for a database client application.

Type: String

Required: No

**EndpointNetworkType**

The network type of the DB proxy endpoint. The network type determines the IP version that the proxy endpoint supports.

Valid values:

- `IPV4` \- The proxy endpoint supports IPv4 only.

- `IPV6` \- The proxy endpoint supports IPv6 only.

- `DUAL` \- The proxy endpoint supports both IPv4 and IPv6.

Type: String

Valid Values: `IPV4 | IPV6 | DUAL`

Required: No

**EngineFamily**

The kinds of databases that the proxy can connect to. This value determines which database network protocol
the proxy recognizes when it interprets network traffic to and from the database. `MYSQL` supports Aurora MySQL,
RDS for MariaDB, and RDS for MySQL databases. `POSTGRESQL` supports Aurora PostgreSQL and RDS for PostgreSQL databases.
`SQLSERVER` supports RDS for Microsoft SQL Server databases.

Type: String

Required: No

**IdleClientTimeout**

The number of seconds a connection to the proxy can have no activity before the proxy drops the client connection.
The proxy keeps the underlying database connection open and puts it back into the connection pool for reuse by
later connection requests.

Default: 1800 (30 minutes)

Constraints: 1 to 28,800

Type: Integer

Required: No

**RequireTLS**

Indicates whether Transport Layer Security (TLS) encryption is required for connections to the proxy.

Type: Boolean

Required: No

**RoleArn**

The Amazon Resource Name (ARN) for the IAM role that the proxy uses to access Amazon Secrets Manager.

Type: String

Required: No

**Status**

The current status of this proxy. A status of `available` means the
proxy is ready to handle requests. Other values indicate that you must wait for
the proxy to be ready, or take some action to resolve an issue.

Type: String

Valid Values: `available | modifying | incompatible-network | insufficient-resource-limits | creating | deleting | suspended | suspending | reactivating`

Required: No

**TargetConnectionNetworkType**

The network type that the proxy uses to connect to the target database. The network type determines the IP version that the proxy uses for connections to the database.

Valid values:

- `IPV4` \- The proxy connects to the database using IPv4 only.

- `IPV6` \- The proxy connects to the database using IPv6 only.

Type: String

Valid Values: `IPV4 | IPV6`

Required: No

**UpdatedDate**

The date and time when the proxy was last updated.

Type: Timestamp

Required: No

**VpcId**

Provides the VPC ID of the DB proxy.

Type: String

Required: No

**VpcSecurityGroupIds.member.N**

Provides a list of VPC security groups that the proxy belongs to.

Type: Array of strings

Required: No

**VpcSubnetIds.member.N**

The EC2 subnet IDs for the proxy.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbproxy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbproxy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbproxy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBParameterGroupStatus

DBProxyEndpoint

All content copied from https://docs.aws.amazon.com/.
