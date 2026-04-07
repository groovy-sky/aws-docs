This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer ArchiveRule

Contains information about an archive rule. Archive rules automatically archive new
findings that meet the criteria you define when you create the rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filter" : [ Filter, ... ],
  "RuleName" : String
}

```

### YAML

```yaml

  Filter:
    - Filter
  RuleName: String

```

## Properties

`Filter`

The criteria for the rule.

_Required_: Yes

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-accessanalyzer-analyzer-filter.html) of [Filter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-accessanalyzer-analyzer-filter.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the rule to create.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z][A-Za-z0-9_.-]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalyzerConfiguration

Filter
