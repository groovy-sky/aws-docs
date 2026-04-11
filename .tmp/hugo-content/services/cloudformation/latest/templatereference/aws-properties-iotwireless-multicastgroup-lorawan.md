This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::MulticastGroup LoRaWAN

The LoRaWAN information that is to be used with the multicast group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DlClass" : String,
  "NumberOfDevicesInGroup" : Integer,
  "NumberOfDevicesRequested" : Integer,
  "RfRegion" : String
}

```

### YAML

```yaml

  DlClass: String
  NumberOfDevicesInGroup: Integer
  NumberOfDevicesRequested: Integer
  RfRegion: String

```

## Properties

`DlClass`

DlClass for LoRaWAN. Valid values are ClassB and ClassC.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfDevicesInGroup`

Number of devices that are associated to the multicast group.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfDevicesRequested`

Number of devices that are requested to be associated with the multicast group.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RfRegion`

The frequency band (RFRegion) value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTWireless::MulticastGroup

Tag

All content copied from https://docs.aws.amazon.com/.
