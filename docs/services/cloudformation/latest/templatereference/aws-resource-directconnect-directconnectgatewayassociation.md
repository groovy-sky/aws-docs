This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::DirectConnectGatewayAssociation

Creates an association between a Direct Connect gateway and a virtual private gateway. The virtual
private gateway must be attached to a VPC and must not be associated with another Direct Connect gateway.

If creating a Direct Connect gateway association between resources in separate AWS accounts, the CloudFormation stack account must own the virtual private gateway. The other account must own the Direct Connect gateway, and must have a role allowing the stack account to accept Direct Connect gateway association proposals.

For more information, see [Direct Connect gateways](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-gateways-intro.html) in the _Direct Connect User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectConnect::DirectConnectGatewayAssociation",
  "Properties" : {
      "AcceptDirectConnectGatewayAssociationProposalRoleArn" : String,
      "AllowedPrefixesToDirectConnectGateway" : [ String, ... ],
      "AssociatedGatewayId" : String,
      "DirectConnectGatewayId" : String
    }
}

```

### YAML

```yaml

Type: AWS::DirectConnect::DirectConnectGatewayAssociation
Properties:
  AcceptDirectConnectGatewayAssociationProposalRoleArn: String
  AllowedPrefixesToDirectConnectGateway:
    - String
  AssociatedGatewayId: String
  DirectConnectGatewayId: String

```

## Properties

`AcceptDirectConnectGatewayAssociationProposalRoleArn`

The Amazon Resource Name (ARN) of the role in a separate account to accept the Direct Connect Gateway association proposal. The role needs to have `directconnect:AcceptDirectConnectGatewayAssociationProposal` permissions.

###### Note

This should only be used when creating an association with a Direct Connect gateway in a separate AWS account.

_Required_: Conditional

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AllowedPrefixesToDirectConnectGateway`

The Amazon VPC prefixes to advertise to the Direct Connect gateway.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociatedGatewayId`

The ID or ARN of the associated gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^((arn:aws[a-z-]*:ec2:[a-z0-9-]+:[0-9]{12}:(vpn-gateway/vgw|transit-gateway/tgw))|(vgw|tgw))-[a-zA-Z0-9]{8,32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DirectConnectGatewayId`

The ID or ARN of the Direct Connect gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[a-z-]*:directconnect::[0-9]{12}:dx-gateway/)?[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Direct Connect gateway association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationId`

The ID of the Direct Connect gateway association.

## Examples

### Direct Connect gateway association with a virtual private gateway in the same account

This example shows a basic Direct Connect gateway association with a virtual private gateway.

#### JSON

```json

{
  "Resources": {
    "myDirectConnectGateway": {
      "Type": "AWS::DirectConnect::DirectConnectGateway",
      "Properties": {
        "DirectConnectGatewayName": "cfn-directconnectgateway-example",
        "AmazonSideAsn": "65412"
      }
    },
    "myDirectConnectGatewayAssociation": {
      "Type": "AWS::DirectConnect::DirectConnectGatewayAssociation",
      "Properties": {
        "AssociatedGatewayId": "vgw-aba37db6",
        "DirectConnectGatewayId": {
          "Ref": "myDirectConnectGateway"
        },
        "AllowedPrefixesToDirectConnectGateway": [
          "10.0.0.0/24"
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
  myDirectConnectGatewayAssociation:
    Type: AWS::DirectConnect::DirectConnectGatewayAssociation
    Properties:
      AssociationGatewayId: vgw-aba37db6
      DirectConnectGatewayId: !Ref myDirectConnectGateway
      AllowedPrefixesToDirectConnectGateway:
      - 10.0.0.0/24
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::DirectConnect::Lag
