This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule LogDeliveryParameters

Configuration parameters for Amazon Bedrock AgentCore logging, including `logType`
settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogTypes" : [ String, ... ]
}

```

### YAML

```yaml

  LogTypes:
    - String

```

## Properties

`LogTypes`

The type of log that the source is sending.

_Required_: No

_Type_: Array of String

_Allowed values_: `APPLICATION_LOGS | USAGE_LOGS`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelNameCondition

LoggingFilter

All content copied from https://docs.aws.amazon.com/.
