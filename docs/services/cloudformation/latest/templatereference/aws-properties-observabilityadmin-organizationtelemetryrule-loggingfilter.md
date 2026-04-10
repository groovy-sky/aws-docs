This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule LoggingFilter

Configuration that determines which WAF log records to keep or drop based on specified
conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultBehavior" : String,
  "Filters" : [ Filter, ... ]
}

```

### YAML

```yaml

  DefaultBehavior: String
  Filters:
    - Filter

```

## Properties

`DefaultBehavior`

The default action (KEEP or DROP) for log records that don't match any filter conditions.

_Required_: No

_Type_: String

_Allowed values_: `KEEP | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filters`

A list of filter conditions that determine log record handling behavior.

_Required_: No

_Type_: Array of [Filter](aws-properties-observabilityadmin-organizationtelemetryrule-filter.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelNameCondition

SingleHeader

All content copied from https://docs.aws.amazon.com/.
