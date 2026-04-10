This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace Label

A label is a name:value pair used to add context to ingested metrics. This structure
defines the name and value for one label that is used in a label set. You can set
ingestion limits on time series that match defined label sets, to help prevent a
workspace from being overwhelmed with unexpected spikes in time series ingestion.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name for this label.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for this label.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogDestination

LimitsPerLabelSet

All content copied from https://docs.aws.amazon.com/.
