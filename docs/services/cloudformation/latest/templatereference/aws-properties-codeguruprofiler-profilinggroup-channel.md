This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeGuruProfiler::ProfilingGroup Channel

Notification medium for users to get alerted for events that occur in application profile. We support SNS topic as a notification channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "channelId" : String,
  "channelUri" : String
}

```

### YAML

```yaml

  channelId: String
  channelUri: String

```

## Properties

`channelId`

The channel ID.

_Required_: No

_Type_: String

_Pattern_: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`channelUri`

The channel URI.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws([-\w]*):[a-z-]+:(([a-z]+-)+[0-9]+)?:([0-9]{12}):[^.]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentPermissions

Tag

All content copied from https://docs.aws.amazon.com/.
