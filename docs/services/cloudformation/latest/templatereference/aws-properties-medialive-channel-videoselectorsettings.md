This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VideoSelectorSettings

Information about the video to extract from the input.

The parent of this entity is VideoSelector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VideoSelectorPid" : VideoSelectorPid,
  "VideoSelectorProgramId" : VideoSelectorProgramId
}

```

### YAML

```yaml

  VideoSelectorPid:
    VideoSelectorPid
  VideoSelectorProgramId:
    VideoSelectorProgramId

```

## Properties

`VideoSelectorPid`

Used to extract video by PID.

_Required_: No

_Type_: [VideoSelectorPid](aws-properties-medialive-channel-videoselectorpid.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoSelectorProgramId`

Used to extract video by program ID.

_Required_: No

_Type_: [VideoSelectorProgramId](aws-properties-medialive-channel-videoselectorprogramid.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoSelectorProgramId

VpcOutputSettings

All content copied from https://docs.aws.amazon.com/.
