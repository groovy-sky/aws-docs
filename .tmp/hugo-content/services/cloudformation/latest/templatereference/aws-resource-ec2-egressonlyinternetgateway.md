This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EgressOnlyInternetGateway

\[IPv6 only\] Specifies an egress-only internet gateway for your VPC. An egress-only
internet gateway is used to enable outbound communication over IPv6 from instances in your
VPC to the internet, and prevents hosts outside of your VPC from initiating an IPv6
connection with your instance.

For more information, see [Egress-only internet gateway](../../../vpc/latest/userguide/egress-only-internet-gateway.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::EgressOnlyInternetGateway",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::EgressOnlyInternetGateway
Properties:
  Tags:
    - Tag
  VpcId: String

```

## Properties

`Tags`

The tags assigned to the egress-only internet gateway.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-egressonlyinternetgateway-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC for which to create the egress-only internet gateway.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the egress-only internet gateway (the physical
resource ID).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the egress-only internet gateway.

## Examples

### Create an egress-only internet gateway

The following example creates an egress-only internet gateway for the
specified VPC.

#### YAML

```yaml

myEgressOnlyInternetGateway:
    Type: AWS::EC2::EgressOnlyInternetGateway
    Properties:
        VpcId: vpc-1a2b3c4d
```

#### JSON

```json

{
    "myEgressOnlyInternetGateway" : {
        "Type" : "AWS::EC2::EgressOnlyInternetGateway",
        "Properties" : {
            "VpcId" : "vpc-1a2b3c4d"
        }
    }
}
```

## See also

- [CreateEgressOnlyInternetGateway](../../../../reference/awsec2/latest/apireference/apireference-query-createegressonlyinternetgateway.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VCpuCountRangeRequest

Tag

All content copied from https://docs.aws.amazon.com/.
