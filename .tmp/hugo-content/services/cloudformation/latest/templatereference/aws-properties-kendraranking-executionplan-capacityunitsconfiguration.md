This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KendraRanking::ExecutionPlan CapacityUnitsConfiguration

Sets additional capacity units configured for your
rescore execution plan. A rescore execution plan is an
Amazon Kendra Intelligent Ranking resource used for
provisioning the `Rescore` API. You can add and
remove capacity units to fit your usage requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RescoreCapacityUnits" : Integer
}

```

### YAML

```yaml

  RescoreCapacityUnits: Integer

```

## Properties

`RescoreCapacityUnits`

The amount of extra capacity for your rescore execution
plan.

A single extra capacity unit for a rescore execution
plan provides 0.01 rescore requests per second. You can add
up to 1000 extra capacity units.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KendraRanking::ExecutionPlan

Tag

All content copied from https://docs.aws.amazon.com/.
