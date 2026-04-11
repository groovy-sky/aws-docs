This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase CompareAction

Compares the action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Input" : Input,
  "Output" : Output
}

```

### YAML

```yaml

  Input:
    Input
  Output:
    Output

```

## Properties

`Input`

The input of the compare action.

_Required_: Yes

_Type_: [Input](aws-properties-apptest-testcase-input.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Output`

The output of the compare action.

_Required_: No

_Type_: [Output](aws-properties-apptest-testcase-output.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudFormationAction

DatabaseCDC

All content copied from https://docs.aws.amazon.com/.
