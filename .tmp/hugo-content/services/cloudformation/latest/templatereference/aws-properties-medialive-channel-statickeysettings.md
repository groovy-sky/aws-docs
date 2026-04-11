This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel StaticKeySettings

The static key settings.

The parent of this entity is KeyProviderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyProviderServer" : InputLocation,
  "StaticKeyValue" : String
}

```

### YAML

```yaml

  KeyProviderServer:
    InputLocation
  StaticKeyValue: String

```

## Properties

`KeyProviderServer`

The URL of the license server that is used for protecting content.

_Required_: No

_Type_: [InputLocation](aws-properties-medialive-channel-inputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticKeyValue`

The static key value as a 32 character hexadecimal string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StandardHlsSettings

TeletextSourceSettings

All content copied from https://docs.aws.amazon.com/.
