This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject ChannelLabelingConfiguration

Enables or disables channel labeling. Channel labeling, when enabled will assign a number to each audio channel, and indicate which channel is being used in each portion of the transcript.
This appears in the response as "ch\_0" for the first channel, and "ch\_1" for the second.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String
}

```

### YAML

```yaml

  State: String

```

## Properties

`State`

State of channel labeling, either enabled or disabled.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlueprintItem

CustomOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
