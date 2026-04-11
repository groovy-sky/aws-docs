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

_Type_: [S3Location](aws-properties-kafkaconnect-customplugin-s3location.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomPluginFileDescription

S3Location

All content copied from https://docs.aws.amazon.com/.
