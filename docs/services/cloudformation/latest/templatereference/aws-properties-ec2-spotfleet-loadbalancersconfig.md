This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet LoadBalancersConfig

Specifies the Classic Load Balancers and target groups to attach to a Spot Fleet
request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClassicLoadBalancersConfig" : ClassicLoadBalancersConfig,
  "TargetGroupsConfig" : TargetGroupsConfig
}

```

### YAML

```yaml

  ClassicLoadBalancersConfig:
    ClassicLoadBalancersConfig
  TargetGroupsConfig:
    TargetGroupsConfig

```

## Properties

`ClassicLoadBalancersConfig`

The Classic Load Balancers.

_Required_: No

_Type_: [ClassicLoadBalancersConfig](aws-properties-ec2-spotfleet-classicloadbalancersconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetGroupsConfig`

The target groups.

_Required_: No

_Type_: [TargetGroupsConfig](aws-properties-ec2-spotfleet-targetgroupsconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplateOverrides

MemoryGiBPerVCpuRequest

All content copied from https://docs.aws.amazon.com/.
