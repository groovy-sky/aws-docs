This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase Batch

Defines a batch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchJobName" : String,
  "BatchJobParameters" : {Key: Value, ...},
  "ExportDataSetNames" : [ String, ... ]
}

```

### YAML

```yaml

  BatchJobName: String
  BatchJobParameters:
    Key: Value
  ExportDataSetNames:
    - String

```

## Properties

`BatchJobName`

The job name of the batch.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BatchJobParameters`

The batch job parameters of the batch.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportDataSetNames`

The export data set names of the batch.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppTest::TestCase

CloudFormationAction

All content copied from https://docs.aws.amazon.com/.
