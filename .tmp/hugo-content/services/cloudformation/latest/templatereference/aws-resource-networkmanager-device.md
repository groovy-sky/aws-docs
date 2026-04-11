This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::Device

Specifies a device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::Device",
  "Properties" : {
      "AWSLocation" : AWSLocation,
      "Description" : String,
      "GlobalNetworkId" : String,
      "Location" : Location,
      "Model" : String,
      "SerialNumber" : String,
      "SiteId" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "Vendor" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::Device
Properties:
  AWSLocation:
    AWSLocation
  Description: String
  GlobalNetworkId: String
  Location:
    Location
  Model: String
  SerialNumber: String
  SiteId: String
  Tags:
    - Tag
  Type: String
  Vendor: String

```

## Properties

`AWSLocation`

The AWS location of the device.

_Required_: No

_Type_: [AWSLocation](aws-properties-networkmanager-device-awslocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the device.

Constraints: Maximum length of 256 characters.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalNetworkId`

The ID of the global network.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Location`

The site location.

_Required_: No

_Type_: [Location](aws-properties-networkmanager-device-location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

The model of the device.

Constraints: Maximum length of 128 characters.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SerialNumber`

The serial number of the device.

Constraints: Maximum length of 128 characters.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SiteId`

The site ID.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the device.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-device-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The device type.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vendor`

The vendor of the device.

Constraints: Maximum length of 128 characters.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the IDs of the global network and device. For example: `global-network-01231231231231231|device-07f6fd08867abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time that the device was created.

`DeviceArn`

The ARN of the device. For example,
`arn:aws:networkmanager::123456789012:device/global-network-01231231231231231/device-07f6fd08867abc123`.

`DeviceId`

The ID of the device. For example, `device-07f6fd08867abc123`.

`State`

The state of the device.

## Examples

### Device

The following example creates a device in a global network.

#### JSON

```json

{
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
        "Location": {
            "Address": "227 W Monroe St, Chicago, IL 60606",
            "Latitude": "41.8",
            "Longitude": "-87.6"
        },
        "Tags": [
            {
                "Key": "Network",
                "Value": "north-america"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: AWS::NetworkManager::Device
Properties:
  Description: "Chicago office device"
  GlobalNetworkId: !Ref GlobalNetwork
  SiteId: !GetAtt Site.SiteId
  Location:
    Address: "227 W Monroe St, Chicago, IL 60606"
    Latitude: "41.8"
    Longitude: "-87.6"
  Tags:
    - Key: Network
      Value: north-america
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::NetworkManager::CustomerGatewayAssociation

AWSLocation

All content copied from https://docs.aws.amazon.com/.
