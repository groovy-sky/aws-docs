This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::Group InsightsConfiguration

The structure containing configurations related to insights.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InsightsEnabled" : Boolean,
  "NotificationsEnabled" : Boolean
}

```

### YAML

```yaml

  InsightsEnabled: Boolean
  NotificationsEnabled: Boolean

```

## Properties

`InsightsEnabled`

Set the InsightsEnabled value to true to enable insights or false to disable
insights.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationsEnabled`

Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be
enabled on a group with InsightsEnabled set to true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::XRay::Group

Tag
