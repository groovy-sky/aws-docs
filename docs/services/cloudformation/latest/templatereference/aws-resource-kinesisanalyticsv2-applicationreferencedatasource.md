This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource

Adds a reference data source to an existing SQL-based Kinesis Data Analytics application.

Kinesis Data Analytics reads reference data (that is, an Amazon S3 object) and creates an
in-application table within your application. In the request, you provide the source (S3
bucket name and object key name), name of the in-application table to create, and the
necessary mapping information that describes how data in an Amazon S3 object maps to columns
in the resulting in-application table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource",
  "Properties" : {
      "ApplicationName" : String,
      "ReferenceDataSource" : ReferenceDataSource
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource
Properties:
  ApplicationName: String
  ReferenceDataSource:
    ReferenceDataSource

```

## Properties

`ApplicationName`

The name of the application.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReferenceDataSource`

For a SQL-based Kinesis Data Analytics application, describes the reference data
source by providing the source information (Amazon S3 bucket name and object key name), the
resulting in-application table name that is created, and the necessary schema to map the data
elements in the Amazon S3 object to the in-application table.

_Required_: Yes

_Type_: [ReferenceDataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referencedatasource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

### Create an ApplicationReferenceDataSource resource

#### JSON

```json

{
    "ApplicationReferenceDataSource": {
        "Type": "AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource",
        "Properties": {
            "ApplicationName": {
                "Ref": "BasicApplication"
            },
            "ReferenceDataSource": {
                "TableName": "exampleTable",
                "ReferenceSchema": {
                    "RecordColumns": [
                        {
                            "Name": "example",
                            "SqlType": "VARCHAR(16)",
                            "Mapping": "$.example"
                        }
                    ],
                    "RecordFormat": {
                        "RecordFormatType": "JSON",
                        "MappingParameters": {
                            "JSONMappingParameters": {
                                "RecordRowPath": "$"
                            }
                        }
                    }
                },
                "S3ReferenceDataSource": {
                    "BucketARN": {
                        "Fn::GetAtt": [
                            "S3Bucket",
                            "Arn"
                        ]
                    },
                    "FileKey": "fakeKey"
                }
            }
        }
    }
}
```

#### YAML

```yaml

ApplicationReferenceDataSource:
  Type: AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource
  Properties:
    ApplicationName:
      Ref: BasicApplication
    ReferenceDataSource:
      TableName: exampleTable
      ReferenceSchema:
        RecordColumns:
        - Name: example
          SqlType: VARCHAR(16)
          Mapping: "$.example"
        RecordFormat:
          RecordFormatType: JSON
          MappingParameters:
            JSONMappingParameters:
              RecordRowPath: "$"
      S3ReferenceDataSource:
        BucketARN:
          Fn::GetAtt:
          - S3Bucket
          - Arn
        FileKey: fakeKey
```

## See also

- [AddApplicationReferenceDataSource](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_AddApplicationReferenceDataSource.html) in the _Amazon_
_Kinesis Data Analytics API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Output

CSVMappingParameters
