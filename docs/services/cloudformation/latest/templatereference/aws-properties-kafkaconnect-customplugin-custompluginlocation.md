This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::CustomPlugin CustomPluginLocation

Information about the location of a custom plugin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Location" : S3Location
}

```

### YAML

```yaml

  S3Location:
    S3Location

```

## Properties

`S3Location`

The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin
file stored in Amazon S3.

_Required_: Yes

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kafkaconnect-customplugin-s3location.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomPluginFileDescription

S3Location
