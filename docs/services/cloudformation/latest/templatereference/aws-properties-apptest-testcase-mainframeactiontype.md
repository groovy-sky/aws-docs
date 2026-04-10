This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase MainframeActionType

Specifies the mainframe action type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Batch" : Batch,
  "Tn3270" : TN3270
}

```

### YAML

```yaml

  Batch:
    Batch
  Tn3270:
    TN3270

```

## Properties

`Batch`

The batch of the mainframe action type.

_Required_: No

_Type_: [Batch](aws-properties-apptest-testcase-batch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tn3270`

The tn3270 port of the mainframe action type.

_Required_: No

_Type_: [TN3270](aws-properties-apptest-testcase-tn3270.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MainframeActionProperties

Output

All content copied from https://docs.aws.amazon.com/.
