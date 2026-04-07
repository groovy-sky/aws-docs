This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::DirectConnectGateway

Creates a Direct Connect gateway, which is an intermediate object that enables you to connect a set
of virtual interfaces and virtual private gateways. A Direct Connect gateway is global and visible in any
AWS Region after it is created. The virtual interfaces and virtual private gateways that
are connected through a Direct Connect gateway can be in different AWS Regions. This enables you to
connect to a VPC in any Region, regardless of the Region in which the virtual interfaces
are located, and pass traffic between them.

For more information, see [Direct Connect gateways](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-gateways-intro.html) in the _Direct Connect User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectConnect::DirectConnectGateway",
  "Properties" : {
      "AmazonSideAsn" : String,
      "DirectConnectGatewayName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DirectConnect::DirectConnectGateway
Properties:
  AmazonSideAsn: String
  DirectConnectGatewayName: String
  Tags:
    - Tag

```

## Properties

`AmazonSideAsn`

The autonomous system number (ASN) for Border Gateway Protocol (BGP) to be configured
on the Amazon side of the connection. The ASN must be in the private range of 64,512 to
65,534 or 4,200,000,000 to 4,294,967,294. The default is 64512.

_Required_: No

_Type_: String

_Pattern_: `^[1-9][0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DirectConnectGatewayName`

The name of the Direct Connect gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w \-_,\/]{1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Information about a tag.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directconnect-directconnectgateway-tag.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the Direct Connect gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DirectConnectGatewayArn`

The Amazon Resource Name (ARN) of the Direct Connect gateway.

`DirectConnectGatewayId`

The ID of the Direct Connect gateway.

## Examples

### Basic Direct Connect gateway

This example shows a basic Direct Connect gateway.

#### JSON

```json

{
  "Resources": {
    "myDirectConnectGateway": {
      "Type": "AWS::DirectConnect::DirectConnectGateway",
      "Properties": {
        "DirectConnectGatewayName": "cfn-directconnectgateway-example",
        "AmazonSideAsn": "65412",
        "Tags": [
          {
            "Key": "example-key",
            "Value": "example-value"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  myDirectConnectGateway:
    Type: AWS::DirectConnect::DirectConnectGateway
    Properties:
      DirectConnectGatewayName: cfn-directconnectgateway-example
      AmazonSideAsn: '65412'
      Tags:
      - Key: example-key
        Value: example-value
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
