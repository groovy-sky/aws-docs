This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel NetworkInputSettings

Information about how to connect to the upstream system.

The parent of this entity is InputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HlsInputSettings" : HlsInputSettings,
  "MulticastInputSettings" : MulticastInputSettings,
  "ServerValidation" : String
}

```

### YAML

```yaml

  HlsInputSettings:
    HlsInputSettings
  MulticastInputSettings:
    MulticastInputSettings
  ServerValidation: String

```

## Properties

`HlsInputSettings`

Information about how to connect to the upstream system.

_Required_: No

_Type_: [HlsInputSettings](aws-properties-medialive-channel-hlsinputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MulticastInputSettings`

Property description not available.

_Required_: No

_Type_: [MulticastInputSettings](aws-properties-medialive-channel-multicastinputsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerValidation`

Checks HTTPS server certificates. When set to checkCryptographyOnly, cryptography in the certificate is checked,
but not the server's name. Certain subdomains (notably S3 buckets that use dots in the bucket name) don't strictly
match the corresponding certificate's wildcard pattern and would otherwise cause the channel to error. This setting
is ignored for protocols that do not use HTTPS.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexProgramChannelDestinationSettings

NielsenCBET

All content copied from https://docs.aws.amazon.com/.
