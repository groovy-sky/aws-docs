This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow EfsInputFileLocation

Specifies the Amazon EFS identifier and the path for the file being used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemId" : String,
  "Path" : String
}

```

### YAML

```yaml

  FileSystemId: String
  Path: String

```

## Properties

`FileSystemId`

The identifier of the file system, assigned by Amazon EFS.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:elasticfilesystem:[0-9a-z-:]+:(access-point/fsap|file-system/fs)-[0-9a-f]{8,40}|fs(ap)?-[0-9a-f]{8,40})$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

The pathname for the folder being used by a workflow.

_Required_: No

_Type_: String

_Pattern_: `^[^\x00]+$`

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteStepDetails

InputFileLocation

All content copied from https://docs.aws.amazon.com/.
