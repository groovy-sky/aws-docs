This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate NetworkPerformanceOptions

Contains settings for the network performance options for the instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BandwidthWeighting" : String
}

```

### YAML

```yaml

  BandwidthWeighting: String

```

## Properties

`BandwidthWeighting`

Specify the bandwidth weighting option to boost the associated type of baseline bandwidth,
as follows:

default

This option uses the standard bandwidth configuration for your instance type.

vpc-1

This option boosts your networking baseline bandwidth and reduces your EBS baseline
bandwidth.

ebs-1

This option boosts your EBS baseline bandwidth and reduces your networking baseline
bandwidth.

_Required_: No

_Type_: String

_Allowed values_: `default | vpc-1 | ebs-1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkInterfaceCount

Placement

All content copied from https://docs.aws.amazon.com/.
