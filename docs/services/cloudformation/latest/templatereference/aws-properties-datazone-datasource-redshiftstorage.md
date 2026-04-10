This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource RedshiftStorage

The details of the Amazon Redshift storage as part of the configuration of an Amazon
Redshift data source run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RedshiftClusterSource" : RedshiftClusterStorage,
  "RedshiftServerlessSource" : RedshiftServerlessStorage
}

```

### YAML

```yaml

  RedshiftClusterSource:
    RedshiftClusterStorage
  RedshiftServerlessSource:
    RedshiftServerlessStorage

```

## Properties

`RedshiftClusterSource`

The details of the Amazon Redshift cluster source.

_Required_: No

_Type_: [RedshiftClusterStorage](aws-properties-datazone-datasource-redshiftclusterstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftServerlessSource`

The details of the Amazon Redshift Serverless workgroup source.

_Required_: No

_Type_: [RedshiftServerlessStorage](aws-properties-datazone-datasource-redshiftserverlessstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftServerlessStorage

RelationalFilterConfiguration

All content copied from https://docs.aws.amazon.com/.
