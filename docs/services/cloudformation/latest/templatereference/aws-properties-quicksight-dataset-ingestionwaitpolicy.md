This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet IngestionWaitPolicy

The wait policy to use when creating or updating a Dataset. The default is to wait for SPICE ingestion to finish
with timeout of 36 hours.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IngestionWaitTimeInHours" : Number,
  "WaitForSpiceIngestion" : Boolean
}

```

### YAML

```yaml

  IngestionWaitTimeInHours: Number
  WaitForSpiceIngestion: Boolean

```

## Properties

`IngestionWaitTimeInHours`

The maximum time (in hours) to wait for Ingestion to complete. Default timeout is 36 hours. Applicable only when
`DataSetImportMode` mode is set to SPICE and `WaitForSpiceIngestion` is set to true.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitForSpiceIngestion`

Wait for SPICE ingestion to finish to mark dataset creation or update as successful. Default (true). Applicable
only when `DataSetImportMode` mode is set to SPICE.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncrementalRefresh

InputColumn

All content copied from https://docs.aws.amazon.com/.
