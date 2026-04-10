This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioHlsRenditionSelection

Selector for HLS audio rendition.

The parent of this entity is AudioSelectorSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupId" : String,
  "Name" : String
}

```

### YAML

```yaml

  GroupId: String
  Name: String

```

## Properties

`GroupId`

Specifies the GROUP-ID in the #EXT-X-MEDIA tag of the target HLS audio rendition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Specifies the NAME in the #EXT-X-MEDIA tag of the target HLS audio rendition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioDolbyEDecode

AudioLanguageSelection

All content copied from https://docs.aws.amazon.com/.
