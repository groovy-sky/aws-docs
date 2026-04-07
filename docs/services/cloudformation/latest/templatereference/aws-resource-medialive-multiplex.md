This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplex

The multiplex object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::Multiplex",
  "Properties" : {
      "AvailabilityZones" : [ String, ... ],
      "Destinations" : [ MultiplexOutputDestination, ... ],
      "MultiplexSettings" : MultiplexSettings,
      "Name" : String,
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::Multiplex
Properties:
  AvailabilityZones:
    - String
  Destinations:
    - MultiplexOutputDestination
  MultiplexSettings:
    MultiplexSettings
  Name: String
  Tags:
    - Tags

```

## Properties

`AvailabilityZones`

A list of availability zones for the multiplex.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Destinations`

A list of the multiplex output destinations.

_Required_: No

_Type_: Array of [MultiplexOutputDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplex-multiplexoutputdestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiplexSettings`

Configuration for a multiplex event.

_Required_: Yes

_Type_: [MultiplexSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplex-multiplexsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the multiplex.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A collection of key-value pairs.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplex-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplex-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The unique arn of the multiplex.

`Id`

The unique id of the multiplex.

`PipelinesRunningCount`

The number of currently healthy pipelines.

`ProgramCount`

The number of programs in the multiplex.

`State`

The current state of the multiplex.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InputWhitelistRuleCidr

MultiplexMediaConnectOutputDestinationSettings
