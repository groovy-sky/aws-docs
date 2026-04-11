This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::EventBus LogConfig

The logging configuration settings for the event bus.

For more information, see [Configuring logs for event buses](../../../eventbridge/latest/userguide/eb-event-bus-logs.md) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncludeDetail" : String,
  "Level" : String
}

```

### YAML

```yaml

  IncludeDetail: String
  Level: String

```

## Properties

`IncludeDetail`

Whether EventBridge include detailed event information in the records it generates.
Detailed data can be useful for troubleshooting and debugging. This information includes details of the event itself, as well as target details.

For more information, see [Including detail data in event bus logs](../../../eventbridge/latest/userguide/eb-event-bus-logs.md#eb-event-logs-data) in the _EventBridge User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `FULL | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Level`

The level of logging detail to include. This applies to all log destinations for the event bus.

For more information, see [Specifying event bus log level](../../../eventbridge/latest/userguide/eb-event-bus-logs.md#eb-event-bus-logs-level) in the _EventBridge User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `INFO | ERROR | TRACE | OFF`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeadLetterConfig

Tag

All content copied from https://docs.aws.amazon.com/.
