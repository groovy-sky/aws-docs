This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline Channel

Determines the source of the messages to be processed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelName" : String,
  "Name" : String,
  "Next" : String
}

```

### YAML

```yaml

  ChannelName: String
  Name: String
  Next: String

```

## Properties

`ChannelName`

The name of the channel from which the messages are processed.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the 'channel' activity.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Next`

The next activity in the pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddAttributes

Datastore

All content copied from https://docs.aws.amazon.com/.
