This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase SourceDatabaseMetadata

Specifies the source database metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaptureTool" : String,
  "Type" : String
}

```

### YAML

```yaml

  CaptureTool: String
  Type: String

```

## Properties

`CaptureTool`

The capture tool of the source database metadata.

_Required_: Yes

_Type_: String

_Allowed values_: `Precisely | AWS DMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the source database metadata.

_Required_: Yes

_Type_: String

_Allowed values_: `z/OS-DB2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Script

Step

All content copied from https://docs.aws.amazon.com/.
