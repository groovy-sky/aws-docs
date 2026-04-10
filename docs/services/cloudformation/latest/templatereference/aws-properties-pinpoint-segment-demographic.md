This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment Demographic

Specifies demographic-based criteria, such as device platform, for the segment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppVersion" : SetDimension,
  "Channel" : SetDimension,
  "DeviceType" : SetDimension,
  "Make" : SetDimension,
  "Model" : SetDimension,
  "Platform" : SetDimension
}

```

### YAML

```yaml

  AppVersion:
    SetDimension
  Channel:
    SetDimension
  DeviceType:
    SetDimension
  Make:
    SetDimension
  Model:
    SetDimension
  Platform:
    SetDimension

```

## Properties

`AppVersion`

The app version criteria for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Channel`

The channel criteria for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceType`

The device type criteria for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Make`

The device make criteria for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

The device model criteria for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Platform`

The device platform criteria for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Coordinates

GPSPoint

All content copied from https://docs.aws.amazon.com/.
