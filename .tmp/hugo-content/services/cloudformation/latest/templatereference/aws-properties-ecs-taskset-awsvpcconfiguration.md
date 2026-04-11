This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskSet AwsVpcConfiguration

An object representing the networking details for a task or service. For example
`awsVpcConfiguration={subnets=["subnet-12344321"],securityGroups=["sg-12344321"]}`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssignPublicIp" : String,
  "SecurityGroups" : [ String, ... ],
  "Subnets" : [ String, ... ]
}

```

### YAML

```yaml

  AssignPublicIp: String
  SecurityGroups:
    - String
  Subnets:
    - String

```

## Properties

`AssignPublicIp`

Whether the task's elastic network interface receives a public IP address.

Consider the following when you set this value:

- When you use `create-service` or `update-service`, the
default is `DISABLED`.

- When the service `deploymentController` is `ECS`, the
value must be `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

The IDs of the security groups associated with the task or service. If you don't
specify a security group, the default security group for the VPC is used. There's a
limit of 5 security groups that can be specified.

###### Note

All specified security groups must be from the same VPC.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subnets`

The IDs of the subnets associated with the task or service. There's a limit of 16
subnets that can be specified.

###### Note

All specified subnets must be from the same VPC.

_Required_: Yes

_Type_: Array of String

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECS::TaskSet

CapacityProviderStrategyItem

All content copied from https://docs.aws.amazon.com/.
