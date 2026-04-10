This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AnalysisRule

A specification about how data from the configured table can be used in a query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Policy" : ConfiguredTableAnalysisRulePolicy,
  "Type" : String
}

```

### YAML

```yaml

  Policy:
    ConfiguredTableAnalysisRulePolicy
  Type: String

```

## Properties

`Policy`

A policy that describes the associated data usage limitations.

_Required_: Yes

_Type_: [ConfiguredTableAnalysisRulePolicy](aws-properties-cleanrooms-configuredtable-configuredtableanalysisrulepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of analysis rule.

_Required_: Yes

_Type_: String

_Allowed values_: `AGGREGATION | LIST | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AggregationConstraint

AnalysisRuleAggregation

All content copied from https://docs.aws.amazon.com/.
