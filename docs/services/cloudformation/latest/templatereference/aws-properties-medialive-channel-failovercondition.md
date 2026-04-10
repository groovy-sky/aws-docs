This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FailoverCondition

Failover Condition settings. There can be multiple failover conditions inside
AutomaticInputFailoverSettings.

The parent of this entity is AutomaticInputFailoverSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailoverConditionSettings" : FailoverConditionSettings
}

```

### YAML

```yaml

  FailoverConditionSettings:
    FailoverConditionSettings

```

## Properties

`FailoverConditionSettings`

Settings for a specific failover condition.

_Required_: No

_Type_: [FailoverConditionSettings](aws-properties-medialive-channel-failoverconditionsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Esam

FailoverConditionSettings

All content copied from https://docs.aws.amazon.com/.
