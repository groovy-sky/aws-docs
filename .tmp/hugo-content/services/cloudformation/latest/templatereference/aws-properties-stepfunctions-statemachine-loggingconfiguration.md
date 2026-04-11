This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachine LoggingConfiguration

Defines what execution history events are logged and where they are logged.

Step Functions provides the log levels — `OFF`, `ALL`, `ERROR`, and `FATAL`. No event types log when set to `OFF` and all event types do when set to `ALL`.

###### Note

By default, the `level` is set to `OFF`. For more information see
[Log\
Levels](../../../step-functions/latest/dg/cloudwatch-log-level.md) in the AWS Step Functions User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destinations" : [ LogDestination, ... ],
  "IncludeExecutionData" : Boolean,
  "Level" : String
}

```

### YAML

```yaml

  Destinations:
    - LogDestination
  IncludeExecutionData: Boolean
  Level: String

```

## Properties

`Destinations`

An array of objects that describes where your execution history events will be logged.
Limited to size 1. Required, if your log level is not set to `OFF`.

_Required_: No

_Type_: Array of [LogDestination](aws-properties-stepfunctions-statemachine-logdestination.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeExecutionData`

Determines whether execution data is included in your log. When set to `false`,
data is excluded.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Level`

Defines which category of execution history events are logged.

_Required_: No

_Type_: String

_Allowed values_: `ALL | ERROR | FATAL | OFF`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogDestination

S3Location

All content copied from https://docs.aws.amazon.com/.
