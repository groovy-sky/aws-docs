This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup LoadBalancerInfo

The `LoadBalancerInfo` property type specifies information about the load
balancer or target group used for an AWS CodeDeploy deployment group. For more
information, see [Integrating\
CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the _AWS CodeDeploy User Guide_.

For CloudFormation to use the properties specified in `LoadBalancerInfo`,
the `DeploymentStyle.DeploymentOption` property must be set to
`WITH_TRAFFIC_CONTROL`. If `DeploymentStyle.DeploymentOption` is not
set to `WITH_TRAFFIC_CONTROL`, CloudFormation ignores any settings specified
in `LoadBalancerInfo`.

###### Note

CloudFormation supports blue/green deployments on the AWS Lambda
compute platform only.

`LoadBalancerInfo` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ElbInfoList" : [ ELBInfo, ... ],
  "TargetGroupInfoList" : [ TargetGroupInfo, ... ],
  "TargetGroupPairInfoList" : [ TargetGroupPairInfo, ... ]
}

```

### YAML

```yaml

  ElbInfoList:
    - ELBInfo
  TargetGroupInfoList:
    - TargetGroupInfo
  TargetGroupPairInfoList:
    - TargetGroupPairInfo

```

## Properties

`ElbInfoList`

An array that contains information about the load balancers to use for load balancing
in a deployment. If you're using Classic Load Balancers, specify those load balancers in
this array.

###### Note

You can add up to 10 load balancers to the array.

###### Note

If you're using Application Load Balancers or Network Load Balancers, use the
`targetGroupInfoList` array instead of this one.

_Required_: No

_Type_: Array of [ELBInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentgroup-elbinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupInfoList`

An array that contains information about the target groups to use for load balancing
in a deployment. If you're using Application Load Balancers and Network Load Balancers,
specify their associated target groups in this array.

###### Note

You can add up to 10 target groups to the array.

###### Note

If you're using Classic Load Balancers, use the `elbInfoList` array
instead of this one.

_Required_: Conditional

_Type_: Array of [TargetGroupInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentgroup-targetgroupinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupPairInfoList`

The target group pair information. This is an array of
`TargeGroupPairInfo` objects with a maximum size of one.

_Required_: No

_Type_: Array of [TargetGroupPairInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GreenFleetProvisioningOption

OnPremisesTagSet
