This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputLossFailoverSettings

MediaLive will perform a failover if content is not detected in this input for the specified period.

The parent of this entity is FailoverConditionSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputLossThresholdMsec" : Integer
}

```

### YAML

```yaml

  InputLossThresholdMsec: Integer

```

## Properties

`InputLossThresholdMsec`

The amount of time (in milliseconds) that no input is detected. After that time, an input failover will occur.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputLossBehavior

InputSettings

All content copied from https://docs.aws.amazon.com/.
