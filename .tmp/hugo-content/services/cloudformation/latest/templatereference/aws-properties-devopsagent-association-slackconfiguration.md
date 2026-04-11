This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association SlackConfiguration

Configuration for Slack workspace integration. Defines the workspace ID, workspace name, and transmission
targets that specify which Slack channels receive notifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TransmissionTarget" : SlackTransmissionTarget,
  "WorkspaceId" : String,
  "WorkspaceName" : String
}

```

### YAML

```yaml

  TransmissionTarget:
    SlackTransmissionTarget
  WorkspaceId: String
  WorkspaceName: String

```

## Properties

`TransmissionTarget`

Transmission targets for agent notifications.

_Required_: Yes

_Type_: [SlackTransmissionTarget](aws-properties-devopsagent-association-slacktransmissiontarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceId`

The associated Slack workspace ID.

The ID must match the pattern `^[TE][A-Z0-9]+$`.

_Required_: Yes

_Type_: String

_Pattern_: `^[TE][A-Z0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceName`

Associated Slack workspace name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlackChannel

SlackTransmissionTarget

All content copied from https://docs.aws.amazon.com/.
