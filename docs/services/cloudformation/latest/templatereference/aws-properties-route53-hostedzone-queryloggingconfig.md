This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::HostedZone QueryLoggingConfig

A complex type that contains information about a configuration for DNS query
logging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsLogGroupArn" : String
}

```

### YAML

```yaml

  CloudWatchLogsLogGroupArn: String

```

## Properties

`CloudWatchLogsLogGroupArn`

The Amazon Resource Name (ARN) of the CloudWatch Logs log group that Amazon Route 53
is publishing logs to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#aws-resource-route53-hostedzone-return-values)
in the topic
[AWS::Route53::HostedZone](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html)

- [QueryLoggingConfig](https://docs.aws.amazon.com/Route53/latest/APIReference/API_QueryLoggingConfig.html)
in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HostedZoneTag

VPC
