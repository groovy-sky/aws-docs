This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster ProvisionedThroughput

Contains information about provisioned throughput for EBS storage volumes attached to kafka broker nodes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "VolumeThroughput" : Integer
}

```

### YAML

```yaml

  Enabled: Boolean
  VolumeThroughput: Integer

```

## Properties

`Enabled`

Provisioned throughput is on or off.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeThroughput`

Throughput value of the EBS volumes for the data drive on each kafka broker node in MiB per second.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Prometheus

PublicAccess

All content copied from https://docs.aws.amazon.com/.
