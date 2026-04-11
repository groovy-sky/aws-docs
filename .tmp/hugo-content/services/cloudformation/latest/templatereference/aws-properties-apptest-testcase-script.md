This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase Script

Specifies the script.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScriptLocation" : String,
  "Type" : String
}

```

### YAML

```yaml

  ScriptLocation: String
  Type: String

```

## Properties

`ScriptLocation`

The script location of the scripts.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the scripts.

_Required_: Yes

_Type_: String

_Allowed values_: `Selenium`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceAction

SourceDatabaseMetadata

All content copied from https://docs.aws.amazon.com/.
