This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::KeyValueStore ImportSource

The import source for the key value store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceArn" : String,
  "SourceType" : String
}

```

### YAML

```yaml

  SourceArn: String
  SourceType: String

```

## Properties

`SourceArn`

The Amazon Resource Name (ARN) of the import source for the key value store.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceType`

The source type of the import source for the key value store.

_Required_: Yes

_Type_: String

_Allowed values_: `S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFront::KeyValueStore

AWS::CloudFront::MonitoringSubscription
