This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::InfluxDBCluster LogDeliveryConfiguration

Configuration for sending InfluxDB engine logs to a specified S3 bucket.

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

Configuration for S3 bucket log delivery.

_Required_: Yes

_Type_: [S3Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-timestream-influxdbcluster-s3configuration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Timestream::InfluxDBCluster

S3Configuration
