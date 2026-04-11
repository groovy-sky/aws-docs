This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket LambdaConfiguration

Describes the AWS Lambda functions to invoke and the events for which to invoke them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Event" : String,
  "Filter" : NotificationFilter,
  "Function" : String
}

```

### YAML

```yaml

  Event: String
  Filter:
    NotificationFilter
  Function: String

```

## Properties

`Event`

The Amazon S3 bucket event for which to invoke the AWS Lambda function. For more information, see [Supported Event Types](../../../s3/latest/dev/notificationhowto.md) in
the _Amazon S3 User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

The filtering rules that determine which objects invoke the AWS Lambda
function. For example, you can create a filter so that only image files with a
`.jpg` extension invoke the function when they are added to the Amazon S3
bucket.

_Required_: No

_Type_: [NotificationFilter](aws-properties-s3-bucket-notificationfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Function`

The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon S3 invokes when the specified event
type occurs.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JournalTableConfiguration

LifecycleConfiguration

All content copied from https://docs.aws.amazon.com/.
