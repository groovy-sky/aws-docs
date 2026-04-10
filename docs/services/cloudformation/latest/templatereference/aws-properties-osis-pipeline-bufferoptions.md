This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline BufferOptions

Options that specify the configuration of a persistent buffer.
To configure how OpenSearch Ingestion encrypts this data, set the `EncryptionAtRestOptions`. For more information, see [Persistent buffering](../../../opensearch-service/latest/developerguide/osis-features-overview.md#persistent-buffering).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PersistentBufferEnabled" : Boolean
}

```

### YAML

```yaml

  PersistentBufferEnabled: Boolean

```

## Properties

`PersistentBufferEnabled`

Whether persistent buffering should be enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OSIS::Pipeline

CloudWatchLogDestination

All content copied from https://docs.aws.amazon.com/.
