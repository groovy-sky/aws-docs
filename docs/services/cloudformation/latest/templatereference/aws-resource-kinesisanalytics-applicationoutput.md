This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationOutput

Adds an external destination to your Amazon Kinesis Analytics application.

If you want Amazon Kinesis Analytics to deliver data from an in-application stream
within your application to an external destination (such as an Amazon Kinesis stream, an
Amazon Kinesis Firehose delivery stream, or an Amazon Lambda function), you add the
relevant configuration to your application using this operation. You can configure one
or more outputs for your application. Each output configuration maps an in-application
stream and an external destination.

You can use one of the output configurations to deliver data from your
in-application error stream to an external destination so that you can analyze the
errors. For more information, see [Understanding Application\
Output (Destination)](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).

Any configuration update, including adding a streaming source using this
operation, results in a new version of the application. You can use the `DescribeApplication` operation to find the current application
version.

For the limits on the number of application inputs and outputs
you can configure, see [Limits](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html).

This operation requires permissions to perform the `kinesisanalytics:AddApplicationOutput` action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalytics::ApplicationOutput",
  "Properties" : {
      "ApplicationName" : String,
      "Output" : Output
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalytics::ApplicationOutput
Properties:
  ApplicationName: String
  Output:
    Output

```

## Properties

`ApplicationName`

Name of the application to which you want to add the output configuration.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Output`

An array of objects, each describing one output configuration. In the output
configuration, you specify the name of an in-application stream, a destination (that is,
an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream, or an AWS Lambda function), and record the formation to use when writing to the
destination.

_Required_: Yes

_Type_: [Output](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalytics-applicationoutput-output.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

### Adding an ApplicationOutput Resource

The following example creates an `ApplicationOutput` resource:

#### YAML

```yaml

Type: AWS::KinesisAnalytics::ApplicationOutput
    Properties:
      ApplicationName: !Ref BasicApplication
      Output:
        Name: "exampleOutput"
        DestinationSchema:
          RecordFormatType: "CSV"
        KinesisStreamsOutput:
          ResourceARN: !GetAtt OutputKinesisStream.Arn
          RoleARN: !GetAtt KinesisAnalyticsRole.Arn

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecordFormat

DestinationSchema
