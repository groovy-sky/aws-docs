This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore FileFormatConfiguration

Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://parquet.apache.org/).

The default file format is JSON. You can specify only one format.

You can't change the file format after you create the data store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonConfiguration" : Json,
  "ParquetConfiguration" : ParquetConfiguration
}

```

### YAML

```yaml

  JsonConfiguration:
    Json
  ParquetConfiguration:
    ParquetConfiguration

```

## Properties

`JsonConfiguration`

Contains the configuration information of the JSON format.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParquetConfiguration`

Contains the configuration information of the Parquet format.

_Required_: No

_Type_: [ParquetConfiguration](aws-properties-iotanalytics-datastore-parquetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatastoreStorage

IotSiteWiseMultiLayerStorage

All content copied from https://docs.aws.amazon.com/.
