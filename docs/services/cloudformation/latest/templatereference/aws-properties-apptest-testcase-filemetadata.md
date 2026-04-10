This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase FileMetadata

Specifies a file metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseCDC" : DatabaseCDC,
  "DataSets" : [ DataSet, ... ]
}

```

### YAML

```yaml

  DatabaseCDC:
    DatabaseCDC
  DataSets:
    - DataSet

```

## Properties

`DatabaseCDC`

The database CDC of the file metadata.

_Required_: No

_Type_: [DatabaseCDC](aws-properties-apptest-testcase-databasecdc.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSets`

The data sets of the file metadata.

_Required_: No

_Type_: Array of [DataSet](aws-properties-apptest-testcase-dataset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSet

Input

All content copied from https://docs.aws.amazon.com/.
