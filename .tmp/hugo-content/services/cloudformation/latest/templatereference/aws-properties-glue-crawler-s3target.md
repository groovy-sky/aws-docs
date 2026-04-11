This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler S3Target

Specifies a data store in Amazon Simple Storage Service (Amazon S3).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionName" : String,
  "DlqEventQueueArn" : String,
  "EventQueueArn" : String,
  "Exclusions" : [ String, ... ],
  "Path" : String,
  "SampleSize" : Integer
}

```

### YAML

```yaml

  ConnectionName: String
  DlqEventQueueArn: String
  EventQueueArn: String
  Exclusions:
    - String
  Path: String
  SampleSize: Integer

```

## Properties

`ConnectionName`

The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DlqEventQueueArn`

A valid Amazon dead-letter SQS ARN. For example, `arn:aws:sqs:region:account:deadLetterQueue`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventQueueArn`

A valid Amazon SQS ARN. For example, `arn:aws:sqs:region:account:sqs`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exclusions`

A list of glob patterns used to exclude from the crawl. For more information, see
[Catalog Tables\
with a Crawler](../../../glue/latest/dg/add-crawler.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to the Amazon S3 target.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleSize`

Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset. If not set, all the files are crawled. A valid value is an integer between 1 and 249.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecrawlPolicy

Schedule

All content copied from https://docs.aws.amazon.com/.
