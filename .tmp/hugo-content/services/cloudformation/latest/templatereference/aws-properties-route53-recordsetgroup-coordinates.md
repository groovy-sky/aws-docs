This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::RecordSetGroup Coordinates

A complex type that lists the coordinates for a geoproximity resource record.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Latitude" : String,
  "Longitude" : String
}

```

### YAML

```yaml

  Latitude: String
  Longitude: String

```

## Properties

`Latitude`

Specifies a coordinate of the north–south position of a geographic point on the surface of
the Earth (-90 - 90).

_Required_: Yes

_Type_: String

_Pattern_: `[-+]?[0-9]{1,2}(\.[0-9]{0,2})?`

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Longitude`

Specifies a coordinate of the east–west position of a geographic point on the surface of
the Earth (-180 - 180).

_Required_: Yes

_Type_: String

_Pattern_: `[-+]?[0-9]{1,3}(\.[0-9]{0,2})?`

_Minimum_: `1`

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CidrRoutingConfig

GeoLocation

All content copied from https://docs.aws.amazon.com/.
