This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup ELBInfo

The `ELBInfo` property type specifies information about the Elastic Load Balancing load balancer used for an CodeDeploy deployment group.

If you specify the `ELBInfo` property, the
`DeploymentStyle.DeploymentOption` property must be set to
`WITH_TRAFFIC_CONTROL` for AWS CodeDeploy to route your traffic using
the specified load balancers.

`ELBInfo` is a property of the [AWS CodeDeploy DeploymentGroup LoadBalancerInfo](../userguide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.md) property type.

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

For blue/green deployments, the name of the load balancer that is used to route traffic
from original instances to replacement instances in a blue/green deployment. For in-place
deployments, the name of the load balancer that instances are deregistered from so they are
not serving traffic during a deployment, and then re-registered with after the deployment is
complete.

###### Note

CloudFormation supports blue/green deployments on AWS Lambda compute
platforms only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ECSService

GitHubLocation

All content copied from https://docs.aws.amazon.com/.
