This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FeatureActivations

Settings to enable specific features. You can't configure these features until you have enabled them in the
channel.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputPrepareScheduleActions" : String,
  "OutputStaticImageOverlayScheduleActions" : String
}

```

### YAML

```yaml

  InputPrepareScheduleActions: String
  OutputStaticImageOverlayScheduleActions: String

```

## Properties

`InputPrepareScheduleActions`

Enables the Input Prepare feature. You can create Input Prepare actions in the schedule only if this feature is enabled.
If you disable the feature on an existing schedule, make sure that you first delete all input prepare actions from the schedule.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputStaticImageOverlayScheduleActions`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailoverConditionSettings

FecOutputSettings

All content copied from https://docs.aws.amazon.com/.
