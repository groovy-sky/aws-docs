This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup TargetGroupInfo

The `TargetGroupInfo` property type specifies information about a target group
in Elastic Load Balancing to use in a deployment. Instances are registered as targets in a
target group, and traffic is routed to the target group. For more information, see [TargetGroupInfo](../../../../reference/codedeploy/latest/apireference/api-targetgroupinfo.md) in the _AWS CodeDeploy API Reference_

If you specify the `TargetGroupInfo` property, the
`DeploymentStyle.DeploymentOption` property must be set to
`WITH_TRAFFIC_CONTROL` for CodeDeploy to route your traffic using the
specified target groups.

`TargetGroupInfo` is a property of the [LoadBalancerInfo](../userguide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

For blue/green deployments, the name of the target group that instances in the original
environment are deregistered from, and instances in the replacement environment registered
with. For in-place deployments, the name of the target group that instances are deregistered
from, so they are not serving traffic during a deployment, and then re-registered with after
the deployment completes. No duplicates allowed.

###### Note

CloudFormation supports blue/green deployments on AWS Lambda compute
platforms only.

This value cannot exceed 32 characters, so you should use the `Name` property
of the target group, or the `TargetGroupName` attribute with the
`Fn::GetAtt` intrinsic function, as shown in the following example. Don't use the
group's Amazon Resource Name (ARN) or `TargetGroupFullName` attribute.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFilter

TargetGroupPairInfo

All content copied from https://docs.aws.amazon.com/.
