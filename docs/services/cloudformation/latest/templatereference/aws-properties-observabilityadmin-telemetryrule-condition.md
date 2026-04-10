This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule Condition

A single condition that can match based on WAF rule action or label name.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionCondition" : ActionCondition,
  "LabelNameCondition" : LabelNameCondition
}

```

### YAML

```yaml

  ActionCondition:
    ActionCondition
  LabelNameCondition:
    LabelNameCondition

```

## Properties

`ActionCondition`

Matches log records based on the WAF rule action taken (ALLOW, BLOCK, COUNT, etc.).

_Required_: No

_Type_: [ActionCondition](aws-properties-observabilityadmin-telemetryrule-actioncondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelNameCondition`

Matches log records based on WAF rule labels applied to the request.

_Required_: No

_Type_: [LabelNameCondition](aws-properties-observabilityadmin-telemetryrule-labelnamecondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudtrailParameters

ELBLoadBalancerLoggingParameters

All content copied from https://docs.aws.amazon.com/.
