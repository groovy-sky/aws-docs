This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationReferenceDataSource

Adds a reference data source to an existing application.

Amazon Kinesis Analytics reads reference data (that is, an Amazon S3 object) and creates an in-application table within your application. In the request, you provide the source (S3 bucket name and object key name), name of the in-application table to create, and the necessary mapping information that describes how data in Amazon S3 object maps to columns in the resulting in-application table.

For conceptual information,
see [Configuring Application Input](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
For the limits on data sources you can add to your application, see
[Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html).

This operation requires permissions to perform the `kinesisanalytics:AddApplicationOutput` action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalytics::ApplicationReferenceDataSource",
  "Properties" : {
      "ApplicationName" : String,
      "ReferenceDataSource" : ReferenceDataSource
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalytics::ApplicationReferenceDataSource
Properties:
  ApplicationName: String
  ReferenceDataSource:
    ReferenceDataSource

```

## Properties

`ApplicationName`

Name of an existing application.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReferenceDataSource`

The reference data source can be an object in your Amazon S3 bucket. Amazon Kinesis
Analytics reads the object and copies the data into the in-application table that is
created. You provide an S3 bucket, object key name, and the resulting in-application
table that is created. You must also provide an IAM role with the necessary permissions
that Amazon Kinesis Analytics can assume to read the object from your S3 bucket on your
behalf.

_Required_: Yes

_Type_: [ReferenceDataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

### Adding an ApplicationReferenceDataSource Resource

The following example creates an `ApplicationReferenceDataSource` resource:

#### YAML

```yaml

ApplicationReferenceDataSource:
    Type: AWS::KinesisAnalytics::ApplicationReferenceDataSource
    Properties:
      ApplicationName: !Ref BasicApplication
      ReferenceDataSource:
        TableName: "exampleTable"
        ReferenceSchema:
          RecordColumns:
            - Name: "example"
              SqlType: "VARCHAR(16)"
              Mapping: "$.example"
          RecordFormat:
            RecordFormatType: "JSON"
            MappingParameters:
              JSONMappingParameters:
                RecordRowPath: "$"
        S3ReferenceDataSource:
          BucketARN: !GetAtt S3Bucket.Arn
          FileKey: 'fakeKey'
          ReferenceRoleARN: !GetAtt KinesisAnalyticsRole.Arn
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Output

CSVMappingParameters
