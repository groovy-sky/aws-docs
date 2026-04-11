This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Files::FileSystemPolicy

The `AWS::S3Files::FileSystemPolicy` resource specifies a resource-based
policy for an Amazon S3 Files file system. Use this resource to control access
permissions for the file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Files::FileSystemPolicy",
  "Properties" : {
      "FileSystemId" : String,
      "Policy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::S3Files::FileSystemPolicy
Properties:
  FileSystemId: String
  Policy: Json

```

## Properties

`FileSystemId`

The ID of the S3 Files file system to which the policy applies.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The JSON formatted resource-based policy for the file system.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the file system ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::S3Files::MountTarget

All content copied from https://docs.aws.amazon.com/.
