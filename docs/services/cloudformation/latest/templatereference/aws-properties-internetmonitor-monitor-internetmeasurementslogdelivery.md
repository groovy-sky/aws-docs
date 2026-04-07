This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InternetMonitor::Monitor InternetMeasurementsLogDelivery

Publish internet measurements to an Amazon S3 bucket in addition to CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Config" : S3Config
}

```

### YAML

```yaml

  S3Config:
    S3Config

```

## Properties

`S3Config`

The configuration for publishing Amazon CloudWatch Internet Monitor internet measurements to Amazon S3.

_Required_: No

_Type_: [S3Config](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-internetmonitor-monitor-s3config.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HealthEventsConfig

LocalHealthEventsConfig
