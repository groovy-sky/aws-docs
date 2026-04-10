This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster EBSStorageInfo

Contains information about the EBS storage volumes attached to the broker nodes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProvisionedThroughput" : ProvisionedThroughput,
  "VolumeSize" : Integer
}

```

### YAML

```yaml

  ProvisionedThroughput:
    ProvisionedThroughput
  VolumeSize: Integer

```

## Properties

`ProvisionedThroughput`

EBS volume provisioned throughput information.

_Required_: No

_Type_: [ProvisionedThroughput](aws-properties-msk-cluster-provisionedthroughput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSize`

The size in GiB of the EBS volume for the data drive on each broker node.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectivityInfo

EncryptionAtRest

All content copied from https://docs.aws.amazon.com/.
