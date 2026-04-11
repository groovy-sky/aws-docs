This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTableAssociation ConfiguredTableAssociationAnalysisRulePolicyV1

Controls on the query specifications that can be run on an associated configured table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : ConfiguredTableAssociationAnalysisRuleAggregation,
  "Custom" : ConfiguredTableAssociationAnalysisRuleCustom,
  "List" : ConfiguredTableAssociationAnalysisRuleList
}

```

### YAML

```yaml

  Aggregation:
    ConfiguredTableAssociationAnalysisRuleAggregation
  Custom:
    ConfiguredTableAssociationAnalysisRuleCustom
  List:
    ConfiguredTableAssociationAnalysisRuleList

```

## Properties

`Aggregation`

Analysis rule type that enables only aggregation queries on a configured table.

_Required_: No

_Type_: [ConfiguredTableAssociationAnalysisRuleAggregation](aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisruleaggregation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Custom`

Analysis rule type that enables the table owner to approve custom SQL queries on their configured tables. It supports differential privacy.

_Required_: No

_Type_: [ConfiguredTableAssociationAnalysisRuleCustom](aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulecustom.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`List`

Analysis rule type that enables only list queries on a configured table.

_Required_: No

_Type_: [ConfiguredTableAssociationAnalysisRuleList](aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulelist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfiguredTableAssociationAnalysisRulePolicy

Tag

All content copied from https://docs.aws.amazon.com/.
