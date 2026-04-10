This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline CloudWatchLogDestination

The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroup" : String
}

```

### YAML

```yaml

  LogGroup: String

```

## Properties

`LogGroup`

The name of the CloudWatch Logs group to send pipeline logs to. You can specify an existing
log group or create a new one. For example,
`/aws/vendedlogs/OpenSearchService/pipelines`.

_Required_: Yes

_Type_: String

_Pattern_: `\/aws\/vendedlogs\/[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BufferOptions

EncryptionAtRestOptions

All content copied from https://docs.aws.amazon.com/.
