This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition Tmpfs

The container path, mount options, and size of the `tmpfs` mount.

###### Note

This object isn't applicable to jobs that are running on Fargate resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerPath" : String,
  "MountOptions" : [ String, ... ],
  "Size" : Integer
}

```

### YAML

```yaml

  ContainerPath: String
  MountOptions:
    - String
  Size: Integer

```

## Properties

`ContainerPath`

The absolute file path in the container where the `tmpfs` volume is
mounted.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountOptions`

The list of `tmpfs` volume mount options.

Valid values: " `defaults`" \| " `ro`" \| " `rw`" \|
" `suid`" \| " `nosuid`" \| " `dev`" \| " `nodev`" \|
" `exec`" \| " `noexec`" \| " `sync`" \| " `async`" \|
" `dirsync`" \| " `remount`" \| " `mand`" \| " `nomand`" \|
" `atime`" \| " `noatime`" \| " `diratime`" \|
" `nodiratime`" \| " `bind`" \| " `rbind" | "unbindable" | "runbindable" |
    "private" | "rprivate" | "shared" | "rshared" | "slave" | "rslave" | "relatime`" \|
" `norelatime`" \| " `strictatime`" \| " `nostrictatime`" \|
" `mode`" \| " `uid`" \| " `gid`" \| " `nr_inodes`" \|
" `nr_blocks`" \| " `mpol`"

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The size (in MiB) of the `tmpfs` volume.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TaskContainerProperties

Ulimit

All content copied from https://docs.aws.amazon.com/.
