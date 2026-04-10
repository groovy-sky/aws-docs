This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationOutput

Adds an external destination to your SQL-based Amazon Kinesis Data Analytics
application.

If you want Kinesis Data Analytics to deliver data from an in-application stream
within your application to an external destination (such as an Kinesis data stream, a
Kinesis Data Firehose delivery stream, or an Amazon Lambda function), you add the
relevant configuration to your application using this operation. You can configure one
or more outputs for your application. Each output configuration maps an in-application
stream and an external destination.

You can use one of the output configurations to deliver data from your in-application
error stream to an external destination so that you can analyze the errors.

Any configuration update, including adding a streaming source using this operation,
results in a new version of the application. You can use the [DescribeApplication](../../../managed-flink/latest/apiv2/api-describeapplication.md) operation to find the current application
version.

###### Note

Creation of multiple outputs should be sequential (use of DependsOn) to avoid a
problem with a stale application version
( _ConcurrentModificationException_).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalyticsV2::ApplicationOutput",
  "Properties" : {
      "ApplicationName" : String,
      "Output" : Output
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalyticsV2::ApplicationOutput
Properties:
  ApplicationName: String
  Output:
    Output

```

## Properties

`ApplicationName`

The name of the application.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Output`

Describes a SQL-based Kinesis Data Analytics application's output configuration,
in which you identify an in-application stream and a destination where you want the
in-application stream data to be written. The destination can be a Kinesis data stream or a
Kinesis Data Firehose delivery stream.

_Required_: Yes

_Type_: [Output](aws-properties-kinesisanalyticsv2-applicationoutput-output.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

### Create an ApplicationOutput object

#### JSON

```json

{
    "Type": "AWS::KinesisAnalyticsV2::ApplicationOutput",
    "Properties": {
        "ApplicationName": {
            "Ref": "BasicApplication"
        },
        "Output": {
            "Name": "exampleOutput",
            "DestinationSchema": {
                "RecordFormatType": "CSV"
            },
            "KinesisStreamsOutput": {
                "ResourceARN": {
                    "Fn::GetAtt": [
                        "OutputKinesisStream",
                        "Arn"
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

Type: AWS::KinesisAnalyticsV2::ApplicationOutput
Properties:
  ApplicationName:
    Ref: BasicApplication
  Output:
    Name: exampleOutput
    DestinationSchema:
      RecordFormatType: CSV
    KinesisStreamsOutput:
      ResourceARN:
        Fn::GetAtt:
        - OutputKinesisStream
        - Arn

```

## See also

- [AddApplicationOutput](../../../managed-flink/latest/apiv2/api-addapplicationoutput.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLoggingOption

DestinationSchema

All content copied from https://docs.aws.amazon.com/.
