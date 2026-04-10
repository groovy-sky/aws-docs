This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ProjectFileSystemLocation

Information about a file system created by Amazon Elastic File System (EFS). For more
information, see [What Is\
Amazon Elastic File System?](../../../efs/latest/ug/whatisefs.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identifier" : String,
  "Location" : String,
  "MountOptions" : String,
  "MountPoint" : String,
  "Type" : String
}

```

### YAML

```yaml

  Identifier: String
  Location: String
  MountOptions: String
  MountPoint: String
  Type: String

```

## Properties

`Identifier`

The name used to access a file system created by Amazon EFS. CodeBuild creates an
environment variable by appending the `identifier` in all capital letters to
`CODEBUILD_`. For example, if you specify `my_efs` for
`identifier`, a new environment variable is create named
`CODEBUILD_MY_EFS`.

The `identifier` is used to mount your file system.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

A string that specifies the location of the file system created by Amazon EFS. Its
format is `efs-dns-name:/directory-path`. You can find the DNS name of file
system when you view it in the Amazon EFS console. The directory path is a path to a
directory in the file system that CodeBuild mounts. For example, if the DNS name of a
file system is `fs-abcd1234.efs.us-west-2.amazonaws.com`, and its mount
directory is `my-efs-mount-directory`, then the `location` is
`fs-abcd1234.efs.us-west-2.amazonaws.com:/my-efs-mount-directory`.

The directory path in the format `efs-dns-name:/directory-path` is
optional. If you do not specify a directory path, the location is only the DNS name and
CodeBuild mounts the entire file system.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountOptions`

The mount options for a file system created by Amazon EFS. The default mount options
used by CodeBuild are
`nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2`. For
more information, see [Recommended NFS Mount\
Options](../../../efs/latest/ug/mounting-fs-nfs-mount-settings.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPoint`

The location in the container where you mount the file system.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the file system. The one supported type is `EFS`.

_Required_: Yes

_Type_: String

_Allowed values_: `EFS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectCache

ProjectFleet

All content copied from https://docs.aws.amazon.com/.
