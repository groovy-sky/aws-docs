This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::LinkAssociation

Describes the association between a device and a link.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::LinkAssociation",
  "Properties" : {
      "DeviceId" : String,
      "GlobalNetworkId" : String,
      "LinkId" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::LinkAssociation
Properties:
  DeviceId: String
  GlobalNetworkId: String
  LinkId: String

```

## Properties

`DeviceId`

The device ID for the link association.

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

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IDs of the global network, device, and link. For example: `global-network-01231231231231231|device-07f6fd08867abc123|link-11112222aaaabbbb1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Link Association

The following example template creates a global network, site, link, and device. It
creates an association between the link and the device.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Create global network and link association",
    "Resources": {
        "GlobalNetwork": {
            "Type": "AWS::NetworkManager::GlobalNetwork"
        },
        "Site": {
            "Type": "AWS::NetworkManager::Site",
            "Properties": {
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
        "Link": {
            "Type": "AWS::NetworkManager::Link",
            "Properties": {
                "Description": "Broadband link",
                "GlobalNetworkId": {
                    "Ref": "GlobalNetwork"
                },
                "SiteId": {
                    "Fn::GetAtt": [
                        "Site",
                        "SiteId"
                    ]
                },
                "Bandwidth": {
                    "DownloadSpeed": 20,
                    "UploadSpeed": 20
                },
                "Provider": "AnyCompany",
                "Type": "Broadband",
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "broadband-link-1"
                    }
                ]
            }
        },
        "Device": {
            "Type": "AWS::NetworkManager::Device",
            "Properties": {
                "Description": "Chicago office device",
                "GlobalNetworkId": {
                    "Ref": "GlobalNetwork"
                },
                "SiteId": {
                    "Fn::GetAtt": [
                        "Site",
                        "SiteId"
                    ]
                },
                "Tags": [
                    {
                        "Key": "Network",
                        "Value": "north-america"
                    }
                ]
            }
        },
        "LinkAssociation": {
            "Type": "AWS::NetworkManager::LinkAssociation",
            "Properties": {
                "GlobalNetworkId": {
                    "Ref": "GlobalNetwork"
                },
                "LinkId": {
                    "Fn::GetAtt": [
                        "Link",
                        "LinkId"
                    ]
                },
                "DeviceId": {
                    "Fn::GetAtt": [
                        "Device",
                        "DeviceId"
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: 'Create global network and link association'
Resources:
  GlobalNetwork:
    Type: AWS::NetworkManager::GlobalNetwork
  Site:
    Type: AWS::NetworkManager::Site
    Properties:
      GlobalNetworkId: !Ref GlobalNetwork
      Location:
        Address: "227 W Monroe St, Chicago, IL 60606"
        Latitude: "41.8"
        Longitude: "-87.6"
  Link:
    Type: AWS::NetworkManager::Link
    Properties:
      Description: Broadband link
      GlobalNetworkId: !Ref GlobalNetwork
      SiteId: !GetAtt Site.SiteId
      Bandwidth:
        DownloadSpeed: 20
        UploadSpeed: 20
      Provider: "AnyCompany"
      Type: "Broadband"
      Tags:
        - Key: Name
          Value: broadband-link-1
  Device:
    Type: AWS::NetworkManager::Device
    Properties:
      Description: Chicago office device
      GlobalNetworkId: !Ref GlobalNetwork
      SiteId: !GetAtt Site.SiteId
      Tags:
        - Key: Network
          Value: north-america
  LinkAssociation:
    Type: AWS::NetworkManager::LinkAssociation
    Properties:
      GlobalNetworkId: !Ref GlobalNetwork
      LinkId: !GetAtt Link.LinkId
      DeviceId: !GetAtt Device.DeviceId
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::NetworkManager::Site
