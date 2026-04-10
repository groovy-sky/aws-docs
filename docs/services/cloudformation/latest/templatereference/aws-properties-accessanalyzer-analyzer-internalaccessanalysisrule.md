This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer InternalAccessAnalysisRule

Contains information about analysis rules for the internal access analyzer. Analysis
rules determine which entities will generate findings based on the criteria you define
when you create the rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Inclusions" : [ InternalAccessAnalysisRuleCriteria, ... ]
}

```

### YAML

```yaml

  Inclusions:
    - InternalAccessAnalysisRuleCriteria

```

## Properties

`Inclusions`

A list of rules for the internal access analyzer containing criteria to include in
analysis. Only resources that meet the rule criteria will generate findings.

_Required_: No

_Type_: Array of [InternalAccessAnalysisRuleCriteria](aws-properties-accessanalyzer-analyzer-internalaccessanalysisrulecriteria.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

InternalAccessAnalysisRuleCriteria

All content copied from https://docs.aws.amazon.com/.
