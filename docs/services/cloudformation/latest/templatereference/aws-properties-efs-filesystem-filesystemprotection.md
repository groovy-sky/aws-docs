This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::FileSystem FileSystemProtection

Describes the protection on the file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReplicationOverwriteProtection" : String
}

```

### YAML

```yaml

  ReplicationOverwriteProtection: String

```

## Properties

`ReplicationOverwriteProtection`

The status of the file system's replication overwrite protection.

- `ENABLED` – The file system cannot be used as the destination file
system in a replication configuration. The file system is writeable. Replication overwrite
protection is `ENABLED` by default.

- `DISABLED` – The file system can be used as the destination file
system in a replication configuration. The file system is read-only and can only be
modified by EFS replication.

- `REPLICATING` – The file system is being used as the destination
file system in a replication configuration. The file system is read-only and is modified
only by EFS replication.

If the replication configuration is deleted, the file system's replication overwrite
protection is re-enabled, the file system becomes writeable.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ElasticFileSystemTag

LifecyclePolicy

All content copied from https://docs.aws.amazon.com/.
