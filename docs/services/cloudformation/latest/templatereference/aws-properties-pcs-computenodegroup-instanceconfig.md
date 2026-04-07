This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::ComputeNodeGroup InstanceConfig

An EC2 instance configuration AWS PCS uses to launch compute nodes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceType" : String
}

```

### YAML

```yaml

  InstanceType: String

```

## Properties

`InstanceType`

The EC2 instance type that AWS PCS can provision in the compute node
group.

Example: `t2.xlarge`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ErrorInfo

ScalingConfiguration
