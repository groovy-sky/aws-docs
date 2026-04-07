This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore

AWS::IoTAnalytics::Datastore resource is a repository for messages. For more information, see
[How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how) in the _AWS IoT Analytics User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTAnalytics::Datastore",
  "Properties" : {
      "DatastoreName" : String,
      "DatastorePartitions" : DatastorePartitions,
      "DatastoreStorage" : DatastoreStorage,
      "FileFormatConfiguration" : FileFormatConfiguration,
      "RetentionPeriod" : RetentionPeriod,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTAnalytics::Datastore
Properties:
  DatastoreName: String
  DatastorePartitions:
    DatastorePartitions
  DatastoreStorage:
    DatastoreStorage
  FileFormatConfiguration:
    FileFormatConfiguration
  RetentionPeriod:
    RetentionPeriod
  Tags:
    - Tag

```

## Properties

`DatastoreName`

The name of the data store.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatastorePartitions`

Information about the partition dimensions in a data store.

_Required_: No

_Type_: [DatastorePartitions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-datastore-datastorepartitions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatastoreStorage`

Where data store data is stored.

_Required_: No

_Type_: [DatastoreStorage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-datastore-datastorestorage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileFormatConfiguration`

Contains the configuration information of file formats. AWS IoT Analytics data stores support JSON and [Parquet](https://parquet.apache.org/).

The default file format is JSON. You can specify only one format.

You can't change the file format after you create the data store.

_Required_: No

_Type_: [FileFormatConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-datastore-fileformatconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriod`

How long, in days, message data is kept for the data store. When
`customerManagedS3` storage is selected, this parameter is ignored.

_Required_: No

_Type_: [RetentionPeriod](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-datastore-retentionperiod.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the data store.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-datastore-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Simple Datastore](#aws-resource-iotanalytics-datastore--examples--Simple_Datastore)

- [Complex Datastore](#aws-resource-iotanalytics-datastore--examples--Complex_Datastore)

### Simple Datastore

The following example creates a simple datastore.

#### JSON

```json

{
    "Description": "Create a simple Datastore",
    "Resources": {
        "Datastore": {
            "Type": "AWS::IoTAnalytics::Datastore",
            "Properties": {
                "DatastoreName": "SimpleDatastore"
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a simple Datastore"
Resources:
  Datastore:
    Type: "AWS::IoTAnalytics::Datastore"
    Properties:
      DatastoreName: "SimpleDatastore"
```

### Complex Datastore

The following example creates a complex datastore.

#### JSON

```json

{
    "Description": "Create a complex Datastore",
    "Resources": {
        "Datastore": {
            "Type": "AWS::IoTAnalytics::Datastore",
            "Properties": {
                "DatastoreName": "ComplexDatastore",
                "RetentionPeriod": {
                    "Unlimited": false,
                    "NumberOfDays": 10
                },
                "Tags": [
                    {
                        "Key": "keyname1",
                        "Value": "value1"
                    },
                    {
                        "Key": "keyname2",
                        "Value": "value2"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Description: "Create a complex Datastore"
Resources:
  Datastore:
    Type: "AWS::IoTAnalytics::Datastore"
    Properties:
      DatastoreName: "ComplexDatastore"
      RetentionPeriod:
        Unlimited: false
        NumberOfDays: 10
      Tags:
        -
          Key: "keyname1"
          Value: "value1"
        -
          Key: "keyname2"
          Value: "value2"
```

## See also

- [How to Use AWS IoT Analytics](https://docs.aws.amazon.com/iotanalytics/latest/userguide/welcome.html#aws-iot-analytics-how)
in the _AWS IoT Analytics User Guide_

- [CreateDatastore](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_CreateDatastore.html) in the
_AWS IoT Analytics API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VersioningConfiguration

Column
