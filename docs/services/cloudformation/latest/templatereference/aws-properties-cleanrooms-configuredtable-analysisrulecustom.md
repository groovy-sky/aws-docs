This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AnalysisRuleCustom

A type of analysis rule that enables the table owner to approve custom SQL queries on
their configured tables. It supports differential privacy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalAnalyses" : String,
  "AllowedAnalyses" : [ String, ... ],
  "AllowedAnalysisProviders" : [ String, ... ],
  "DifferentialPrivacy" : DifferentialPrivacy,
  "DisallowedOutputColumns" : [ String, ... ]
}

```

### YAML

```yaml

  AdditionalAnalyses: String
  AllowedAnalyses:
    - String
  AllowedAnalysisProviders:
    - String
  DifferentialPrivacy:
    DifferentialPrivacy
  DisallowedOutputColumns:
    - String

```

## Properties

`AdditionalAnalyses`

An indicator as to whether additional analyses (such as AWS Clean Rooms ML) can be applied to the output of the direct query.

_Required_: No

_Type_: String

_Allowed values_: `ALLOWED | REQUIRED | NOT_ALLOWED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedAnalyses`

The ARN of the analysis templates that are allowed by the custom analysis rule.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedAnalysisProviders`

The IDs of the AWS accounts that are allowed to query by the custom analysis rule. Required when
`allowedAnalyses` is `ANY_QUERY`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DifferentialPrivacy`

The differential privacy configuration.

_Required_: No

_Type_: [DifferentialPrivacy](aws-properties-cleanrooms-configuredtable-differentialprivacy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisallowedOutputColumns`

A list of columns that aren't allowed to be shown in the query output.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisRuleAggregation

AnalysisRuleList

All content copied from https://docs.aws.amazon.com/.
