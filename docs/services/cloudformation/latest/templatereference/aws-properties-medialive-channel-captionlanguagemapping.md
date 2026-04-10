This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CaptionLanguageMapping

Maps a captions channel to an ISO 693-2 language code (http://www.loc.gov/standards/iso639-2), with an optional
description.

The parent of this entity is HlsGroupSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaptionChannel" : Integer,
  "LanguageCode" : String,
  "LanguageDescription" : String
}

```

### YAML

```yaml

  CaptionChannel: Integer
  LanguageCode: String
  LanguageDescription: String

```

## Properties

`CaptionChannel`

The closed caption channel being described by this CaptionLanguageMapping. Each channel mapping must have a
unique channel number (maximum of 4).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageCode`

A three-character ISO 639-2 language code (see http://www.loc.gov/standards/iso639-2).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageDescription`

The textual description of language.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptionDestinationSettings

CaptionRectangle

All content copied from https://docs.aws.amazon.com/.
