This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service NetworkConfiguration

The network configuration for a task or service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsvpcConfiguration" : AwsVpcConfiguration
}

```

### YAML

```yaml

  AwsvpcConfiguration:
    AwsVpcConfiguration

```

## Properties

`AwsvpcConfiguration`

The VPC subnets and security groups that are associated with a task.

###### Note

All specified subnets and security groups must be from the same VPC.

_Required_: No

_Type_: [AwsVpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-service-awsvpcconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LogConfiguration

PlacementConstraint
