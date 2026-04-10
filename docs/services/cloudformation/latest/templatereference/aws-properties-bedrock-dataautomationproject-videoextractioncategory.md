This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject VideoExtractionCategory

Settings for generating categorical data from video.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String,
  "Types" : [ String, ... ]
}

```

### YAML

```yaml

  State: String
  Types:
    - String

```

## Properties

`State`

Whether generating categorical data from video is enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Types`

The types of data to generate.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoBoundingBox

VideoOverrideConfiguration

All content copied from https://docs.aws.amazon.com/.
