This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VideoSelectorProgramId

Used to extract video by the program ID.

The parent of this entity is VideoSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProgramId" : Integer
}

```

### YAML

```yaml

  ProgramId: Integer

```

## Properties

`ProgramId`

Selects a specific program from within a multi-program transport stream. If the program doesn't exist, MediaLive
selects the first program within the transport stream by default.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoSelectorPid

VideoSelectorSettings

All content copied from https://docs.aws.amazon.com/.
