This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::Discovery

###### Note

If you have existing `AWS::ApplicationSignals::Discovery` resources that were created prior to the Application Map release, you will need to delete and recreate these resources in your account to enable Application Map.

Enables this AWS account to be able to use CloudWatch Application Signals by creating the `AWSServiceRoleForCloudWatchApplicationSignals ` service-linked role. This service-linked role has the following permissions:

- `xray:GetServiceGraph`

- `logs:StartQuery`

- `logs:GetQueryResults`

- `cloudwatch:GetMetricData`

- `cloudwatch:ListMetrics`

- `tag:GetResources`

- `autoscaling:DescribeAutoScalingGroups`

A service-linked CloudTrail event channel is created to process CloudTrail events and return change event information. This includes last deployment time, userName, eventName, and other event metadata.

After completing this step, you still need to instrument your Java and Python applications to send data to Application Signals. For more information, see
[Enabling Application Signals](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Application-Signals-Enable.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApplicationSignals::Discovery"
}

```

### YAML

```yaml

Type: AWS::ApplicationSignals::Discovery
```

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the 12 digit AWS Account ID for the account.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The 12 digit AWS Account ID for the account.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatch Application Signals

AWS::ApplicationSignals::GroupingConfiguration
