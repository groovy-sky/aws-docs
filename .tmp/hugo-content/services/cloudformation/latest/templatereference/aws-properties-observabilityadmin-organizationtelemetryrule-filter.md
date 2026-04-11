This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule Filter

A single filter condition that specifies behavior, requirement, and matching conditions
for WAF log records.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Behavior" : String,
  "Conditions" : [ Condition, ... ],
  "Requirement" : String
}

```

### YAML

```yaml

  Behavior: String
  Conditions:
    - Condition
  Requirement: String

```

## Properties

`Behavior`

The action to take for log records matching this filter (KEEP or DROP).

_Required_: No

_Type_: String

_Allowed values_: `KEEP | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

The list of conditions that determine if a log record matches this filter.

_Required_: No

_Type_: Array of [Condition](aws-properties-observabilityadmin-organizationtelemetryrule-condition.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Requirement`

Whether the log record must meet all conditions (MEETS\_ALL) or any condition (MEETS\_ANY)
to match this filter.

_Required_: No

_Type_: String

_Allowed values_: `MEETS_ALL | MEETS_ANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldToMatch

LabelNameCondition

All content copied from https://docs.aws.amazon.com/.
