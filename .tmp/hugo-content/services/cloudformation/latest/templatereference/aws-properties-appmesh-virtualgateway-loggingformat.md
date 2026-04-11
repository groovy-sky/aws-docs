This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway LoggingFormat

An object that represents the format for the logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Json" : [ JsonFormatRef, ... ],
  "Text" : String
}

```

### YAML

```yaml

  Json:
    - JsonFormatRef
  Text: String

```

## Properties

`Json`

The logging format for JSON.

_Required_: No

_Type_: Array of [JsonFormatRef](aws-properties-appmesh-virtualgateway-jsonformatref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Text`

The logging format for text.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JsonFormatRef

SubjectAlternativeNameMatchers

All content copied from https://docs.aws.amazon.com/.
