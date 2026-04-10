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

_Type_: [Array](aws-properties-accessanalyzer-analyzer-filter.md) of [Filter](aws-properties-accessanalyzer-analyzer-filter.md)

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalyzerConfiguration

Filter

All content copied from https://docs.aws.amazon.com/.
