This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::Map

The `AWS::Location::Map` resource specifies a map resource in your
AWS account, which provides map tiles of different styles sourced
from global location data providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Location::Map",
  "Properties" : {
      "Configuration" : MapConfiguration,
      "Description" : String,
      "MapName" : String,
      "PricingPlan" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Location::Map
Properties:
  Configuration:
    MapConfiguration
  Description: String
  MapName: String
  PricingPlan: String
  Tags:
    - Tag

```

## Properties

`Configuration`

Specifies the `MapConfiguration`, including the map style, for the map
resource that you create. The map style defines the look of maps and the data provider
for your map resource.

_Required_: Yes

_Type_: [MapConfiguration](aws-properties-location-map-mapconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

An optional description for the map resource.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapName`

The name for the map resource.

Requirements:

- Must contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-),
periods (.), and underscores (\_).

- Must be a unique map resource name.

- No spaces allowed. For example, `ExampleMap`.

_Required_: Yes

_Type_: String

_Pattern_: `^[-._\w]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PricingPlan`

No longer used. If included, the only allowed value is
`RequestBasedUsage`.

_Allowed Values_: `RequestBasedUsage`

_Required_: No

_Type_: String

_Allowed values_: `RequestBasedUsage`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Applies one or more tags to the map resource. A tag is a key-value pair helps manage,
identify, search, and filter your resources by labelling them.

Format: `"key" : "value"`

Restrictions:

- Maximum 50 tags per resource

- Each resource tag must be unique with a maximum of one value.

- Maximum key length: 128 Unicode characters in UTF-8

- Maximum value length: 256 Unicode characters in UTF-8

- Can use alphanumeric characters (A–Z, a–z, 0–9), and the following characters:
\+ \- = . \_ : / @.

- Cannot use "aws:" as a prefix for a key.

_Required_: No

_Type_: Array of [Tag](aws-properties-location-map-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Map` name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across
all AWS.

- Format example:
`arn:aws:geo:region:account-id:maps/ExampleMap`

`CreateTime`

The timestamp for when the map resource was created in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format:
`YYYY-MM-DDThh:mm:ss.sssZ`.

`MapArn`

Synonym for `Arn`. The Amazon Resource Name (ARN) for the map resource. Used to specify a resource across
all AWS.

- Format example:
`arn:aws:geo:region:account-id:maps/ExampleMap`

`UpdateTime`

The timestamp for when the map resource was last updated in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format:
`YYYY-MM-DDThh:mm:ss.sssZ`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

MapConfiguration

All content copied from https://docs.aws.amazon.com/.
