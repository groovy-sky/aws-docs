This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow MetadataCatalogConfig

Specifies the configuration that Amazon AppFlow uses when it catalogs your data. When
Amazon AppFlow catalogs your data, it stores metadata in a data catalog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlueDataCatalog" : GlueDataCatalog
}

```

### YAML

```yaml

  GlueDataCatalog:
    GlueDataCatalog

```

## Properties

`GlueDataCatalog`

Specifies the configuration that Amazon AppFlow uses when it catalogs your data with
the AWS Glue Data Catalog.

_Required_: No

_Type_: [GlueDataCatalog](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-gluedatacatalog.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MarketoSourceProperties

PardotSourceProperties
