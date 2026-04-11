This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase TN3270

Specifies the TN3270 protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExportDataSetNames" : [ String, ... ],
  "Script" : Script
}

```

### YAML

```yaml

  ExportDataSetNames:
    - String
  Script:
    Script

```

## Properties

`ExportDataSetNames`

The data set names of the TN3270 protocol.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Script`

The script of the TN3270 protocol.

_Required_: Yes

_Type_: [Script](aws-properties-apptest-testcase-script.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TestCaseLatestVersion

Next

All content copied from https://docs.aws.amazon.com/.
