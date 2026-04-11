This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider

Creates a capacity provider. Capacity providers are associated with a cluster and are
used in capacity provider strategies to facilitate cluster auto scaling. You can create
capacity providers for Amazon ECS Managed Instances and EC2 instances. AWS Fargate has the
predefined `FARGATE` and `FARGATE_SPOT` capacity providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::CapacityProvider",
  "Properties" : {
      "AutoScalingGroupProvider" : AutoScalingGroupProvider,
      "ClusterName" : String,
      "ManagedInstancesProvider" : ManagedInstancesProvider,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECS::CapacityProvider
Properties:
  AutoScalingGroupProvider:
    AutoScalingGroupProvider
  ClusterName: String
  ManagedInstancesProvider:
    ManagedInstancesProvider
  Name: String
  Tags:
    - Tag

```

## Properties

`AutoScalingGroupProvider`

The Auto Scaling group settings for the capacity provider.

_Required_: No

_Type_: [AutoScalingGroupProvider](aws-properties-ecs-capacityprovider-autoscalinggroupprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The cluster that this capacity provider is associated with. Managed instances capacity
providers are cluster-scoped, meaning they can only be used within their associated
cluster.

This is required for Managed instances.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedInstancesProvider`

The configuration for the Amazon ECS Managed Instances provider. This includes the
infrastructure role, the launch template configuration, and tag propagation
settings.

_Required_: No

_Type_: [ManagedInstancesProvider](aws-properties-ecs-capacityprovider-managedinstancesprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the capacity provider. If a name is specified, it cannot start with
`aws`, `ecs`, or `fargate`. If no name is specified, a
default name in the `CFNStackName-CFNResourceName-RandomString` format is
used.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The metadata that you apply to the capacity provider to help you categorize and
organize it. Each tag consists of a key and an optional value. You define both.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length - 128 Unicode characters in UTF-8

- Maximum value length - 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with
this prefix. Tags with this prefix do not count against your tags per resource
limit.

_Required_: No

_Type_: Array of [Tag](aws-properties-ecs-capacityprovider-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

In the following example, the `Ref` function returns the name of the capacity
provider, such as `MyStack-MyCapacityProvider-JrwYBzxovGfr`.

`{ "Ref": "MyCapacityProvider" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating an Amazon ECS capacity provider

The following example creates a capacity provider that uses the Auto Scaling group
`MyAutoScalingGroup`, has managed scaling and managed termination
protection enabled. This configuration is used for Amazon ECS cluster auto
scaling.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
    "MyCapacityProvider": {
            "Type": "AWS::ECS::CapacityProvider",
            "Properties": {
                "AutoScalingGroupProvider": {
                    "AutoScalingGroupArn": "arn:aws:autoscaling:us-west-2:123456789012:autoScalingGroup:a1b2c3d4-5678-90ab-cdef-EXAMPLE11111:autoScalingGroupName/MyAutoScalingGroup",
                    "ManagedScaling": {
                        "MaximumScalingStepSize": 10,
                        "MinimumScalingStepSize": 1,
                        "Status": "ENABLED",
                        "TargetCapacity": 100
                    },
                    "ManagedTerminationProtection": "ENABLED"
                },
                "Tags": [
                    {
                        "Key": "environment",
                        "Value": "production"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
 MyCapacityProvider:
    Type: AWS::ECS::CapacityProvider
    Properties:
        AutoScalingGroupProvider:
            AutoScalingGroupArn: arn:aws:autoscaling:us-west-2:123456789012:autoScalingGroup:a1b2c3d4-5678-90ab-cdef-EXAMPLE11111:autoScalingGroupName/MyAutoScalingGroup
            ManagedScaling:
                MaximumScalingStepSize: 10
                MinimumScalingStepSize: 1
                Status: ENABLED
                TargetCapacity: 100
            ManagedTerminationProtection: ENABLED
        Tags:
            - Key: environment
              Value: production
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon ECS

AcceleratorCountRequest

All content copied from https://docs.aws.amazon.com/.
