This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem UserAndGroupQuotas

Used to configure quotas that define how much storage a user or group can use on an
FSx for OpenZFS volume. For more information, see
[Volume properties](../../../fsx/latest/openzfsguide/managing-volumes.md#volume-properties)
in the FSx for OpenZFS User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : Integer,
  "StorageCapacityQuotaGiB" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  Id: Integer
  StorageCapacityQuotaGiB: Integer
  Type: String

```

## Properties

`Id`

The ID of the user or group that the quota applies to.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageCapacityQuotaGiB`

The user or group's storage quota, in gibibytes (GiB).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

Specifies whether the quota applies to a user or group.

_Required_: No

_Type_: String

_Allowed values_: `USER | GROUP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WindowsConfiguration

All content copied from https://docs.aws.amazon.com/.
