This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer

Creates or updates a _log transformer_ for a single log group. You
use log transformers to transform log events into a different format, making them easier
for you to process and analyze. You can also transform logs from different sources into
standardized formats that contains relevant, source-specific information.

After you have created a transformer, CloudWatch Logs performs the transformations at
the time of log ingestion. You can then refer to the transformed versions of the logs
during operations such as querying with CloudWatch Logs Insights or creating metric
filters or subscription filers.

You can also use a transformer to copy metadata from metadata keys into the log events
themselves. This metadata can include log group name, log stream name, account ID and
Region.

A transformer for a log group is a series of processors, where each processor applies
one type of transformation to the log events ingested into this log group. The
processors work one after another, in the order that you list them, like a pipeline. For
more information about the available processors to use in a transformer, see [Processors that you can use](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-Processors).

Having log events in standardized format enables visibility across your applications
for your log analysis, reporting, and alarming needs. CloudWatch Logs provides
transformation for common log types with out-of-the-box transformation templates for
major AWS log sources such as VPC flow logs, Lambda, and Amazon RDS. You
can use pre-built transformation templates or create custom transformation
policies.

You can create transformers only for the log groups in the Standard log class.

You can also set up a transformer at the account level. For more information, see
[PutAccountPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putaccountpolicy.md). If there is both a log-group level transformer created
with `PutTransformer` and an account-level transformer that could apply to
the same log group, the log group uses only the log-group level transformer. It ignores
the account-level transformer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::Transformer",
  "Properties" : {
      "LogGroupIdentifier" : String,
      "TransformerConfig" : [ Processor, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Logs::Transformer
Properties:
  LogGroupIdentifier: String
  TransformerConfig:
    - Processor

```

## Properties

`LogGroupIdentifier`

Specify either the name or ARN of the log group to create the transformer for.

_Required_: Yes

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransformerConfig`

This structure is an array that contains the configuration of this log transformer. A
log transformer is an array of processors, where each processor applies one type of
transformation to the log events that are ingested.

_Required_: Yes

_Type_: Array of [Processor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-processor.html)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Logs::SubscriptionFilter

AddKeyEntry
