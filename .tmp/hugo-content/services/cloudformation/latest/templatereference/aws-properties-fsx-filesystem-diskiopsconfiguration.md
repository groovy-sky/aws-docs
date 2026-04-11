This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem DiskIopsConfiguration

The SSD IOPS (input/output operations per second) configuration for an Amazon FSx for NetApp ONTAP, Amazon FSx for Windows File Server, or FSx for OpenZFS file system. By default, Amazon FSx
automatically provisions 3 IOPS per GB of storage capacity. You can provision additional IOPS per
GB of storage. The configuration consists of the total number of provisioned SSD IOPS
and how it is was provisioned, or the mode (by the customer or by Amazon FSx).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iops" : Integer,
  "Mode" : String
}

```

### YAML

```yaml

  Iops: Integer
  Mode: String

```

## Properties

`Iops`

The total number of SSD IOPS provisioned for the file system.

The minimum and maximum values for this property depend on the value of `HAPairs` and `StorageCapacity`. The minimum value is calculated as `StorageCapacity` \\* 3 \* `HAPairs` (3 IOPS per GB of `StorageCapacity`). The maximum value is calculated as 200,000 \* `HAPairs`.

Amazon FSx responds with an HTTP status code 400 (Bad Request) if the value of `Iops` is outside of the minimum or maximum values.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Specifies whether the file system is
using the `AUTOMATIC` setting of SSD IOPS of 3 IOPS per GB of storage capacity, or
if it using a `USER_PROVISIONED` value.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | USER_PROVISIONED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataReadCacheConfiguration

FsrmConfiguration

All content copied from https://docs.aws.amazon.com/.
