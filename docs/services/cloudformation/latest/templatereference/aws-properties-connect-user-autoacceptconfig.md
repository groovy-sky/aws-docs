This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User AutoAcceptConfig

Configuration settings for auto-accept for a specific channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentFirstCallbackAutoAccept" : Boolean,
  "AutoAccept" : Boolean,
  "Channel" : String
}

```

### YAML

```yaml

  AgentFirstCallbackAutoAccept: Boolean
  AutoAccept: Boolean
  Channel: String

```

## Properties

`AgentFirstCallbackAutoAccept`

Indicates whether auto-accept is enabled for agent-first callbacks. This setting only applies to the VOICE channel.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoAccept`

Indicates whether auto-accept is enabled for this channel. When enabled, available agents are automatically connected to contacts from this channel.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Channel`

The channel for this auto-accept configuration. Valid values: VOICE, CHAT, TASK, EMAIL.

_Required_: Yes

_Type_: String

_Allowed values_: `VOICE | CHAT | TASK | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AfterContactWorkConfigPerChannel

PersistentConnectionConfig

All content copied from https://docs.aws.amazon.com/.
