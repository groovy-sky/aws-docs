This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Logging EventConfiguration

Configuration for event-based logging that specifies which event types to log and their logging settings. Used for account-level logging overrides.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventType" : String,
  "LogDestination" : String,
  "LogLevel" : String
}

```

### YAML

```yaml

  EventType: String
  LogDestination: String
  LogLevel: String

```

## Properties

`EventType`

The type of event to log. These include event types like Connect, Publish, and Disconnect.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogDestination`

CloudWatch Log Group for event-based logging. Specifies where log events should be sent. The log destination for event-based logging overrides default Log Group for the specified event type and applies to all resources associated with that event.

_Required_: No

_Type_: String

_Pattern_: `^(?!aws/)[a-zA-Z0-9_\-/.#]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

The logging level for the specified event type. Determines the verbosity of log messages generated for this event type. Valid Values: `ERROR | WARN | INFO | DEBUG | DISABLED`

_Required_: No

_Type_: String

_Allowed values_: `ERROR | WARN | INFO | DEBUG | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::Logging

AWS::IoT::MitigationAction

All content copied from https://docs.aws.amazon.com/.
