This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBProxy

The `AWS::RDS::DBProxy` resource creates or updates a DB proxy.

For information about RDS Proxy for Amazon RDS, see [Managing Connections with Amazon\
RDS Proxy](../../../amazonrds/latest/userguide/rds-proxy.md) in the _Amazon RDS User Guide_.

For information about RDS Proxy for Amazon Aurora, see [Managing Connections with\
Amazon RDS Proxy](../../../amazonrds/latest/aurorauserguide/rds-proxy.md) in the _Amazon Aurora User Guide_.

###### Note

Limitations apply to RDS Proxy, including DB engine version limitations and AWS
Region limitations.

For information about limitations that apply to RDS Proxy for Amazon RDS, see
[Limitations for RDS Proxy](../../../amazonrds/latest/userguide/rds-proxy.md#rds-proxy.limitations) in the _Amazon RDS User Guide_.

For information about that apply to RDS Proxy for Amazon Aurora, see [Limitations for RDS Proxy](../../../amazonrds/latest/aurorauserguide/rds-proxy.md#rds-proxy.limitations) in the _Amazon Aurora User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBProxy",
  "Properties" : {
      "Auth" : [ AuthFormat, ... ],
      "DBProxyName" : String,
      "DebugLogging" : Boolean,
      "DefaultAuthScheme" : String,
      "EndpointNetworkType" : String,
      "EngineFamily" : String,
      "IdleClientTimeout" : Integer,
      "RequireTLS" : Boolean,
      "RoleArn" : String,
      "Tags" : [ TagFormat, ... ],
      "TargetConnectionNetworkType" : String,
      "VpcSecurityGroupIds" : [ String, ... ],
      "VpcSubnetIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBProxy
Properties:
  Auth:
    - AuthFormat
  DBProxyName: String
  DebugLogging: Boolean
  DefaultAuthScheme: String
  EndpointNetworkType: String
  EngineFamily: String
  IdleClientTimeout: Integer
  RequireTLS: Boolean
  RoleArn: String
  Tags:
    - TagFormat
  TargetConnectionNetworkType: String
  VpcSecurityGroupIds:
    - String
  VpcSubnetIds:
    - String

```

## Properties

`Auth`

The authorization mechanism that the proxy uses.

_Required_: No

_Type_: Array of [AuthFormat](aws-properties-rds-dbproxy-authformat.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBProxyName`

The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.

_Required_: Yes

_Type_: String

_Pattern_: `[0-z]*`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DebugLogging`

Specifies whether the proxy logs detailed connection and query information.
When you enable `DebugLogging`, the proxy captures connection details
and connection pool behavior from your queries. Debug logging increases CloudWatch costs
and can impact proxy performance. Enable this option only when you need
to troubleshoot connection or performance issues.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultAuthScheme`

The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.
Valid values are `NONE` and `IAM_AUTH`.
When set to `IAM_AUTH`, the proxy uses end-to-end IAM authentication to connect to the database.

_Required_: No

_Type_: String

_Allowed values_: `IAM_AUTH | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointNetworkType`

The network type of the DB proxy endpoint. The network type determines the IP version that the proxy endpoint supports.

Valid values:

- `IPV4` \- The proxy endpoint supports IPv4 only.

- `IPV6` \- The proxy endpoint supports IPv6 only.

- `DUAL` \- The proxy endpoint supports both IPv4 and IPv6.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | IPV6 | DUAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineFamily`

The kinds of databases that the proxy can connect to.
This value determines which database network protocol the proxy recognizes when it interprets
network traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for MySQL databases, specify `MYSQL`.
For Aurora PostgreSQL and RDS for PostgreSQL databases, specify `POSTGRESQL`. For RDS for Microsoft SQL Server, specify
`SQLSERVER`.

_Required_: Yes

_Type_: String

_Allowed values_: `MYSQL | POSTGRESQL | SQLSERVER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdleClientTimeout`

The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. You can set this
value higher or lower than the connection timeout limit for the associated database.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireTLS`

Specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
By enabling this setting, you can enforce encrypted TLS connections to the proxy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.

_Required_: No

_Type_: Array of [TagFormat](aws-properties-rds-dbproxy-tagformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetConnectionNetworkType`

The network type that the proxy uses to connect to the target database. The network type determines the IP version that the proxy uses for connections to the database.

Valid values:

- `IPV4` \- The proxy connects to the database using IPv4 only.

- `IPV6` \- The proxy connects to the database using IPv6 only.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | IPV6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSecurityGroupIds`

One or more VPC security group IDs to associate with the new proxy.

If you plan to update the resource, don't specify VPC security groups in a shared VPC.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSubnetIds`

One or more VPC subnet IDs to associate with the new proxy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the DB proxy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DBProxyArn`

The Amazon Resource Name (ARN) for the proxy.

`Endpoint`

The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the
connection string for a database client application.

`VpcId`

The VPC ID to associate with the DB proxy.

## Examples

### Creating a DB proxy and registering a DB instance

The following example creates a DB proxy and registers a DB instance.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "ProxyName": {
            "Type": "String",
            "Default": "exampleProxy"
        },
        "InstanceName": {
            "Type": "String",
            "Default": "database-1"
        },
        "BootstrapSecretReaderRoleArn": {
            "Type": "String",
            "Default": "arn:aws:iam::123456789012:role/RDSSecretReaderRoleForDBProxy"
        },
        "BootstrapProxySecretArn": {
            "Type": "String",
            "Default": "arn:aws:secretsmanager:us-west-2:123456789012:secret:MySecretForDBProxy"
        },
        "SubnetIds": {
            "Type": "String",
            "Default": "subnet-01b761b31fb498f20,subnet-012b9a958ef0f9949,subnet-05e8e263052025378"
        }
    },
    "Resources": {
        "TestDBProxy": {
            "Type": "AWS::RDS::DBProxy",
            "Properties": {
                "DebugLogging": true,
                "DBProxyName": {
                    "Ref": "ProxyName"
                },
                "EngineFamily": "MYSQL",
                "IdleClientTimeout": 120,
                "RequireTLS": true,
                "RoleArn": {
                    "Ref": "BootstrapSecretReaderRoleArn"
                },
                "Auth": [
                    {
                        "AuthScheme": "SECRETS",
                        "SecretArn": {
                            "Ref": "BootstrapProxySecretArn"
                        },
                        "IAMAuth": "DISABLED"
                    }
                ],
                "VpcSubnetIds": {
                    "Fn::Split": [
                        ",",
                        {
                            "Ref": "SubnetIds"
                        }
                    ]
                }
            }
        },
        "ProxyTargetGroup": {
            "Type": "AWS::RDS::DBProxyTargetGroup",
            "Properties": {
                "DBProxyName": {
                    "Ref": "TestDBProxy"
                },
                "DBInstanceIdentifiers": [
                    {
                        "Ref": "InstanceName"
                    }
                ],
                "TargetGroupName": "default",
                "ConnectionPoolConfigurationInfo": {
                    "MaxConnectionsPercent": 100,
                    "MaxIdleConnectionsPercent": 50,
                    "ConnectionBorrowTimeout": 120
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  ProxyName:
    Type: String
    Default: exampleProxy
  InstanceName:
    Type: String
    Default: database-1
  BootstrapSecretReaderRoleArn:
    Type: String
    Default: arn:aws:iam::123456789012:role/RDSSecretReaderRoleForDBProxy
  BootstrapProxySecretArn:
    Type: String
    Default: arn:aws:secretsmanager:us-west-2:123456789012:secret:MySecretForDBProxy
  SubnetIds:
    Type: String
    Default: subnet-01b761b31fb498f20,subnet-012b9a958ef0f9949,subnet-05e8e263052025378
Resources:
  TestDBProxy:
    Type: AWS::RDS::DBProxy
    Properties:
      DebugLogging: true
      DBProxyName: !Ref ProxyName
      EngineFamily: MYSQL
      IdleClientTimeout: 120
      RequireTLS: true
      RoleArn:
        !Ref BootstrapSecretReaderRoleArn
      Auth:
        - {AuthScheme: SECRETS, SecretArn: !Ref BootstrapProxySecretArn, IAMAuth: DISABLED}
      VpcSubnetIds:
        Fn::Split: [",", !Ref SubnetIds]
  ProxyTargetGroup:
    Type: AWS::RDS::DBProxyTargetGroup
    Properties:
      DBProxyName: !Ref TestDBProxy
      DBInstanceIdentifiers: [!Ref InstanceName]
      TargetGroupName: default
      ConnectionPoolConfigurationInfo:
          MaxConnectionsPercent: 100
          MaxIdleConnectionsPercent: 50
          ConnectionBorrowTimeout: 120

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AuthFormat
