This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export

Creates a data export and specifies the data query, the delivery preference, and any
optional resource tags.

A `DataQuery` consists of both a `QueryStatement` and
`TableConfigurations`.

The `QueryStatement` is an SQL statement. Data Exports only supports a limited
subset of the SQL syntax. For more information on the SQL syntax that is supported, see [Data query](https://docs.aws.amazon.com/cur/latest/userguide/de-data-query.html). To
view the available tables and columns, see the [Data Exports table\
dictionary](https://docs.aws.amazon.com/cur/latest/userguide/de-table-dictionary.html).

The `TableConfigurations` is a collection of specified
`TableProperties` for the table being queried in the `QueryStatement`.
TableProperties are additional configurations you can provide to change the data and schema of
a table. Each table can have different TableProperties. However, tables are not required to
have any TableProperties. Each table property has a default value that it assumes if not
specified. For more information on table configurations, see [Data query](https://docs.aws.amazon.com/cur/latest/userguide/de-data-query.html). To
view the table properties available for each table, see the [Data Exports table\
dictionary](https://docs.aws.amazon.com/cur/latest/userguide/de-table-dictionary.html) or use the `ListTables` API to get a response of all tables
and their available properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BCMDataExports::Export",
  "Properties" : {
      "Export" : Export,
      "Tags" : [ ResourceTag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BCMDataExports::Export
Properties:
  Export:
    Export
  Tags:
    - ResourceTag

```

## Properties

`Export`

The details that are available for an export.

_Required_: Yes

_Type_: [Export](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bcmdataexports-export-export.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [ResourceTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bcmdataexports-export-resourcetag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Export.ExportArn`

The Amazon Resource Name (ARN) for this export.

`ExportArn`

The Amazon Resource Name (ARN) for this export.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Data Exports

DataQuery
