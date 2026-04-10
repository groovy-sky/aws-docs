This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTableAssociation ConfiguredTableAssociationAnalysisRuleAggregation

The configured table association analysis rule applied to a configured table with the aggregation analysis rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedAdditionalAnalyses" : [ String, ... ],
  "AllowedResultReceivers" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedAdditionalAnalyses:
    - String
  AllowedResultReceivers:
    - String

```

## Properties

`AllowedAdditionalAnalyses`

The list of resources or wildcards (ARNs) that are allowed to perform additional analysis on query output.

The `allowedAdditionalAnalyses` parameter is currently supported for the list
analysis rule ( `AnalysisRuleList`) and the custom analysis rule
( `AnalysisRuleCustom`).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedResultReceivers`

The list of collaboration members who are allowed to receive results of queries run
with this configured table.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfiguredTableAssociationAnalysisRule

ConfiguredTableAssociationAnalysisRuleCustom

All content copied from https://docs.aws.amazon.com/.
