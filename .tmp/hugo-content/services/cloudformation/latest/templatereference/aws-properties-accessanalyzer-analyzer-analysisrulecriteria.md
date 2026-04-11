This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer AnalysisRuleCriteria

The criteria for an analysis rule for an analyzer. The criteria determine which entities
will generate findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountIds" : [ String, ... ],
  "ResourceTags" : [ [ , ... ], ... ]
}

```

### YAML

```yaml

  AccountIds:
    - String
  ResourceTags:
    -
    -

```

## Properties

`AccountIds`

A list of AWS account IDs to apply to the analysis rule criteria. The accounts cannot
include the organization analyzer owner account. Account IDs can only be applied to the
analysis rule criteria for organization-level analyzers. The list cannot include more than
2,000 account IDs.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ResourceTags`

An array of key-value pairs to match for your resources. You can use the set of Unicode
letters, digits, whitespace, `_`, `.`, `/`,
`=`, `+`, and `-`.

For the tag key, you can specify a value that is 1 to 128 characters in length and
cannot be prefixed with `aws:`.

For the tag value, you can specify a value that is 0 to 256 characters in length. If the
specified tag value is 0 characters, the rule is applied to all principals with the
specified tag key.

_Required_: No

_Type_: Array of Array

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisRule

AnalyzerConfiguration

All content copied from https://docs.aws.amazon.com/.
