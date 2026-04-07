This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceGroupConfig EbsConfiguration

The Amazon EBS configuration of a cluster instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EbsBlockDeviceConfigs" : [ EbsBlockDeviceConfig, ... ],
  "EbsOptimized" : Boolean
}

```

### YAML

```yaml

  EbsBlockDeviceConfigs:
    - EbsBlockDeviceConfig
  EbsOptimized: Boolean

```

## Properties

`EbsBlockDeviceConfigs`

An array of Amazon EBS volume specifications attached to a cluster
instance.

_Required_: No

_Type_: Array of [EbsBlockDeviceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-instancegroupconfig-ebsblockdeviceconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EbsOptimized`

Indicates whether an Amazon EBS volume is EBS-optimized. The default is false. You should explicitly set this value to true to enable the Amazon EBS-optimized setting
for an EC2 instance.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EbsBlockDeviceConfig

MetricDimension
