This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment Location

Specifies location-based criteria, such as region or GPS coordinates, for the
segment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Country" : SetDimension,
  "GPSPoint" : GPSPoint
}

```

### YAML

```yaml

  Country:
    SetDimension
  GPSPoint:
    GPSPoint

```

## Properties

`Country`

The country or region code, in ISO 3166-1 alpha-2 format, for the segment.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-segment-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GPSPoint`

The GPS point dimension for the segment.

_Required_: No

_Type_: [GPSPoint](aws-properties-pinpoint-segment-gpspoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Groups

Recency

All content copied from https://docs.aws.amazon.com/.
