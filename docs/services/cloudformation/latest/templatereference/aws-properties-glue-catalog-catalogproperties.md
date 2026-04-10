This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Catalog CatalogProperties

A structure that specifies data lake access properties and other custom properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomProperties" : {Key: Value, ...},
  "DataLakeAccessProperties" : DataLakeAccessProperties
}

```

### YAML

```yaml

  CustomProperties:
    Key: Value
  DataLakeAccessProperties:
    DataLakeAccessProperties

```

## Properties

`CustomProperties`

Additional key-value properties for the catalog, such as column statistics optimizations.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLakeAccessProperties`

A `DataLakeAccessProperties` object that specifies properties to configure data lake access for your catalog resource in the AWS Glue Data Catalog.

_Required_: No

_Type_: [DataLakeAccessProperties](aws-properties-glue-catalog-datalakeaccessproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::Catalog

DataLakeAccessProperties

All content copied from https://docs.aws.amazon.com/.
