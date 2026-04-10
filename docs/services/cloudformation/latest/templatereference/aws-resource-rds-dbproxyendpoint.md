This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBProxyEndpoint

The `AWS::RDS::DBProxyEndpoint` resource creates or updates a DB proxy endpoint. You can use custom proxy endpoints to access a proxy through a different
VPC than the proxy's default VPC.

For more information about RDS Proxy, see
[AWS::RDS::DBProxy](../userguide/aws-resource-rds-dbproxy.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBProxyEndpoint",
  "Properties" : {
      "DBProxyEndpointName" : String,
      "DBProxyName" : String,
      "EndpointNetworkType" : String,
      "Tags" : [ TagFormat, ... ],
      "TargetRole" : String,
      "VpcSecurityGroupIds" : [ String, ... ],
      "VpcSubnetIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBProxyEndpoint
Properties:
  DBProxyEndpointName: String
  DBProxyName: String
  EndpointNetworkType: String
  Tags:
    - TagFormat
  TargetRole: String
  VpcSecurityGroupIds:
    - String
  VpcSubnetIds:
    - String

```

## Properties

`DBProxyEndpointName`

The name of the DB proxy endpoint to create.

_Required_: Yes

_Type_: String

_Pattern_: `[0-z]*`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBProxyName`

The name of the DB proxy associated with the DB proxy endpoint that you create.

_Required_: Yes

_Type_: String

_Pattern_: `[0-z]*`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

`Tags`

An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.

_Required_: No

_Type_: Array of [TagFormat](aws-properties-rds-dbproxyendpoint-tagformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRole`

A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.

_Required_: No

_Type_: String

_Allowed values_: `READ_WRITE | READ_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

The VPC security group IDs for the DB proxy endpoint that you create. You can
specify a different set of security group IDs than for the original DB proxy.
The default is the default security group for the VPC.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSubnetIds`

The VPC subnet IDs for the DB proxy endpoint that you create. You can specify a
different set of subnet IDs than for the original DB proxy.

_Required_: Yes

_Type_: Array of String

_Minimum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the DB proxy endpoint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DBProxyEndpointArn`

The Amazon Resource Name (ARN) for the DB proxy endpoint.

`Endpoint`

The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the
connection string for a database client application.

`IsDefault`

Indicates whether this endpoint is the default endpoint for the associated DB proxy.
Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the
DB proxy can be either read/write or read-only.

`VpcId`

Provides the VPC ID of the DB proxy endpoint.

## Examples

### Creating a custom DB proxy endpoint

The following example creates a custom DB proxy endpoint.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters": {
    "ProxyEndpointName": {
      "Type": "String",
      "Default": "exampleProxyEndpoint"
    },
    "ProxyName": {
      "Type": "String",
      "Default": "exampleProxy"
    },
    "SecurityGroupIds": {
      "Type": "String",
      "Default": "sg-12345678"
    },
    "SubnetIds": {
      "Type": "String",
      "Default": "subnet-12345677,subnet-12345678,subnet-12345679"
    }
  },
  "Resources": {
    "TestDBProxy": {
      "Type": "AWS::RDS::DBProxyEndpoint",
      "Properties": {
        "DBProxyEndpointName": {
          "Ref": "ProxyEndpointName"
        },
        "DBProxyName": {
          "Ref": "ProxyName"
        },
        "VpcSecurityGroupIds": {
          "Fn::Split": [
            ",",
            {
              "Ref": "SecurityGroupIds"
            }
          ]
        },
        "VpcSubnetIds": {
          "Fn::Split": [
            ",",
            {
              "Ref": "SubnetIds"
            }
          ]
        },
        "TargetRole": "READ_ONLY"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  ProxyEndpointName:
    Type: String
    Default: exampleProxyEndpoint
  ProxyName:
    Type: String
    Default: exampleProxy
  SubnetIds:
    Type: String
    Default: subnet-12345677,subnet-12345678,subnet-12345679
  SecurityGroupIds:
    Type: String
    Default: sg-12345678
Resources:
  TestDBProxyEndpoint:
    Type: AWS::RDS::DBProxyEndpoint
    Properties:
      DBProxyEndpointName: !Ref ProxyEndpointName
      DBProxyName: !Ref ProxyName
      VpcSubnetIds:
        Fn::Split: [",", !Ref SubnetIds]
      VpcSecurityGroupIds:
        Fn::Split: [",", !Ref SecurityGroupIds]
      TargetRole: READ_ONLY
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFormat

TagFormat

All content copied from https://docs.aws.amazon.com/.
