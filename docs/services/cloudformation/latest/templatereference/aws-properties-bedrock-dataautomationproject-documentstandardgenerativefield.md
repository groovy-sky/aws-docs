This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject DocumentStandardGenerativeField

Settings for generating descriptions of documents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String
}

```

### YAML

```yaml

  State: String

```

## Properties

`State`

Whether generating descriptions is enabled for documents.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentStandardExtraction

DocumentStandardOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
