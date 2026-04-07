# VgwTelemetry

Describes telemetry for a VPN tunnel.

## Contents

**acceptedRouteCount**

The number of accepted routes.

Type: Integer

Required: No

**certificateArn**

The Amazon Resource Name (ARN) of the VPN tunnel endpoint certificate.

Type: String

Required: No

**lastStatusChange**

The date and time of the last change in status. This field is updated when changes in IKE (Phase 1), IPSec (Phase 2), or BGP status are detected.

Type: Timestamp

Required: No

**outsideIpAddress**

The Internet-routable IP address of the virtual private gateway's outside
interface.

Type: String

Required: No

**status**

The status of the VPN tunnel.

Type: String

Valid Values: `UP | DOWN`

Required: No

**statusMessage**

If an error occurs, a description of the error.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vgwtelemetry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vgwtelemetry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vgwtelemetry.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessTrustProviderCondensed

Volume
