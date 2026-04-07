This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::TransitGatewayRegistration

Registers a transit gateway in your global network. Not all Regions support transit
gateways for global networks. For a list of the supported Regions, see [Region Availability](https://docs.aws.amazon.com/network-manager/latest/tgwnm/what-are-global-networks.html#nm-available-regions) in the _AWS Transit Gateways for Global_
_Networks User Guide_. The transit gateway can be in any of the supported
AWS Regions, but it must be owned by the same AWS account that owns the global
network. You cannot register a transit gateway in more than one global network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::TransitGatewayRegistration",
  "Properties" : {
      "GlobalNetworkId" : String,
      "TransitGatewayArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::TransitGatewayRegistration
Properties:
  GlobalNetworkId: String
  TransitGatewayArn: String

```

## Properties

`GlobalNetworkId`

The ID of the global network.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayArn`

The Amazon Resource Name (ARN) of the transit gateway.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the global network and the ARN of the transit gateway. For example: `global-network-01231231231231231|arn:aws:ec2:us-west-2:123456789012:transit-gateway/tgw-123abc05e04123abc`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Transit Gateway Registration

The following example registers a transit gateway in a global network.

#### JSON

```json

{
    "Type": "AWS::NetworkManager::TransitGatewayRegistration",
    "Properties": {
        "GlobalNetworkId": {
            "Ref": "GlobalNetwork"
        },
        "TransitGatewayArn": {
            "Fn::Sub": "arn:aws:ec2:${AWS::Region}:${AWS::AccountId}:transit-gateway/${TransitGateway}"
        }
    }
}
```

#### YAML

```yaml

Type: AWS::NetworkManager::TransitGatewayRegistration
Properties:
  GlobalNetworkId: !Ref GlobalNetwork
  TransitGatewayArn: !Sub 'arn:aws:ec2:${AWS::Region}:${AWS::AccountId}:transit-gateway/${TransitGateway}'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::NetworkManager::TransitGatewayRouteTableAttachment
