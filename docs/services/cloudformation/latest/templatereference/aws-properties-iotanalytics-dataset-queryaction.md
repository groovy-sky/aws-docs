This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset QueryAction

An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filters" : [ Filter, ... ],
  "SqlQuery" : String
}

```

### YAML

```yaml

  Filters:
    - Filter
  SqlQuery: String

```

## Properties

`Filters`

Pre-filters applied to message data.

_Required_: No

_Type_: Array of [Filter](aws-properties-iotanalytics-dataset-filter.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlQuery`

An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputFileUriValue

ResourceConfiguration

All content copied from https://docs.aws.amazon.com/.
