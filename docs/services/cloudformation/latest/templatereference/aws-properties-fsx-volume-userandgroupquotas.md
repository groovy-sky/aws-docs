This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume UserAndGroupQuotas

Configures how much storage users and groups can use on the volume.

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

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageCapacityQuotaGiB`

The user or group's storage quota, in gibibytes (GiB).

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies whether the quota applies to a user or group.

_Required_: Yes

_Type_: String

_Allowed values_: `USER | GROUP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TieringPolicy

Next

All content copied from https://docs.aws.amazon.com/.
