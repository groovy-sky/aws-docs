This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer AnalysisRule

Contains information about analysis rules for the analyzer. Analysis rules determine
which entities will generate findings based on the criteria you define when you create the
rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exclusions" : [ AnalysisRuleCriteria, ... ]
}

```

### YAML

```yaml

  Exclusions:
    - AnalysisRuleCriteria

```

## Properties

`Exclusions`

A list of rules for the analyzer containing criteria to exclude from analysis. Entities
that meet the rule criteria will not generate findings.

_Required_: No

_Type_: Array of [AnalysisRuleCriteria](aws-properties-accessanalyzer-analyzer-analysisrulecriteria.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AccessAnalyzer::Analyzer

AnalysisRuleCriteria

All content copied from https://docs.aws.amazon.com/.
