This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::ScheduledQuery DestinationConfiguration

Configuration for where to deliver scheduled query results. Specifies the destination type
and associated settings for result delivery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Configuration" : S3Configuration
}

```

### YAML

```yaml

  S3Configuration:
    S3Configuration

```

## Properties

`S3Configuration`

Configuration for delivering query results to Amazon S3.

_Required_: No

_Type_: [S3Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-scheduledquery-s3configuration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Logs::ScheduledQuery

S3Configuration
