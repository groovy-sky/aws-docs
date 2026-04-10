This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment OntapUnixFileSystemUser

The FSx for ONTAP UNIX file system user that is used for authorizing all file access requests that are made using the S3 access point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the UNIX user. The name can be up to 256 characters long.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\u0000\u0085\u2028\u2029\r\n]{1,256}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OntapFileSystemIdentity

OntapWindowsFileSystemUser

All content copied from https://docs.aws.amazon.com/.
