This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem ReadCacheConfiguration

The configuration for the optional provisioned SSD read cache on Amazon FSx for OpenZFS file systems that use the Intelligent-Tiering storage class.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SizeGiB" : Integer,
  "SizingMode" : String
}

```

### YAML

```yaml

  SizeGiB: Integer
  SizingMode: String

```

## Properties

`SizeGiB`

Required if `SizingMode` is set to `USER_PROVISIONED`. Specifies the size of the file system's SSD read cache, in gibibytes (GiB).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizingMode`

Specifies how the provisioned SSD read cache is sized, as follows:

- Set to `NO_CACHE` if you do not want to use an SSD read cache with your Intelligent-Tiering file system.

- Set to `USER_PROVISIONED` to specify the exact size of your SSD read cache.

- Set to `PROPORTIONAL_TO_THROUGHPUT_CAPACITY` to have your SSD read cache automatically sized based on your throughput capacity.

_Required_: No

_Type_: String

_Allowed values_: `NO_CACHE | USER_PROVISIONED | PROPORTIONAL_TO_THROUGHPUT_CAPACITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenZFSConfiguration

RootVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
