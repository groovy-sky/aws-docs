This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL AWSManagedRulesBotControlRuleSet

Details for your use of the Bot Control managed rule group, `AWSManagedRulesBotControlRuleSet`. This configuration is used in `ManagedRuleGroupConfig`.

For additional information about this and the other intelligent threat mitigation rule groups,
see [Intelligent threat mitigation in AWS WAF](../../../waf/latest/developerguide/waf-managed-protections.md)
and [AWS Managed Rules rule groups list](../../../waf/latest/developerguide/aws-managed-rule-groups-list.md)
in the _AWS WAF Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableMachineLearning" : Boolean,
  "InspectionLevel" : String
}

```

### YAML

```yaml

  EnableMachineLearning: Boolean
  InspectionLevel: String

```

## Properties

`EnableMachineLearning`

Applies only to the targeted inspection level.

Determines whether to use machine learning (ML) to
analyze your web traffic for bot-related activity. Machine learning is required for the Bot Control rules `TGT_ML_CoordinatedActivityLow` and `TGT_ML_CoordinatedActivityMedium`, which
inspect for anomalous behavior that might indicate distributed, coordinated bot activity.

For more information about this choice, see the listing for these rules in the table at [Bot Control rules listing](../../../waf/latest/developerguide/aws-managed-rule-groups-bot.md#aws-managed-rule-groups-bot-rules) in the
_AWS WAF Developer Guide_.

Default: `TRUE`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InspectionLevel`

The inspection level to use for the Bot Control rule group. The common level is the least expensive. The
targeted level includes all common level rules and adds rules with more advanced inspection criteria. For
details, see [AWS WAF Bot Control rule group](../../../waf/latest/developerguide/aws-managed-rule-groups-bot.md)
in the _AWS WAF Developer Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `COMMON | TARGETED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSManagedRulesATPRuleSet

BlockAction

All content copied from https://docs.aws.amazon.com/.
