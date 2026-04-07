This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::Input InputDefinition

The definition of the input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ Attribute, ... ]
}

```

### YAML

```yaml

  Attributes:
    - Attribute

```

## Properties

`Attributes`

The attributes from the JSON payload that are made available by the input. Inputs are
derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such
message contains a JSON payload, and those attributes (and their paired values) specified here
are available for use in the `condition` expressions used by detectors that monitor
this input.

_Required_: Yes

_Type_: Array of [Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotevents-input-attribute.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Attribute

Tag
