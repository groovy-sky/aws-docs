This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet ClassicLoadBalancersConfig

Specifies the Classic Load Balancers to attach to a Spot Fleet. Spot Fleet registers the
running Spot Instances with these Classic Load Balancers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClassicLoadBalancers" : [ ClassicLoadBalancer, ... ]
}

```

### YAML

```yaml

  ClassicLoadBalancers:
    - ClassicLoadBalancer

```

## Properties

`ClassicLoadBalancers`

One or more Classic Load Balancers.

_Required_: Yes

_Type_: Array of [ClassicLoadBalancer](aws-properties-ec2-spotfleet-classicloadbalancer.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClassicLoadBalancer

CpuPerformanceFactorRequest

All content copied from https://docs.aws.amazon.com/.
