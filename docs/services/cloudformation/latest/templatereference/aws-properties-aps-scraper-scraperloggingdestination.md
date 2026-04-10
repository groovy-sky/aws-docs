This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper ScraperLoggingDestination

The destination where scraper logs are sent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogDestination
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogDestination

```

## Properties

`CloudWatchLogs`

The CloudWatch Logs configuration for the scraper logging destination.

_Required_: No

_Type_: [CloudWatchLogDestination](aws-properties-aps-scraper-cloudwatchlogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScraperLoggingConfiguration

Source

All content copied from https://docs.aws.amazon.com/.
