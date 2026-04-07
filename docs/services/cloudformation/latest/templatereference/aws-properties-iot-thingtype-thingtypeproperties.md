This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingType ThingTypeProperties

The ThingTypeProperties contains information about the thing type including: a thing type description,
and a list of searchable thing attribute names.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mqtt5Configuration" : Mqtt5Configuration,
  "SearchableAttributes" : [ String, ... ],
  "ThingTypeDescription" : String
}

```

### YAML

```yaml

  Mqtt5Configuration:
    Mqtt5Configuration
  SearchableAttributes:
    - String
  ThingTypeDescription: String

```

## Properties

`Mqtt5Configuration`

The configuration to add user-defined properties to enrich MQTT 5 messages.

_Required_: No

_Type_: [Mqtt5Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-thingtype-mqtt5configuration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SearchableAttributes`

A list of searchable thing attribute names.

_Required_: No

_Type_: Array of String

_Maximum_: `128 | 3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingTypeDescription`

The description of the thing type.

_Required_: No

_Type_: String

_Pattern_: `[\p{Graph}\x20]*`

_Maximum_: `2028`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IoT::TopicRule
