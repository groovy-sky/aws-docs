This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application WindowsEvent

The `AWS::ApplicationInsights::Application WindowsEvent` property type specifies a Windows Event to monitor for the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventLevels" : [ String, ... ],
  "EventName" : String,
  "LogGroupName" : String,
  "PatternSet" : String
}

```

### YAML

```yaml

  EventLevels:
    - String
  EventName: String
  LogGroupName: String
  PatternSet: String

```

## Properties

`EventLevels`

The levels of event to log. You must specify each level to log. Possible values include `INFORMATION`, `WARNING`, `ERROR`, `CRITICAL`, and `VERBOSE`. This field is required for each type of Windows Event to log.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventName`

The type of Windows Events to log, equivalent to the Windows Event log channel name. For
example, System, Security, CustomEventName, and so on. This field is required for each
type of Windows event to log.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_ \\/-]+$`

_Minimum_: `1`

_Maximum_: `260`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The CloudWatch log group name to be associated with the monitored log.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternSet`

The log pattern set.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9.-_]*`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
