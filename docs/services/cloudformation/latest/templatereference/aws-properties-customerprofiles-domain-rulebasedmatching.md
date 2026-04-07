This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain RuleBasedMatching

The process of matching duplicate profiles using Rule-Based matching. If
`RuleBasedMatching = true`, Amazon Connect Customer Profiles will
start to match and merge your profiles according to your configuration in the
`RuleBasedMatchingRequest`. You can use the
`ListRuleBasedMatches` and `GetSimilarProfiles` API to return
and review the results. Also, if you have configured `ExportingConfig` in the
`RuleBasedMatchingRequest`, you can download the results from S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeTypesSelector" : AttributeTypesSelector,
  "ConflictResolution" : ConflictResolution,
  "Enabled" : Boolean,
  "ExportingConfig" : ExportingConfig,
  "MatchingRules" : [ MatchingRule, ... ],
  "MaxAllowedRuleLevelForMatching" : Integer,
  "MaxAllowedRuleLevelForMerging" : Integer,
  "Status" : String
}

```

### YAML

```yaml

  AttributeTypesSelector:
    AttributeTypesSelector
  ConflictResolution:
    ConflictResolution
  Enabled: Boolean
  ExportingConfig:
    ExportingConfig
  MatchingRules:
    - MatchingRule
  MaxAllowedRuleLevelForMatching: Integer
  MaxAllowedRuleLevelForMerging: Integer
  Status: String

```

## Properties

`AttributeTypesSelector`

Configures information about the `AttributeTypesSelector` where the
rule-based identity resolution uses to match profiles.

_Required_: No

_Type_: [AttributeTypesSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-attributetypesselector.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConflictResolution`

Determines how the auto-merging process should resolve conflicts between different
profiles. For example, if Profile A and Profile B have the same `FirstName`
and `LastName`, `ConflictResolution` specifies which
`EmailAddress` should be used.

_Required_: No

_Type_: [ConflictResolution](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-conflictresolution.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

The flag that enables the matching process of duplicate profiles.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportingConfig`

The S3 location where Identity Resolution Jobs write result files.

_Required_: No

_Type_: [ExportingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-exportingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchingRules`

Configures how the rule-based matching process should match profiles. You can have up
to 15 `MatchingRule` in the `MatchingRules`.

_Required_: No

_Type_: Array of [MatchingRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-domain-matchingrule.html)

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAllowedRuleLevelForMatching`

Indicates the maximum allowed rule level for matching.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAllowedRuleLevelForMerging`

Indicates the maximum allowed rule level for merging.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of rule-based matching rule.

_Required_: No

_Type_: String

_Allowed values_: `PENDING | IN_PROGRESS | ACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Readiness

S3ExportingConfig
