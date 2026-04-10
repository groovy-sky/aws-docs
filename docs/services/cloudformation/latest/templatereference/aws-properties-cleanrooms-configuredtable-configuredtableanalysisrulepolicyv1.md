This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable ConfiguredTableAnalysisRulePolicyV1

Controls on the query specifications that can be run on a configured table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : AnalysisRuleAggregation,
  "Custom" : AnalysisRuleCustom,
  "List" : AnalysisRuleList
}

```

### YAML

```yaml

  Aggregation:
    AnalysisRuleAggregation
  Custom:
    AnalysisRuleCustom
  List:
    AnalysisRuleList

```

## Properties

`Aggregation`

Analysis rule type that enables only aggregation queries on a configured table.

_Required_: No

_Type_: [AnalysisRuleAggregation](aws-properties-cleanrooms-configuredtable-analysisruleaggregation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Custom`

Analysis rule type that enables custom SQL queries on a configured table.

_Required_: No

_Type_: [AnalysisRuleCustom](aws-properties-cleanrooms-configuredtable-analysisrulecustom.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`List`

Analysis rule type that enables only list queries on a configured table.

_Required_: No

_Type_: [AnalysisRuleList](aws-properties-cleanrooms-configuredtable-analysisrulelist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfiguredTableAnalysisRulePolicy

DifferentialPrivacy

All content copied from https://docs.aws.amazon.com/.
