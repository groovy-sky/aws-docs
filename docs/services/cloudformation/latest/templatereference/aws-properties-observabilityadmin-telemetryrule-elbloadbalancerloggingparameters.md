This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule ELBLoadBalancerLoggingParameters

Configuration parameters for ELB load balancer logging, including output format and field
delimiter settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldDelimiter" : String,
  "OutputFormat" : String
}

```

### YAML

```yaml

  FieldDelimiter: String
  OutputFormat: String

```

## Properties

`FieldDelimiter`

The delimiter character used to separate fields in ELB access log entries when using
plain text format.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFormat`

The format for ELB access log entries (plain text or JSON format).

_Required_: No

_Type_: String

_Allowed values_: `plain | json`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

FieldToMatch

All content copied from https://docs.aws.amazon.com/.
