This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline LogPublishingOptions

Container for the values required to configure logging for the pipeline. If you don't
specify these values, OpenSearch Ingestion will not publish logs from your application to
CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogDestination" : CloudWatchLogDestination,
  "IsLoggingEnabled" : Boolean
}

```

### YAML

```yaml

  CloudWatchLogDestination:
    CloudWatchLogDestination
  IsLoggingEnabled: Boolean

```

## Properties

`CloudWatchLogDestination`

The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch Logs. This
parameter is required if `IsLoggingEnabled` is set to `true`.

_Required_: No

_Type_: [CloudWatchLogDestination](aws-properties-osis-pipeline-cloudwatchlogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsLoggingEnabled`

Whether logs should be published.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionAtRestOptions

ResourcePolicy

All content copied from https://docs.aws.amazon.com/.
