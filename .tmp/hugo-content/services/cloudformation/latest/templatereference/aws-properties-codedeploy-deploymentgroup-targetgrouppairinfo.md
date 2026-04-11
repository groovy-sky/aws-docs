This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup TargetGroupPairInfo

Information about two target groups and how traffic is routed during an Amazon ECS deployment. An optional test traffic route can be specified.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProdTrafficRoute" : TrafficRoute,
  "TargetGroups" : [ TargetGroupInfo, ... ],
  "TestTrafficRoute" : TrafficRoute
}

```

### YAML

```yaml

  ProdTrafficRoute:
    TrafficRoute
  TargetGroups:
    - TargetGroupInfo
  TestTrafficRoute:
    TrafficRoute

```

## Properties

`ProdTrafficRoute`

The path used by a load balancer to route production traffic when an Amazon ECS deployment is complete.

_Required_: No

_Type_: [TrafficRoute](aws-properties-codedeploy-deploymentgroup-trafficroute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroups`

One pair of target groups. One is associated with the original task set. The second
is associated with the task set that serves traffic after the deployment is complete.

_Required_: No

_Type_: Array of [TargetGroupInfo](aws-properties-codedeploy-deploymentgroup-targetgroupinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TestTrafficRoute`

An optional path used by a load balancer to route test traffic after an Amazon ECS deployment. Validation can occur while test traffic is served during a
deployment.

_Required_: No

_Type_: [TrafficRoute](aws-properties-codedeploy-deploymentgroup-trafficroute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetGroupInfo

TrafficRoute

All content copied from https://docs.aws.amazon.com/.
