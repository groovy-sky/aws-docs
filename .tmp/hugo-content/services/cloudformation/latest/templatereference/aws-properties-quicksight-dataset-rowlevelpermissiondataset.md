This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet RowLevelPermissionDataSet

Information about a dataset that contains permissions for row-level security (RLS).
The permissions dataset maps fields to users or groups. For more information, see
[Using Row-Level Security (RLS) to Restrict Access to a Dataset](../../../quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.md) in the _Quick Sight User_
_Guide_.

The option to deny permissions by setting `PermissionPolicy` to `DENY_ACCESS` is
not supported for new RLS datasets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "FormatVersion" : String,
  "Namespace" : String,
  "PermissionPolicy" : String,
  "Status" : String
}

```

### YAML

```yaml

  Arn: String
  FormatVersion: String
  Namespace: String
  PermissionPolicy: String
  Status: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the dataset that contains permissions for RLS.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatVersion`

The user or group rules associated with the dataset that contains permissions for RLS.

By default, `FormatVersion` is `VERSION_1`. When `FormatVersion` is `VERSION_1`, `UserName` and `GroupName` are required. When `FormatVersion` is `VERSION_2`, `UserARN` and `GroupARN` are required, and `Namespace` must not exist.

_Required_: No

_Type_: String

_Allowed values_: `VERSION_1 | VERSION_2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace associated with the dataset that contains permissions for RLS.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]*$`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionPolicy`

The type of permissions to use when interpreting the permissions for RLS. `DENY_ACCESS`
is included for backward compatibility only.

_Required_: Yes

_Type_: String

_Allowed values_: `GRANT_ACCESS | DENY_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the row-level security permission dataset. If enabled, the status is `ENABLED`. If disabled, the status is `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RowLevelPermissionConfiguration

RowLevelPermissionTagConfiguration

All content copied from https://docs.aws.amazon.com/.
