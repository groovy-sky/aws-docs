This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisVideo::Stream StreamStorageConfiguration

The configuration for stream storage, including the default storage tier for stream data. This configuration determines how stream data is stored and accessed, with different tiers offering varying levels of performance and cost optimization.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultStorageTier" : String
}

```

### YAML

```yaml

  DefaultStorageTier: String

```

## Properties

`DefaultStorageTier`

The default storage tier for the stream data. This setting determines the storage class used for stream data, affecting both performance characteristics and storage costs.

Available storage tiers:

- `HOT` \- Optimized for frequent access with the lowest latency and highest performance. Ideal for real-time applications and frequently accessed data.

- `WARM` \- Balanced performance and cost for moderately accessed data. Suitable for data that is accessed regularly but not continuously.

_Required_: No

_Type_: String

_Allowed values_: `HOT | WARM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KinesisVideo::Stream

Tag

All content copied from https://docs.aws.amazon.com/.
