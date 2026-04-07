This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLiftStreams::Application

The `AWS::GameLiftStreams::Application` resource defines an Amazon GameLift Streams application. An application specifies
the content that you want to stream, such as a game or other software, and its runtime environment (Microsoft Windows, Ubuntu, or
Proton).

Before you create an Amazon GameLift Streams application, upload your _uncompressed_ game files (do not upload a
.zip file) to an Amazon Simple Storage Service (Amazon S3) standard bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLiftStreams::Application",
  "Properties" : {
      "ApplicationLogOutputUri" : String,
      "ApplicationLogPaths" : [ String, ... ],
      "ApplicationSourceUri" : String,
      "Description" : String,
      "ExecutablePath" : String,
      "RuntimeEnvironment" : RuntimeEnvironment,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::GameLiftStreams::Application
Properties:
  ApplicationLogOutputUri: String
  ApplicationLogPaths:
    - String
  ApplicationSourceUri: String
  Description: String
  ExecutablePath: String
  RuntimeEnvironment:
    RuntimeEnvironment
  Tags:
    Key: Value

```

## Properties

`ApplicationLogOutputUri`

An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs. Required if you specify one or more `ApplicationLogPaths`.

_Required_: No

_Type_: String

_Pattern_: `^$|^s3://([a-zA-Z0-9][a-zA-Z0-9._-]{1,61}[a-zA-Z0-9])(/[a-zA-Z0-9._-]+)*/?$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationLogPaths`

Locations of log files that your content generates during a stream session. Enter path
values that are relative to the `ApplicationSourceUri` location.
You can specify up to 10 log paths.
Amazon GameLift Streams uploads designated log files to the Amazon S3 bucket that you specify in `ApplicationLogOutputUri`
at the end of a stream session. To retrieve stored log files, call [GetStreamSession](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html)
and get the `LogFileLocationUri`.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `1024 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationSourceUri`

The location of the content that you want to stream. Enter an Amazon S3 URI to a bucket that contains your game or other application. The
location can have a multi-level prefix structure, but it must include all the files needed to run the content. Amazon GameLift Streams copies everything
under the specified location.

This value is immutable. To designate a different content location, create a new application.

###### Note

The Amazon S3 bucket and the Amazon GameLift Streams application must be in the same AWS Region.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A human-readable label for the application. You can update this value later.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_.!+@/][a-zA-Z0-9-_.!+@/ ]*$`

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutablePath`

The relative path and file name of the executable file that Amazon GameLift Streams will stream. Specify a path relative to the location set in
`ApplicationSourceUri`. The file must be contained within the application's root folder. For Windows applications, the file
must be a valid Windows executable or batch file with a filename ending in .exe, .cmd, or .bat. For Linux applications, the file must be a
valid Linux binary executable or a script that contains an initial interpreter line starting with a shebang (' `#!`').

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RuntimeEnvironment`

A set of configuration settings to run the application on a stream group. This configures the operating system, and can include
compatibility layers and other drivers.

_Required_: Yes

_Type_: [RuntimeEnvironment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gameliftstreams-application-runtimeenvironment.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of labels to assign to the new application resource. Tags are developer-defined key-value pairs. Tagging AWS resources is
useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the _AWS General Reference_.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns an [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html)
that uniquely identifies the application resource across all AWS Regions. For example:

`arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html)
that uniquely identifies the application resource across all AWS Regions. For example:

`arn:aws:gameliftstreams:us-west-2:123456789012:application/a-9ZY8X7Wv6`.

`Id`

An ID that uniquely identifies the application resource. For example: `a-9ZY8X7Wv6`.

## See also

- [Prepare an application in Amazon GameLift\
Streams](https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/applications.html) in the _Amazon GameLift Streams Developer Guide_

- [CreateApplication](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_CreateApplication.html) in the
_Amazon GameLift Streams API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon GameLift Streams

RuntimeEnvironment
