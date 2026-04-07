This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskSet NetworkConfiguration

The network configuration for a task or service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsVpcConfiguration" : AwsVpcConfiguration
}

```

### YAML

```yaml

  AwsVpcConfiguration:
    AwsVpcConfiguration

```

## Properties

`AwsVpcConfiguration`

The VPC subnets and security groups that are associated with a task.

###### Note

All specified subnets and security groups must be from the same VPC.

_Required_: No

_Type_: [AwsVpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-awsvpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LoadBalancer

Scale
