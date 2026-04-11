This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace LimitsPerLabelSet

This defines a label set for the workspace, and defines the ingestion limit for active
time series that match that label set. Each label name in a label set must be
unique.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelSet" : [ Label, ... ],
  "Limits" : LimitsPerLabelSetEntry
}

```

### YAML

```yaml

  LabelSet:
    - Label
  Limits:
    LimitsPerLabelSetEntry

```

## Properties

`LabelSet`

This defines one label set that will have an enforced ingestion limit. You can set
ingestion limits on time series that match defined label sets, to help prevent a
workspace from being overwhelmed with unexpected spikes in time series ingestion.

Label values accept all UTF-8 characters with one exception. If the label name is
metric name label `__name__`, then the
_metric_ part of the name must conform to the following pattern:
`[a-zA-Z_:][a-zA-Z0-9_:]*`

_Required_: Yes

_Type_: Array of [Label](aws-properties-aps-workspace-label.md)

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Limits`

This structure contains the information about the limits that apply to time series
that match this label set.

_Required_: Yes

_Type_: [LimitsPerLabelSetEntry](aws-properties-aps-workspace-limitsperlabelsetentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Label

LimitsPerLabelSetEntry

All content copied from https://docs.aws.amazon.com/.
