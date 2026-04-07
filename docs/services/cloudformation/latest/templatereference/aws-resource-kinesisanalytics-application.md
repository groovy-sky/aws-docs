This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::Application

The `AWS::KinesisAnalytics::Application` resource creates an Amazon Kinesis Data Analytics application. For more information, see the [Amazon Kinesis Data Analytics Developer Guide](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/what-is.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalytics::Application",
  "Properties" : {
      "ApplicationCode" : String,
      "ApplicationDescription" : String,
      "ApplicationName" : String,
      "Inputs" : [ Input, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalytics::Application
Properties:
  ApplicationCode: String
  ApplicationDescription: String
  ApplicationName: String
  Inputs:
    - Input

```

## Properties

`ApplicationCode`

One or more SQL statements that read input data, transform it, and generate output.
For example, you can write a SQL statement that reads data from one in-application
stream, generates a running average of the number of advertisement clicks by vendor, and
insert resulting rows in another in-application stream using pumps. For more information
about the typical pattern, see [Application\
Code](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-app-code.html).

You can provide such series of SQL statements, where output of one statement can be
used as the input for the next statement. You store intermediate results by creating
in-application streams and pumps.

Note that the application code must create the streams with names specified in the
`Outputs`. For example, if your `Outputs` defines output
streams named `ExampleOutputStream1` and `ExampleOutputStream2`,
then your application code must create these streams.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `102400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationDescription`

Summary description of the application.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationName`

Name of your Amazon Kinesis Analytics application (for example,
`sample-app`).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Inputs`

Use this parameter to configure the application input.

You can configure your application to receive input from a single streaming source. In
this configuration, you map this streaming source to an in-application stream that is
created. Your application code can then query the in-application stream like a table
(you can think of it as a constantly updating table).

For the streaming source, you provide its Amazon Resource Name (ARN) and format of
data on the stream (for example, JSON, CSV, etc.). You also must provide an IAM role
that Amazon Kinesis Analytics can assume to read this stream on your behalf.

To create the in-application stream, you need to specify a schema to transform your
data into a schematized version used in SQL. In the schema, you provide the necessary
mapping of the data elements in the streaming source to record columns in the in-app
stream.

_Required_: Yes

_Type_: Array of [Input](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalytics-application-input.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

### Creating an Amazon Kinesis Data Analytics Application

The following example demonstrates how to create and configure a Kinesis Data Analytics application.

#### YAML

```yaml

---
Description: "Sample KinesisAnalytics via CloudFormation"
Resources:
  BasicApplication:
    Type: AWS::KinesisAnalytics::Application
    Properties:
      ApplicationName: "sampleApplication"
      ApplicationDescription: "SampleApp"
      ApplicationCode: "Example Application Code"
      Inputs:
        - NamePrefix: "exampleNamePrefix"
          InputSchema:
            RecordColumns:
             - Name: "example"
               SqlType: "VARCHAR(16)"
               Mapping: "$.example"
            RecordFormat:
              RecordFormatType: "JSON"
              MappingParameters:
                JSONMappingParameters:
                  RecordRowPath: "$"
          KinesisStreamsInput:
            ResourceARN: !GetAtt InputKinesisStream.Arn
            RoleARN: !GetAtt KinesisAnalyticsRole.Arn
  InputKinesisStream:
    Type: AWS::Kinesis::Stream
    Properties:
      ShardCount: 1
  KinesisAnalyticsRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: Allow
            Principal:
              Service: kinesisanalytics.amazonaws.com
            Action: "sts:AssumeRole"
      Path: "/"
      Policies:
        - PolicyName: Open
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Action: "*"
                Resource: "*"
  BasicApplicationOutputs:
    Type: AWS::KinesisAnalytics::ApplicationOutput
    DependsOn: BasicApplication
    Properties:
      ApplicationName: !Ref BasicApplication
      Output:
        Name: "exampleOutput"
        DestinationSchema:
          RecordFormatType: "CSV"
        KinesisStreamsOutput:
          ResourceARN: !GetAtt OutputKinesisStream.Arn
          RoleARN: !GetAtt KinesisAnalyticsRole.Arn
  OutputKinesisStream:
    Type: AWS::Kinesis::Stream
    Properties:
      ShardCount: 1
  ApplicationReferenceDataSource:
    Type: AWS::KinesisAnalytics::ApplicationReferenceDataSource
    DependsOn: BasicApplicationOutputs
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
  S3Bucket:
    Type: AWS::S3::Bucket
Outputs:
  ApplicationPhysicalResourceId:
    Value: !Ref BasicApplication
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Managed Service for Apache Flink

CSVMappingParameters
