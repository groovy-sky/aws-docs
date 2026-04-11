# ClientRouteEnforcementOptions

Client Route Enforcement is a feature of Client VPN that helps enforce administrator defined
routes on devices connected through the VPN. This feature helps improve your security
posture by ensuring that network traffic originating from a connected client is not
inadvertently sent outside the VPN tunnel.

Client Route Enforcement works by monitoring the route table of a connected device for
routing policy changes to the VPN connection. If the feature detects any VPN routing
policy modifications, it will automatically force an update to the route table,
reverting it back to the expected route configurations.

## Contents

**Enforced**

Enable or disable Client Route Enforcement. The state can either be `true`
(enabled) or `false` (disabled). The default is `false`.

Valid values: `true | false`

Default value: `false`

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/clientrouteenforcementoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/clientrouteenforcementoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/clientrouteenforcementoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientLoginBannerResponseOptions

ClientRouteEnforcementResponseOptions
