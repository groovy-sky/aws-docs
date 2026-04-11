This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer InternalAccessConfiguration

Specifies the configuration of an internal access analyzer for an AWS
organization or account. This configuration determines how the analyzer evaluates
internal access within your AWS environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InternalAccessAnalysisRule" : InternalAccessAnalysisRule
}

```

### YAML

```yaml

  InternalAccessAnalysisRule:
    InternalAccessAnalysisRule

```

## Properties

`InternalAccessAnalysisRule`

Contains information about analysis rules for the internal access analyzer. These
rules determine which resources and access patterns will be analyzed.

_Required_: No

_Type_: [InternalAccessAnalysisRule](aws-properties-accessanalyzer-analyzer-internalaccessanalysisrule.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InternalAccessAnalysisRuleCriteria

Tag

All content copied from https://docs.aws.amazon.com/.
