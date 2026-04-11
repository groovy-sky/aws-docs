This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AutomaticInputFailoverSettings

Settings to configure the conditions that will define the input as unhealthy and that will make MediaLive fail
over to the other input in the input failover pair.

The parent of this entity is InputAttachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorClearTimeMsec" : Integer,
  "FailoverConditions" : [ FailoverCondition, ... ],
  "InputPreference" : String,
  "SecondaryInputId" : String
}

```

### YAML

```yaml

  ErrorClearTimeMsec: Integer
  FailoverConditions:
    - FailoverCondition
  InputPreference: String
  SecondaryInputId: String

```

## Properties

`ErrorClearTimeMsec`

This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailoverConditions`

A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input.

_Required_: No

_Type_: Array of [FailoverCondition](aws-properties-medialive-channel-failovercondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputPreference`

Input preference when deciding which input to make active when a previously failed input has recovered.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryInputId`

The input ID of the secondary input in the automatic input failover pair.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioWatermarkSettings

Av1ColorSpaceSettings

All content copied from https://docs.aws.amazon.com/.
