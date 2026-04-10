This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::InternetGateway

Allocates an internet gateway for use with a VPC. After creating the Internet gateway,
you then attach it to a VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::InternetGateway",
  "Properties" : {
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::InternetGateway
Properties:
  Tags:
    - Tag

```

## Properties

`Tags`

Any tags to assign to the internet gateway.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-internetgateway-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the internet gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`InternetGatewayId`

The ID of the internet gateway.

## Examples

### Create an internet gateway

The following example creates an internet gateway and assigns it a tag.

#### JSON

```json

"Resources" : {
   "myInternetGateway" : {
      "Type" : "AWS::EC2::InternetGateway",
      "Properties" : {
        "Tags" : [ {"Key" : "stack", "Value" : "production"}]
      }
   }
}
```

#### YAML

```yaml

  myInternetGateway:
    Type: AWS::EC2::InternetGateway
    Properties:
      Tags:
      - Key: stack
        Value: production
```

## See also

- [CreateInternetGateway](../../../../reference/awsec2/latest/apireference/api-createinternetgateway.md) in the _Amazon EC2 API_
_Reference_

- [Internet gateways](../../../vpc/latest/userguide/vpc-internet-gateway.md)
in the _Amazon VPC User Guide_

- Use the [AWS::EC2::VPCGatewayAttachment](../userguide/aws-resource-ec2-vpc-gateway-attachment.md) resource to associate an internet
gateway with a VPC

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
