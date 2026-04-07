This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessEndpoint RdsOptions

Describes the RDS options for a Verified Access endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : Integer,
  "Protocol" : String,
  "RdsDbClusterArn" : String,
  "RdsDbInstanceArn" : String,
  "RdsDbProxyArn" : String,
  "RdsEndpoint" : String,
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  Port: Integer
  Protocol: String
  RdsDbClusterArn: String
  RdsDbInstanceArn: String
  RdsDbProxyArn: String
  RdsEndpoint: String
  SubnetIds:
    - String

```

## Properties

`Port`

The port.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol.

_Required_: No

_Type_: String

_Allowed values_: `http | https | tcp`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RdsDbClusterArn`

The ARN of the DB cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RdsDbInstanceArn`

The ARN of the RDS instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RdsDbProxyArn`

The ARN of the RDS proxy.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RdsEndpoint`

The RDS endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The IDs of the subnets. You can specify only one subnet per Availability Zone.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PortRange

SseSpecification
