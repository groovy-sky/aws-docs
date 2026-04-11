# ActiveVpnTunnelStatus

Contains information about the current security configuration of an active VPN tunnel.

## Contents

**ikeVersion**

The version of the Internet Key Exchange (IKE) protocol being used.

Type: String

Required: No

**phase1DHGroup**

The Diffie-Hellman group number being used in Phase 1 IKE negotiations.

Type: Integer

Required: No

**phase1EncryptionAlgorithm**

The encryption algorithm negotiated in Phase 1 IKE negotiations.

Type: String

Required: No

**phase1IntegrityAlgorithm**

The integrity algorithm negotiated in Phase 1 IKE negotiations.

Type: String

Required: No

**phase2DHGroup**

The Diffie-Hellman group number being used in Phase 2 IKE negotiations.

Type: Integer

Required: No

**phase2EncryptionAlgorithm**

The encryption algorithm negotiated in Phase 2 IKE negotiations.

Type: String

Required: No

**phase2IntegrityAlgorithm**

The integrity algorithm negotiated in Phase 2 IKE negotiations.

Type: String

Required: No

**provisioningStatus**

The current provisioning status of the VPN tunnel.

Type: String

Valid Values: `available | pending | failed`

Required: No

**provisioningStatusReason**

The reason for the current provisioning status.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/activevpntunnelstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/activevpntunnelstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/activevpntunnelstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ActiveInstance

AddedPrincipal
