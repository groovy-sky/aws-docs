This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig EbsConfiguration

`EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.

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

_Type_: Array of [EbsBlockDeviceConfig](aws-properties-emr-instancefleetconfig-ebsblockdeviceconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EbsOptimized`

Indicates whether an Amazon EBS volume is EBS-optimized. The default is false. You should explicitly set this value to true to enable the Amazon EBS-optimized setting
for an EC2 instance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EbsBlockDeviceConfig

InstanceFleetProvisioningSpecifications

All content copied from https://docs.aws.amazon.com/.
