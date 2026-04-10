This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup TriggerConfig

Information about notification triggers for the deployment group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TriggerEvents" : [ String, ... ],
  "TriggerName" : String,
  "TriggerTargetArn" : String
}

```

### YAML

```yaml

  TriggerEvents:
    - String
  TriggerName: String
  TriggerTargetArn: String

```

## Properties

`TriggerEvents`

The event type or types that trigger notifications.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerName`

The name of the notification trigger.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerTargetArn`

The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through
which notifications about deployment or instance events are sent.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrafficRoute

Next

All content copied from https://docs.aws.amazon.com/.
