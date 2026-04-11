This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer InternalAccessAnalysisRuleCriteria

The criteria for an analysis rule for an internal access analyzer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountIds" : [ String, ... ],
  "ResourceArns" : [ String, ... ],
  "ResourceTypes" : [ String, ... ]
}

```

### YAML

```yaml

  AccountIds:
    - String
  ResourceArns:
    - String
  ResourceTypes:
    - String

```

## Properties

`AccountIds`

A list of AWS account IDs to apply to the internal access analysis rule
criteria. Account IDs can only be applied to the analysis rule criteria for
organization-level analyzers.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ResourceArns`

A list of resource ARNs to apply to the internal access analysis rule criteria. The
analyzer will only generate findings for resources that match these ARNs.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ResourceTypes`

A list of resource types to apply to the internal access analysis rule criteria. The
analyzer will only generate findings for resources of these types. These resource types
are currently supported for internal access analyzers:

- `AWS::S3::Bucket`

- `AWS::RDS::DBSnapshot`

- `AWS::RDS::DBClusterSnapshot`

- `AWS::S3Express::DirectoryBucket`

- `AWS::DynamoDB::Table`

- `AWS::DynamoDB::Stream`

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InternalAccessAnalysisRule

InternalAccessConfiguration

All content copied from https://docs.aws.amazon.com/.
