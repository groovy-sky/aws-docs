This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::Channel SlateSource

Slate VOD source configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceLocationName" : String,
  "VodSourceName" : String
}

```

### YAML

```yaml

  SourceLocationName: String
  VodSourceName: String

```

## Properties

`SourceLocationName`

The name of the source location where the slate VOD source is stored.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VodSourceName`

The slate VOD source name. The VOD source must already exist in a source location before it can be used for slate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestOutputItem

Tag

All content copied from https://docs.aws.amazon.com/.
