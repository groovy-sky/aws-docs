This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel OutputGroup

The settings for one output group.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "OutputGroupSettings" : OutputGroupSettings,
  "Outputs" : [ Output, ... ]
}

```

### YAML

```yaml

  Name: String
  OutputGroupSettings:
    OutputGroupSettings
  Outputs:
    - Output

```

## Properties

`Name`

A custom output group name that you can optionally define. Only letters, numbers, and the underscore character
are allowed. The maximum length is 32 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputGroupSettings`

The settings associated with the output group.

_Required_: No

_Type_: [OutputGroupSettings](aws-properties-medialive-channel-outputgroupsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outputs`

The settings for the outputs in the output group.

_Required_: No

_Type_: Array of [Output](aws-properties-medialive-channel-output.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputDestinationSettings

OutputGroupSettings

All content copied from https://docs.aws.amazon.com/.
