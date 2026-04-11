This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::NotificationChannel NotificationFilterConfig

The filter configurations for the Amazon SNS notification topic you use with DevOps Guru. You can choose to specify which events or message types to receive notifications for.
You can also choose to specify which severity levels to receive notifications for.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MessageTypes" : [ String, ... ],
  "Severities" : [ String, ... ]
}

```

### YAML

```yaml

  MessageTypes:
    - String
  Severities:
    - String

```

## Properties

`MessageTypes`

The events that you want to receive notifications for. For example, you can choose to receive notifications only when the severity level is upgraded or a new insight is created.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Severities`

The severity levels that you want to receive notifications for. For example, you can choose to receive notifications only for insights with `HIGH` and `MEDIUM` severity levels.
For more information, see [Understanding insight severities](../../../devops-guru/latest/userguide/working-with-insights.md#understanding-insights-severities).

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationChannelConfig

SnsChannelConfig

All content copied from https://docs.aws.amazon.com/.
