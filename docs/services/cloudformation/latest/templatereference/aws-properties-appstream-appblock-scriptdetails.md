This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::AppBlock ScriptDetails

The details of the script.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutableParameters" : String,
  "ExecutablePath" : String,
  "ScriptS3Location" : S3Location,
  "TimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  ExecutableParameters: String
  ExecutablePath: String
  ScriptS3Location:
    S3Location
  TimeoutInSeconds: Integer

```

## Properties

`ExecutableParameters`

The parameters used in the run path for the script.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutablePath`

The run path for the script.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScriptS3Location`

The S3 object location of the script.

_Required_: Yes

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appstream-appblock-s3location.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutInSeconds`

The run timeout, in seconds, for the script.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Location

Tag
