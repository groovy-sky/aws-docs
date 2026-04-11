This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task Options

Represents the options that are available to control the behavior of a [StartTaskExecution](../../../datasync/latest/userguide/api-starttaskexecution.md) operation. This behavior includes preserving metadata,
such as user ID (UID), group ID (GID), and file permissions; overwriting files in the
destination; data integrity verification; and so on.

A task has a set of default options associated with it. If you don't specify an
option in [StartTaskExecution](../../../datasync/latest/userguide/api-starttaskexecution.md), the default value is used. You can override the default
options on each task execution by specifying an overriding `Options` value to
[StartTaskExecution](../../../datasync/latest/userguide/api-starttaskexecution.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Atime" : String,
  "BytesPerSecond" : Integer,
  "Gid" : String,
  "LogLevel" : String,
  "Mtime" : String,
  "ObjectTags" : String,
  "OverwriteMode" : String,
  "PosixPermissions" : String,
  "PreserveDeletedFiles" : String,
  "PreserveDevices" : String,
  "SecurityDescriptorCopyFlags" : String,
  "TaskQueueing" : String,
  "TransferMode" : String,
  "Uid" : String,
  "VerifyMode" : String
}

```

### YAML

```yaml

  Atime: String
  BytesPerSecond: Integer
  Gid: String
  LogLevel: String
  Mtime: String
  ObjectTags: String
  OverwriteMode: String
  PosixPermissions: String
  PreserveDeletedFiles: String
  PreserveDevices: String
  SecurityDescriptorCopyFlags: String
  TaskQueueing: String
  TransferMode: String
  Uid: String
  VerifyMode: String

```

## Properties

`Atime`

A file metadata value that shows the last time that a file was accessed (that is,
when the file was read or written to). If you set `Atime` to
`BEST_EFFORT`, AWS DataSync attempts to preserve the original
`Atime` attribute on all source files (that is, the version before the
PREPARING phase). However, `Atime`'s behavior is not fully standard across
platforms, so AWS DataSync can only do this on a best-effort basis.

Default value: `BEST_EFFORT`

`BEST_EFFORT`: Attempt to preserve the per-file `Atime` value
(recommended).

`NONE`: Ignore `Atime`.

###### Note

If `Atime` is set to `BEST_EFFORT`, `Mtime` must
be set to `PRESERVE`.

If `Atime` is set to `NONE`, `Mtime` must also be
`NONE`.

_Required_: No

_Type_: String

_Allowed values_: `NONE | BEST_EFFORT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BytesPerSecond`

A value that limits the bandwidth used by AWS DataSync. For example, if
you want AWS DataSync to use a maximum of 1 MB, set this value to
`1048576` (=1024\*1024).

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gid`

The group ID (GID) of the file's owners.

Default value:
`INT_VALUE`

`INT_VALUE`: Preserve the integer value of the user ID (UID) and group ID
(GID) (recommended).

`NAME`: Currently not supported.

`NONE`: Ignore the UID and GID.

_Required_: No

_Type_: String

