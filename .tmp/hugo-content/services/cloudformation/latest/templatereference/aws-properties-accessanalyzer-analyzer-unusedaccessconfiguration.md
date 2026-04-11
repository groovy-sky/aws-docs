This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer UnusedAccessConfiguration

Contains information about an unused access analyzer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnalysisRule" : AnalysisRule,
  "UnusedAccessAge" : Integer
}

```

### YAML

```yaml

  AnalysisRule:
    AnalysisRule
  UnusedAccessAge: Integer

```

## Properties

`AnalysisRule`

Contains information about analysis rules for the analyzer. Analysis rules determine
which entities will generate findings based on the criteria you define when you create the
rule.

_Required_: No

_Type_: [AnalysisRule](aws-properties-accessanalyzer-analyzer-analysisrule.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UnusedAccessAge`

The specified access age in days for which to generate findings for unused access. For
example, if you specify 90 days, the analyzer will generate findings for IAM entities
within the accounts of the selected organization for any access that hasn't been used in 90
or more days since the analyzer's last scan. You can choose a value between 1 and 365
days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `365`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
