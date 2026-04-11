This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice

Provisions a wireless device.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::WirelessDevice",
  "Properties" : {
      "Description" : String,
      "DestinationName" : String,
      "LastUplinkReceivedAt" : String,
      "LoRaWAN" : LoRaWANDevice,
      "Name" : String,
      "Positioning" : String,
      "Tags" : [ Tag, ... ],
      "ThingArn" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::WirelessDevice
Properties:
  Description: String
  DestinationName: String
  LastUplinkReceivedAt: String
  LoRaWAN:
    LoRaWANDevice
  Name: String
  Positioning: String
  Tags:
    - Tag
  ThingArn: String
  Type: String

```

## Properties

`Description`

The description of the new resource. Maximum length is 2048.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationName`

The name of the destination to assign to the new wireless device. Can have only have
alphanumeric, - (hyphen) and \_ (underscore) characters and it can't have any spaces.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUplinkReceivedAt`

The date and time when the most recent uplink was received.

_Required_: No

_Type_: String

_Pattern_: `^([\+-]?\d{4}(?!\d{2}\b))((-?)((0[1-9]|1[0-2])(\3([12]\d|0[1-9]|3[01]))?|W([0-4]\d|5[0-2])(-?[1-7])?|(00[1-9]|0[1-9]\d|[12]\d{2}|3([0-5]\d|6[1-6])))([T\s]((([01]\d|2[0-3])((:?)[0-5]\d)?|24\:?00)([\.,]\d+(?!:))?)?(\17[0-5]\d([\.,]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoRaWAN`

The device configuration information to use to create the wireless device. Must be at
least one of OtaaV10x, OtaaV11, AbpV11, or AbpV10x.

_Required_: No

_Type_: [LoRaWANDevice](aws-properties-iotwireless-wirelessdevice-lorawandevice.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the new resource.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Positioning`

FPort values for the GNSS, Stream, and ClockSync functions of the positioning
information.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-wirelessdevice-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingArn`

The ARN of the thing to associate with the wireless device.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The wireless device type.

_Required_: Yes

_Type_: String

_Allowed values_: `Sidewalk | LoRaWAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the wireless device ID.

### Fn::GetAtt

`Arn`

The ARN of the wireless device created.

`Id`

The ID of the wireless device created.

`ThingName`

The name of the thing associated with the wireless device. The value is empty if a thing
isn't associated with the device.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateWirelessGatewayTaskCreate

AbpV10x

All content copied from https://docs.aws.amazon.com/.
