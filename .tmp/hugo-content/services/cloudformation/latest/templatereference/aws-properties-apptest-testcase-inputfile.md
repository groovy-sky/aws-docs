This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase InputFile

Specifies the input file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileMetadata" : FileMetadata,
  "SourceLocation" : String,
  "TargetLocation" : String
}

```

### YAML

```yaml

  FileMetadata:
    FileMetadata
  SourceLocation: String
  TargetLocation: String

```

## Properties

`FileMetadata`

The file metadata of the input file.

_Required_: Yes

_Type_: [FileMetadata](aws-properties-apptest-testcase-filemetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLocation`

The source location of the input file.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetLocation`

The target location of the input file.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Input

M2ManagedActionProperties

All content copied from https://docs.aws.amazon.com/.
