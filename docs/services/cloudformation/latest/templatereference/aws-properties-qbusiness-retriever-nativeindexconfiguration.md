This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Retriever NativeIndexConfiguration

Configuration information for an Amazon Q Business index.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IndexId" : String
}

```

### YAML

```yaml

  IndexId: String

```

## Properties

`IndexId`

The identifier for the Amazon Q Business index.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KendraIndexConfiguration

RetrieverConfiguration

All content copied from https://docs.aws.amazon.com/.
