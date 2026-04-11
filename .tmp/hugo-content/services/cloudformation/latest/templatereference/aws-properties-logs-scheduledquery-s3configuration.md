This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::ScheduledQuery S3Configuration

Configuration for Amazon S3 destination where scheduled query results are delivered.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationIdentifier" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  DestinationIdentifier: String
  RoleArn: String

```

## Properties

`DestinationIdentifier`

The Amazon S3 URI where query results are delivered. Must be a valid S3 URI format.

_Required_: Yes

_Type_: String

_Pattern_: `^s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants permissions to write query results to the specified
Amazon S3 destination.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationConfiguration

TagsItems

All content copied from https://docs.aws.amazon.com/.
