This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ApplicationCodeConfiguration

Describes code configuration for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeContent" : CodeContent,
  "CodeContentType" : String
}

```

### YAML

```yaml

  CodeContent:
    CodeContent
  CodeContentType: String

```

## Properties

`CodeContent`

The location and type of the application code.

_Required_: Yes

_Type_: [CodeContent](aws-properties-kinesisanalyticsv2-application-codecontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeContentType`

Specifies whether the code content is in text or zip format.

_Required_: Yes

_Type_: String

_Allowed values_: `PLAINTEXT | ZIPFILE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ApplicationCodeConfiguration](../../../managed-flink/latest/apiv2/api-applicationcodeconfiguration.md) in the _Amazon Kinesis_
_Data Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KinesisAnalyticsV2::Application

ApplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
