This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint ClientRouteEnforcementOptions

Client Route Enforcement is a feature of Client VPN that helps enforce administrator defined
routes on devices connected through the VPN. This feature helps improve your security
posture by ensuring that network traffic originating from a connected client is not
inadvertently sent outside the VPN tunnel.

Client Route Enforcement works by monitoring the route table of a connected device for
routing policy changes to the VPN connection. If the feature detects any VPN routing
policy modifications, it will automatically force an update to the route table,
reverting it back to the expected route configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enforced" : Boolean
}

```

### YAML

```yaml

  Enforced: Boolean

```

## Properties

`Enforced`

Enable or disable Client Route Enforcement. The state can either be `true`
(enabled) or `false` (disabled). The default is `false`.

Valid values: `true | false`

Default value: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClientLoginBannerOptions

ConnectionLogOptions
