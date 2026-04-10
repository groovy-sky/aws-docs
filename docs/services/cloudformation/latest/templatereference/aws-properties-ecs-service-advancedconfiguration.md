This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service AdvancedConfiguration

The advanced settings for a load balancer used in blue/green deployments. Specify the
alternate target group, listener rules, and IAM role required for traffic shifting during
blue/green deployments. For more information, see [Required resources for Amazon ECS\
blue/green deployments](../../../amazonecs/latest/developerguide/blue-green-deployment-implementation.md) in the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlternateTargetGroupArn" : String,
  "ProductionListenerRule" : String,
  "RoleArn" : String,
  "TestListenerRule" : String
}

```

### YAML

```yaml

  AlternateTargetGroupArn: String
  ProductionListenerRule: String
  RoleArn: String
  TestListenerRule: String

```

## Properties

`AlternateTargetGroupArn`

The Amazon Resource Name (ARN) of the alternate target group for Amazon ECS blue/green deployments.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductionListenerRule`

The Amazon Resource Name (ARN) that that identifies the production listener rule (in the case of an Application Load Balancer) or
listener (in the case for an Network Load Balancer) for routing production traffic.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that grants Amazon ECS permission to call the Elastic Load Balancing APIs for you.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TestListenerRule`

The Amazon Resource Name (ARN) that identifies ) that identifies the test listener rule (in the case of an Application Load Balancer) or
listener (in the case for an Network Load Balancer) for routing test traffic.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECS::Service

AwsVpcConfiguration

All content copied from https://docs.aws.amazon.com/.
