This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::HealthCheck AlarmIdentifier

A complex type that identifies the CloudWatch alarm that you want Amazon Route 53
health checkers to use to determine whether the specified health check is
healthy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Region" : String
}

```

### YAML

```yaml

  Name: String
  Region: String

```

## Properties

`Name`

The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use
to determine whether this health check is healthy.

###### Note

Route 53 supports CloudWatch alarms with the following features:

- Standard-resolution metrics. High-resolution metrics aren't supported. For
more information, see [High-Resolution Metrics](../../../amazoncloudwatch/latest/developerguide/publishingmetrics.md#high-resolution-metrics) in the _Amazon CloudWatch User_
_Guide_.

- Statistics: Average, Minimum, Maximum, Sum, and SampleCount. Extended
statistics aren't supported.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

For the CloudWatch alarm that you want Route 53 health checkers to use to determine
whether this health check is healthy, the region that the alarm was created in.

For the current list of CloudWatch regions, see [Amazon CloudWatch endpoints and\
quotas](../../../../general/latest/gr/cw-region.md) in the _Amazon Web Services General_
_Reference_.

_Required_: Yes

_Type_: String

_Allowed values_: `us-east-1 | us-east-2 | us-west-1 | us-west-2 | ca-central-1 | eu-central-1 | eu-central-2 | eu-west-1 | eu-west-2 | eu-west-3 | ap-east-1 | me-south-1 | me-central-1 | ap-south-1 | ap-south-2 | ap-southeast-1 | ap-southeast-2 | ap-southeast-3 | ap-northeast-1 | ap-northeast-2 | ap-northeast-3 | eu-north-1 | sa-east-1 | cn-northwest-1 | cn-north-1 | af-south-1 | eu-south-1 | eu-south-2 | us-gov-west-1 | us-gov-east-1 | us-iso-east-1 | us-iso-west-1 | us-isob-east-1 | ap-southeast-4 | il-central-1 | ca-west-1 | ap-southeast-5 | mx-central-1 | us-isof-south-1 | us-isof-east-1 | ap-southeast-7 | ap-east-2 | eu-isoe-west-1 | ap-southeast-6 | us-isob-west-1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](../userguide/aws-resource-route53-healthcheck.md#aws-resource-route53-healthcheck-return-values)
in the topic
[AWS::Route53::HealthCheck](../userguide/aws-resource-route53-healthcheck.md)

- [AlarmIdentifier](../../../../reference/route53/latest/apireference/api-alarmidentifier.md) in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Route53::HealthCheck

HealthCheckConfig

All content copied from https://docs.aws.amazon.com/.
