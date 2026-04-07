This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace WorkspaceConfiguration

Use this structure to define label sets and the ingestion limits for time series that
match label sets, and to specify the retention period of the workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LimitsPerLabelSets" : [ LimitsPerLabelSet, ... ],
  "RetentionPeriodInDays" : Integer
}

```

### YAML

```yaml

  LimitsPerLabelSets:
    - LimitsPerLabelSet
  RetentionPeriodInDays: Integer

```

## Properties

`LimitsPerLabelSets`

This is an array of structures, where each structure defines a label set for the
workspace, and defines the ingestion limit for active time series for each of those
label sets. Each label name in a label set must be unique.

_Required_: No

_Type_: Array of [LimitsPerLabelSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-workspace-limitsperlabelset.html)

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriodInDays`

Specifies how many days that metrics will be retained in the workspace.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
