This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet BaselineEbsBandwidthMbpsRequest

The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps. For more information, see
[Amazon\
EBS–optimized instances](../../../ec2/latest/userguide/ebs-optimized.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Integer,
  "Min" : Integer
}

```

### YAML

```yaml

  Max: Integer
  Min: Integer

```

## Properties

`Max`

The maximum baseline bandwidth, in Mbps. To specify no maximum limit, omit
this parameter.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Min`

The minimum baseline bandwidth, in Mbps. To specify no minimum limit, omit
this parameter.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AcceleratorTotalMemoryMiBRequest

BaselinePerformanceFactorsRequest

All content copied from https://docs.aws.amazon.com/.
