This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::FileSystem LifecyclePolicy

Describes a policy used by Lifecycle management that
specifies when to transition files into and out of the EFS storage classes. For more information,
see [Managing file system storage](../../../efs/latest/ug/lifecycle-management-efs.md).

###### Note

- Each `LifecyclePolicy` object can have only a single transition. This means that in a
request body, `LifecyclePolicies` must be structured as an array of
`LifecyclePolicy` objects, one object for each transition,
`TransitionToIA`, `TransitionToArchive`, `TransitionToPrimaryStorageClass`.

- See the AWS::EFS::FileSystem examples for the correct `LifecyclePolicy` structure. Do not
use the syntax shown on this page.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TransitionToArchive" : String,
  "TransitionToIA" : String,
  "TransitionToPrimaryStorageClass" : String
}

```

### YAML

```yaml

  TransitionToArchive: String
  TransitionToIA: String
  TransitionToPrimaryStorageClass: String

```

## Properties

`TransitionToArchive`

The number of days after files were last accessed in primary storage (the
Standard storage class) at which to move them to Archive
storage. Metadata operations such as listing the contents of a directory don't count as
file access events.

_Required_: No

_Type_: String

_Allowed values_: `AFTER_1_DAY | AFTER_7_DAYS | AFTER_14_DAYS | AFTER_30_DAYS | AFTER_60_DAYS | AFTER_90_DAYS | AFTER_180_DAYS | AFTER_270_DAYS | AFTER_365_DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionToIA`

The number of days after files were last accessed in primary storage (the
Standard storage class) at which to move them to Infrequent Access
(IA) storage. Metadata operations such as listing the contents of a directory
don't count as file access events.

_Required_: No

_Type_: String

_Allowed values_: `AFTER_7_DAYS | AFTER_14_DAYS | AFTER_30_DAYS | AFTER_60_DAYS | AFTER_90_DAYS | AFTER_1_DAY | AFTER_180_DAYS | AFTER_270_DAYS | AFTER_365_DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionToPrimaryStorageClass`

Whether to move files back to primary (Standard) storage after they are
accessed in IA or Archive storage. Metadata operations such as
listing the contents of a directory don't count as file access events.

_Required_: No

_Type_: String

_Allowed values_: `AFTER_1_ACCESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FileSystemProtection

ReplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
