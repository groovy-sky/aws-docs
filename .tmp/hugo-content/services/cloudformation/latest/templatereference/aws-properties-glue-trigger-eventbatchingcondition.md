This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Trigger EventBatchingCondition

Batch condition that must be met (specified number of events received or batch time window expired) before EventBridge event trigger fires.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchSize" : Integer,
  "BatchWindow" : Integer
}

```

### YAML

```yaml

  BatchSize: Integer
  BatchWindow: Integer

```

## Properties

`BatchSize`

Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BatchWindow`

Window of time in seconds after which EventBridge event trigger fires. Window starts when first event is received.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

NotificationProperty

All content copied from https://docs.aws.amazon.com/.
