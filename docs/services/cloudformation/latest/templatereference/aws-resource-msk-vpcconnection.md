This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::VpcConnection

Create remote VPC connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::VpcConnection",
  "Properties" : {
      "Authentication" : String,
      "ClientSubnets" : [ String, ... ],
      "SecurityGroups" : [ String, ... ],
      "Tags" : {Key: Value, ...},
      "TargetClusterArn" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::MSK::VpcConnection
Properties:
  Authentication: String
  ClientSubnets:
    - String
  SecurityGroups:
    - String
  Tags:
    Key: Value
  TargetClusterArn: String
  VpcId: String

```

## Properties

`Authentication`

The type of private link authentication.

_Required_: Yes

_Type_: String

_Allowed values_: `SASL_IAM | SASL_SCRAM | TLS`

_Minimum_: `3`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientSubnets`

The list of subnets in the client VPC to connect to.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

The security groups to attach to the ENIs for the broker nodes.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An arbitrary set of tags (key-value pairs) you specify while creating the VPC connection.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetClusterArn`

The Amazon Resource Name (ARN) of the cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[\w-]+:kafka:[\w-]+:\d+:cluster.*\Z`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The VPC ID of the remote client.

_Required_: Yes

_Type_: String

_Pattern_: `^(vpc-)([a-z0-9]+)\Z`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the VPC connection.

For Amazon MSK VPC connection `MyVpcConnection`, `Ref` returns the ARN of the VPC connection whose logical ID is `MyVpcConnection`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the VPC connection.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MSK::Topic

Next
