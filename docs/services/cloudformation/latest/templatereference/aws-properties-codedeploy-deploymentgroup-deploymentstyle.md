This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup DeploymentStyle

Information about the type of deployment, either in-place or blue/green, you want to
run and whether to route deployment traffic behind a load balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentOption" : String,
  "DeploymentType" : String
}

```

### YAML

```yaml

  DeploymentOption: String
  DeploymentType: String

```

## Properties

`DeploymentOption`

Indicates whether to route deployment traffic behind a load balancer.

###### Note

An Amazon EC2Application Load Balancer or Network Load Balancer is required for an Amazon ECS
deployment.

_Required_: No

_Type_: String

_Allowed values_: `WITH_TRAFFIC_CONTROL | WITHOUT_TRAFFIC_CONTROL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentType`

Indicates whether to run an in-place or blue/green deployment.

_Required_: No

_Type_: String

_Allowed values_: `IN_PLACE | BLUE_GREEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [EC2TagFilter](../../../../reference/codedeploy/latest/apireference/api-ec2tagfilter.md) in the _AWS CodeDeploy API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentReadyOption

EC2TagFilter

All content copied from https://docs.aws.amazon.com/.
