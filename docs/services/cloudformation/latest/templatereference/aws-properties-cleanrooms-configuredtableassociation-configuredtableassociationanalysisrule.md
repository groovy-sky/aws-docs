This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTableAssociation ConfiguredTableAssociationAnalysisRule

An analysis rule for a configured table association. This analysis rule specifies how
data from the table can be used within its associated collaboration. In the console, the
`ConfiguredTableAssociationAnalysisRule` is referred to as the
_collaboration analysis rule_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Policy" : ConfiguredTableAssociationAnalysisRulePolicy,
  "Type" : String
}

```

### YAML

```yaml

  Policy:
    ConfiguredTableAssociationAnalysisRulePolicy
  Type: String

```

## Properties

`Policy`

The policy of the configured table association analysis rule.

_Required_: Yes

_Type_: [ConfiguredTableAssociationAnalysisRulePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-configuredtableassociation-configuredtableassociationanalysisrulepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the configured table association analysis rule.

_Required_: Yes

_Type_: String

_Allowed values_: `AGGREGATION | LIST | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CleanRooms::ConfiguredTableAssociation

ConfiguredTableAssociationAnalysisRuleAggregation
