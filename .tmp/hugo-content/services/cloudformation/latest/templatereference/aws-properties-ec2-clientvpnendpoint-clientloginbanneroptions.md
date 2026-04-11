This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint ClientLoginBannerOptions

Options for enabling a customizable text banner that will be displayed on
AWS provided clients when a VPN session is established.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BannerText" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  BannerText: String
  Enabled: Boolean

```

## Properties

`BannerText`

Customizable text that will be displayed in a banner on AWS provided
clients when a VPN session is established. UTF-8 encoded characters only. Maximum of
1400 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Enable or disable a customizable text banner that will be displayed on
AWS provided clients when a VPN session is established.

Valid values: `true | false`

Default value: `false`

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientConnectOptions

ClientRouteEnforcementOptions

All content copied from https://docs.aws.amazon.com/.