_Allowed values_: `NONE | INT_VALUE | NAME | BOTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

Specifies the type of logs that DataSync publishes to a Amazon CloudWatch Logs
log group. To specify the log group, see [CloudWatchLogGroupArn](../../../datasync/latest/userguide/api-createtask.md#DataSync-CreateTask-request-CloudWatchLogGroupArn).

- `BASIC` \- Publishes logs with only basic information (such as transfer
errors).

- `TRANSFER` \- Publishes logs for all files or objects that your DataSync task transfers and performs data-integrity checks on.

- `OFF` \- No logs are published.

_Required_: No

_Type_: String

_Allowed values_: `OFF | BASIC | TRANSFER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mtime`

A value that indicates the last time that a file was modified (that is, a file was
written to) before the PREPARING phase. This option is required for cases when you need
to run the same task more than one time.

Default value: `PRESERVE`

`PRESERVE`: Preserve original `Mtime` (recommended)

`NONE`: Ignore `Mtime`.

###### Note

If `Mtime` is set to `PRESERVE`, `Atime` must be
set to `BEST_EFFORT`.

If `Mtime` is set to `NONE`, `Atime` must also be
set to `NONE`.

_Required_: No

_Type_: String

_Allowed values_: `NONE | PRESERVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectTags`

Specifies whether you want DataSync to `PRESERVE` object tags
(default behavior) when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the `NONE` value.

_Required_: No

_Type_: String

_Allowed values_: `PRESERVE | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverwriteMode`

Specifies whether DataSync should modify or preserve data at the destination
location.

- `ALWAYS` (default) - DataSync modifies data in the destination
location when source data (including metadata) has changed.

If DataSync overwrites objects, you might incur additional charges for
certain Amazon S3 storage classes (for example, for retrieval or early deletion).
For more information, see [Storage\
class considerations with Amazon S3 transfers](../../../datasync/latest/userguide/create-s3-location.md#using-storage-classes).

- `NEVER` \- DataSync doesn't overwrite data in the destination
location even if the source data has changed. You can use this option to protect against
overwriting changes made to files or objects in the destination.

_Required_: No

_Type_: String

_Allowed values_: `ALWAYS | NEVER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PosixPermissions`

A value that determines which users or groups can access a file for a specific
purpose, such as reading, writing, or execution of the file. This option should be set
only for Network File System (NFS), Amazon EFS, and Amazon S3 locations. For more
information about what metadata is copied by DataSync, see [Metadata Copied\
by DataSync](../../../datasync/latest/userguide/special-files.md#metadata-copied).

Default value: `PRESERVE`

`PRESERVE`: Preserve POSIX-style permissions (recommended).

`NONE`: Ignore permissions.

###### Note

AWS DataSync can preserve extant permissions of a source
location.

_Required_: No

_Type_: String

_Allowed values_: `NONE | PRESERVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreserveDeletedFiles`

A value that specifies whether files in the destination that don't exist in the
source file system are preserved. This option can affect your storage costs. If your
task deletes objects, you might incur minimum storage duration charges for certain
storage classes. For detailed information, see [Considerations when working with Amazon S3 storage classes in DataSync](../../../datasync/latest/userguide/create-s3-location.md#using-storage-classes) in
the _AWS DataSync User Guide_.

Default value: `PRESERVE`

`PRESERVE`: Ignore destination files that aren't present in the source
(recommended).

`REMOVE`: Delete destination files that aren't present in the
source.

_Required_: No

_Type_: String

_Allowed values_: `PRESERVE | REMOVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreserveDevices`

A value that determines whether AWS DataSync should preserve the metadata
of block and character devices in the source file system, and re-create the files with
that device name and metadata on the destination. DataSync does not copy the contents of
such devices, only the name and metadata.

###### Note

AWS DataSync can't sync the actual contents of such devices, because
they are nonterminal and don't return an end-of-file (EOF) marker.

Default value: `NONE`

`NONE`: Ignore special devices (recommended).

`PRESERVE`: Preserve character and block device metadata. This option
isn't currently supported for Amazon EFS.

_Required_: No

_Type_: String

_Allowed values_: `NONE | PRESERVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityDescriptorCopyFlags`

A value that determines which components of the SMB security descriptor are copied
from source to destination objects.

This value is only used for transfers between SMB and Amazon FSx for Windows File
Server locations, or between two Amazon FSx for Windows File Server locations. For more
information about how DataSync handles metadata, see [How DataSync Handles Metadata and\
Special Files](../../../datasync/latest/userguide/special-files.md).

Default value: `OWNER_DACL`

`OWNER_DACL`: For each copied object, DataSync copies the following
metadata:

- Object owner.

- NTFS discretionary access control lists (DACLs), which determine whether to
grant access to an object.

When you use option, DataSync does NOT copy the NTFS system access control lists
(SACLs), which are used by administrators to log attempts to access a secured
object.

`OWNER_DACL_SACL`: For each copied object, DataSync copies the following
metadata:

- Object owner.

- NTFS discretionary access control lists (DACLs), which determine whether to
grant access to an object.

- NTFS system access control lists (SACLs), which are used by administrators to
log attempts to access a secured object.

Copying SACLs requires granting additional permissions to the Windows user that
DataSync uses to access your SMB location. For information about choosing a user that
ensures sufficient permissions to files, folders, and metadata, see [user](../../../datasync/latest/userguide/create-smb-location.md#SMBuser).

`NONE`: None of the SMB security descriptor components are copied.
Destination objects are owned by the user that was provided for accessing the
destination location. DACLs and SACLs are set based on the destination server’s
configuration.

_Required_: No

_Type_: String

_Allowed values_: `NONE | OWNER_DACL | OWNER_DACL_SACL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskQueueing`

Specifies whether your transfer tasks should be put into a queue during certain scenarios
when [running multiple\
tasks](../../../datasync/latest/userguide/run-task.md#running-multiple-tasks). This is `ENABLED` by default.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransferMode`

A value that determines whether DataSync transfers only the data and metadata that
differ between the source and the destination location, or whether DataSync transfers
all the content from the source, without comparing it to the destination location.

`CHANGED`: DataSync copies only data or metadata that is new or different
from the source location to the destination location.

`ALL`: DataSync copies all source location content to the destination,
without comparing it to existing content on the destination.

_Required_: No

_Type_: String

_Allowed values_: `CHANGED | ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uid`

The user ID (UID) of the file's owner.

Default value:
`INT_VALUE`

`INT_VALUE`: Preserve the integer value of the UID and group ID (GID)
(recommended).

`NAME`: Currently not supported

`NONE`: Ignore the UID and GID.

_Required_: No

_Type_: String

_Allowed values_: `NONE | INT_VALUE | NAME | BOTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifyMode`

A value that determines whether a data integrity verification is performed at the
end of a task execution after all data and metadata have been transferred. For more
information, see [Configure task settings](../../../datasync/latest/userguide/create-task.md).

Default value: `POINT_IN_TIME_CONSISTENT`

`ONLY_FILES_TRANSFERRED` (recommended): Perform verification only on
files that were transferred.

`POINT_IN_TIME_CONSISTENT`: Scan the entire source and entire
destination at the end of the transfer to verify that the source and destination are
fully synchronized. This option isn't supported when transferring to S3 Glacier or S3
Glacier Deep Archive storage classes.

`NONE`: No additional verification is done at the end of the transfer,
but all data transmissions are integrity-checked with checksum verification during the
transfer.

_Required_: No

_Type_: String

_Allowed values_: `POINT_IN_TIME_CONSISTENT | ONLY_FILES_TRANSFERRED | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManifestConfigSourceS3

Overrides

All content copied from https://docs.aws.amazon.com/.
