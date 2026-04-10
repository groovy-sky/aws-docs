This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::CustomerGatewayAssociation

Specifies an association between a customer gateway, a device, and optionally, a link.
If you specify a link, it must be associated with the specified device. The customer
gateway must be connected to a VPN attachment on a transit gateway that's registered in
your global network.

You cannot associate a customer gateway with more than one device and link.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::CustomerGatewayAssociation",
  "Properties" : {
      "CustomerGatewayArn" : String,
      "DeviceId" : String,
      "GlobalNetworkId" : String,
      "LinkId" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::CustomerGatewayAssociation
Properties:
  CustomerGatewayArn: String
  DeviceId: String
  GlobalNetworkId: String
  LinkId: String

```

## Properties

`CustomerGatewayArn`

The Amazon Resource Name (ARN) of the customer gateway.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeviceId`

The ID of the device.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GlobalNetworkId`

The ID of the global network.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LinkId`

The ID of the link.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the global network and the Amazon Resource Name
(ARN) of the customer gateway. For example: `global-network-01231231231231231|arn:aws:ec2:eu-central-1:123456789012:customer-gateway/cgw-00112233aabbcc112`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Customer Gateway Association

The following example template creates a global network, device, customer
gateway, VPN connection, and transit gateway. It registers the transit gateway
in the global network, and creates an association between the customer gateway
and device. The creation of the customer gateway association depends on the VPN
connection and transit gateway registration.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Create a global network and customer gateway association",
    "Resources": {
        "GlobalNetwork": {
            "Type": "AWS::NetworkManager::GlobalNetwork"
        },
        "Device": {
            "Type": "AWS::NetworkManager::Device",
            "Properties": {
                "Description": "Chicago office device",
                "GlobalNetworkId": {
                    "Ref": "GlobalNetwork"
                },
                "Location": {
                    "Address": "227 W Monroe St, Chicago, IL 60606",
                    "Latitude": "41.8",
                    "Longitude": "-87.6"
                }
            }
        },
        "TransitGateway": {
            "Type": "AWS::EC2::TransitGateway"
        },
        "TransitGatewayRegistration": {
            "Type": "AWS::NetworkManager::TransitGatewayRegistration",
            "Properties": {
                "GlobalNetworkId": {
                    "Ref": "GlobalNetwork"
                },
                "TransitGatewayArn": {
                    "Fn::Sub": "arn:aws:ec2:${AWS::Region}:${AWS::AccountId}:transit-gateway/${TransitGateway}"
                }
            }
        },
        "CustomerGateway": {
            "Type": "AWS::EC2::CustomerGateway",
            "Properties": {
                "Type": "ipsec.1",
                "BgpAsn": 65534,
                "IpAddress": "12.1.2.3"
            }
        },
        "VPNConnection": {
            "Type": "AWS::EC2::VPNConnection",
            "Properties": {
                "Type": "ipsec.1",
                "StaticRoutesOnly": true,
                "CustomerGatewayId": {
                    "Ref": "CustomerGateway"
                },
                "TransitGatewayId": {
                    "Ref": "TransitGateway"
                }
            }
        },
        "CustomerGatewayAssociation": {
            "DependsOn": [
                "VPNConnection",
                "TransitGatewayRegistration"
            ],
            "Type": "AWS::NetworkManager::CustomerGatewayAssociation",
            "Properties": {
                "GlobalNetworkId": {
                    "Ref": "GlobalNetwork"
                },
                "DeviceId": {
                    "Fn::GetAtt": [
                        "Device",
                        "DeviceId"
                    ]
                },
                "CustomerGatewayArn": {
                    "Fn::Sub": "arn:aws:ec2:${AWS::Region}:${AWS::AccountId}:customer-gateway/${CustomerGateway}"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: 'Create a global network and customer gateway association'
Resources:
  GlobalNetwork:
    Type: AWS::NetworkManager::GlobalNetwork
  Device:
    Type: AWS::NetworkManager::Device
    Properties:
      Description: Chicago office device
      GlobalNetworkId: !Ref GlobalNetwork
      Location:
        Address: "227 W Monroe St, Chicago, IL 60606"
        Latitude: "41.8"
        Longitude: "-87.6"
  TransitGateway:
    Type: AWS::EC2::TransitGateway
  TransitGatewayRegistration:
    Type: AWS::NetworkManager::TransitGatewayRegistration
    Properties:
      GlobalNetworkId: !Ref GlobalNetwork
      TransitGatewayArn: !Sub 'arn:aws:ec2:${AWS::Region}:${AWS::AccountId}:transit-gateway/${TransitGateway}'
  CustomerGateway:
    Type: AWS::EC2::CustomerGateway
    Properties:
      Type: ipsec.1
      BgpAsn: 65534
      IpAddress: 12.1.2.3
  VPNConnection:
    Type: AWS::EC2::VPNConnection
    Properties:
      Type: ipsec.1
      StaticRoutesOnly: true
      CustomerGatewayId:
        !Ref CustomerGateway
      TransitGatewayId:
        !Ref TransitGateway
  CustomerGatewayAssociation:
    DependsOn:
      - VPNConnection
      - TransitGatewayRegistration
    Type: AWS::NetworkManager::CustomerGatewayAssociation
    Properties:
      GlobalNetworkId: !Ref GlobalNetwork
      DeviceId: !GetAtt Device.DeviceId
      CustomerGatewayArn: !Sub 'arn:aws:ec2:${AWS::Region}:${AWS::AccountId}:customer-gateway/${CustomerGateway}'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::NetworkManager::CoreNetworkPrefixListAssociation

AWS::NetworkManager::Device

All content copied from https://docs.aws.amazon.com/.
