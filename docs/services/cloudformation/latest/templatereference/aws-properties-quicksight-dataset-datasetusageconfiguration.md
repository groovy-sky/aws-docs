This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetUsageConfiguration

The usage configuration to apply to child datasets that reference this dataset as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisableUseAsDirectQuerySource" : Boolean,
  "DisableUseAsImportedSource" : Boolean
}

```

### YAML

```yaml

  DisableUseAsDirectQuerySource: Boolean
  DisableUseAsImportedSource: Boolean

```

## Properties

`DisableUseAsDirectQuerySource`

An option that controls whether a child dataset of a direct query can use this dataset as a source.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableUseAsImportedSource`

An option that controls whether a child dataset that's stored in Quick Sight can use this dataset as a source.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetStringListFilterValue

DateTimeDatasetParameter

All content copied from https://docs.aws.amazon.com/.
