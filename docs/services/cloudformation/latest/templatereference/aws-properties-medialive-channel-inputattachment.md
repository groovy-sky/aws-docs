This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputAttachment

An input to attach to this channel.

This entity is at the top level in the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutomaticInputFailoverSettings" : AutomaticInputFailoverSettings,
  "InputAttachmentName" : String,
  "InputId" : String,
  "InputSettings" : InputSettings,
  "LogicalInterfaceNames" : [ String, ... ]
}

```

### YAML

```yaml

  AutomaticInputFailoverSettings:
    AutomaticInputFailoverSettings
  InputAttachmentName: String
  InputId: String
  InputSettings:
    InputSettings
  LogicalInterfaceNames:
    - String

```

## Properties

`AutomaticInputFailoverSettings`

Settings to implement automatic input failover in this input.

_Required_: No

_Type_: [AutomaticInputFailoverSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-automaticinputfailoversettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputAttachmentName`

A name for the attachment. This is required if you want to use this input in an input switch action.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputId`

The ID of the input to attach.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSettings`

Information about the content to extract from the input and about the general handling of the content.

_Required_: No

_Type_: [InputSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-channel-inputsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalInterfaceNames`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InferenceSettings

InputChannelLevel
