This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config

Creates a `Config` with the specified parameters.

Config objects provide Ground Station with the details necessary in order to schedule and execute satellite contacts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GroundStation::Config",
  "Properties" : {
      "ConfigData" : ConfigData,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GroundStation::Config
Properties:
  ConfigData:
    ConfigData
  Name: String
  Tags:
    - Tag

```

## Properties

`ConfigData`

Object containing the parameters of a config.
Only one subtype may be specified per config.
See the subtype definitions for a description of each config subtype.

_Required_: Yes

_Type_: [ConfigData](aws-properties-groundstation-config-configdata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the config object.

_Required_: Yes

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_:-]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags assigned to a resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-groundstation-config-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the config. For example:

`{ "Ref": "Config" }`

For the Ground Station config, `Ref` returns the ARN of the config.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the config, such as
`arn:aws:groundstation:us-east-2:012345678910:config/tracking/9940bf3b-d2ba-427e-9906-842b5e5d2296`.

`Id`

The ID of the config, such as
`9940bf3b-d2ba-427e-9906-842b5e5d2296`.

`Type`

The type of the config, such as
`tracking`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Ground Station

AntennaDownlinkConfig

All content copied from https://docs.aws.amazon.com/.
