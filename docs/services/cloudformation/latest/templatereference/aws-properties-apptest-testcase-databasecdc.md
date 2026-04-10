This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase DatabaseCDC

Defines the Change Data Capture (CDC) of the database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceMetadata" : SourceDatabaseMetadata,
  "TargetMetadata" : TargetDatabaseMetadata
}

```

### YAML

```yaml

  SourceMetadata:
    SourceDatabaseMetadata
  TargetMetadata:
    TargetDatabaseMetadata

```

## Properties

`SourceMetadata`

The source metadata of the database CDC.

_Required_: Yes

_Type_: [SourceDatabaseMetadata](aws-properties-apptest-testcase-sourcedatabasemetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetMetadata`

The target metadata of the database CDC.

_Required_: Yes

_Type_: [TargetDatabaseMetadata](aws-properties-apptest-testcase-targetdatabasemetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompareAction

DataSet

All content copied from https://docs.aws.amazon.com/.
