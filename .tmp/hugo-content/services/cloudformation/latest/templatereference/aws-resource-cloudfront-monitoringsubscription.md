This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::MonitoringSubscription

A monitoring subscription. This structure contains information about whether
additional CloudWatch metrics are enabled for a given CloudFront distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::MonitoringSubscription",
  "Properties" : {
      "DistributionId" : String,
      "MonitoringSubscription" : MonitoringSubscription
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::MonitoringSubscription
Properties:
  DistributionId: String
  MonitoringSubscription:
    MonitoringSubscription

```

## Properties

`DistributionId`

The ID of the distribution that you are enabling metrics for.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitoringSubscription`

A subscription configuration for additional CloudWatch metrics.

_Required_: Yes

_Type_: [MonitoringSubscription](aws-properties-cloudfront-monitoringsubscription-monitoringsubscription.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier for the monitoring subscription, which
is the same as the distribution ID of the distribution that the subscription is for.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportSource

MonitoringSubscription

All content copied from https://docs.aws.amazon.com/.
