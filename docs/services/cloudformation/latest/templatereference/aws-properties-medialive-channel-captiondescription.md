This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CaptionDescription

The encoding information for output captions.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Accessibility" : String,
  "CaptionDashRoles" : [ String, ... ],
  "CaptionSelectorName" : String,
  "DestinationSettings" : CaptionDestinationSettings,
  "DvbDashAccessibility" : String,
  "LanguageCode" : String,
  "LanguageDescription" : String,
  "Name" : String
}

```

### YAML

```yaml

  Accessibility: String
  CaptionDashRoles:
    - String
  CaptionSelectorName: String
  DestinationSettings:
    CaptionDestinationSettings
  DvbDashAccessibility: String
  LanguageCode: String
  LanguageDescription: String
  Name: String

```

## Properties

`Accessibility`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionDashRoles`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionSelectorName`

Specifies which input captions selector to use as a captions source when generating output captions. This field
should match a captionSelector name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationSettings`

Additional settings for a captions destination that depend on the destination type.

_Required_: No

_Type_: [CaptionDestinationSettings](aws-properties-medialive-channel-captiondestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DvbDashAccessibility`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageCode`

An ISO 639-2 three-digit code. For more information, see http://www.loc.gov/standards/iso639-2/.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageDescription`

Human-readable information to indicate the captions that are available for players (for example, English or
Spanish).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the captions description. The name is used to associate a captions description with an output. Names
must be unique within a channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BurnInDestinationSettings

CaptionDestinationSettings

All content copied from https://docs.aws.amazon.com/.
