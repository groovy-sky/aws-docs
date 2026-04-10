This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection VpnTunnelOptionsSpecification

The tunnel options for a single VPN tunnel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DPDTimeoutAction" : String,
  "DPDTimeoutSeconds" : Integer,
  "EnableTunnelLifecycleControl" : Boolean,
  "IKEVersions" : [ IKEVersionsRequestListValue, ... ],
  "LogOptions" : VpnTunnelLogOptionsSpecification,
  "Phase1DHGroupNumbers" : [ Phase1DHGroupNumbersRequestListValue, ... ],
  "Phase1EncryptionAlgorithms" : [ Phase1EncryptionAlgorithmsRequestListValue, ... ],
  "Phase1IntegrityAlgorithms" : [ Phase1IntegrityAlgorithmsRequestListValue, ... ],
  "Phase1LifetimeSeconds" : Integer,
  "Phase2DHGroupNumbers" : [ Phase2DHGroupNumbersRequestListValue, ... ],
  "Phase2EncryptionAlgorithms" : [ Phase2EncryptionAlgorithmsRequestListValue, ... ],
  "Phase2IntegrityAlgorithms" : [ Phase2IntegrityAlgorithmsRequestListValue, ... ],
  "Phase2LifetimeSeconds" : Integer,
  "PreSharedKey" : String,
  "RekeyFuzzPercentage" : Integer,
  "RekeyMarginTimeSeconds" : Integer,
  "ReplayWindowSize" : Integer,
  "StartupAction" : String,
  "TunnelInsideCidr" : String,
  "TunnelInsideIpv6Cidr" : String
}

```

### YAML

```yaml

  DPDTimeoutAction: String
  DPDTimeoutSeconds: Integer
  EnableTunnelLifecycleControl: Boolean
  IKEVersions:
    - IKEVersionsRequestListValue
  LogOptions:
    VpnTunnelLogOptionsSpecification
  Phase1DHGroupNumbers:
    - Phase1DHGroupNumbersRequestListValue
  Phase1EncryptionAlgorithms:
    - Phase1EncryptionAlgorithmsRequestListValue
  Phase1IntegrityAlgorithms:
    - Phase1IntegrityAlgorithmsRequestListValue
  Phase1LifetimeSeconds: Integer
  Phase2DHGroupNumbers:
    - Phase2DHGroupNumbersRequestListValue
  Phase2EncryptionAlgorithms:
    - Phase2EncryptionAlgorithmsRequestListValue
  Phase2IntegrityAlgorithms:
    - Phase2IntegrityAlgorithmsRequestListValue
  Phase2LifetimeSeconds: Integer
  PreSharedKey: String
  RekeyFuzzPercentage: Integer
  RekeyMarginTimeSeconds: Integer
  ReplayWindowSize: Integer
  StartupAction: String
  TunnelInsideCidr: String
  TunnelInsideIpv6Cidr: String

```

## Properties

`DPDTimeoutAction`

The action to take after DPD timeout occurs. Specify `restart` to restart
the IKE initiation. Specify `clear` to end the IKE session.

Valid Values: `clear` \| `none` \| `restart`

Default: `clear`

_Required_: No

_Type_: String

_Allowed values_: `clear | none | restart`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DPDTimeoutSeconds`

The number of seconds after which a DPD timeout occurs.

Constraints: A value greater than or equal to 30.

Default: `30`

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableTunnelLifecycleControl`

Turn on or off tunnel endpoint lifecycle control feature.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IKEVersions`

The IKE versions that are permitted for the VPN tunnel.

Valid values: `ikev1` \| `ikev2`

_Required_: No

_Type_: Array of [IKEVersionsRequestListValue](aws-properties-ec2-vpnconnection-ikeversionsrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogOptions`

Options for logging VPN tunnel activity.

_Required_: No

_Type_: [VpnTunnelLogOptionsSpecification](aws-properties-ec2-vpnconnection-vpntunnellogoptionsspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase1DHGroupNumbers`

One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for
phase 1 IKE negotiations.

Valid values: `2` \| `14` \| `15` \| `16` \|
`17` \| `18` \| `19` \| `20` \|
`21` \| `22` \| `23` \| `24`

_Required_: No

_Type_: Array of [Phase1DHGroupNumbersRequestListValue](aws-properties-ec2-vpnconnection-phase1dhgroupnumbersrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase1EncryptionAlgorithms`

One or more encryption algorithms that are permitted for the VPN tunnel for phase 1
IKE negotiations.

Valid values: `AES128` \| `AES256` \| `AES128-GCM-16` \|
`AES256-GCM-16`

_Required_: No

_Type_: Array of [Phase1EncryptionAlgorithmsRequestListValue](aws-properties-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase1IntegrityAlgorithms`

One or more integrity algorithms that are permitted for the VPN tunnel for phase 1 IKE
negotiations.

Valid values: `SHA1` \| `SHA2-256` \| `SHA2-384` \|
`SHA2-512`

_Required_: No

_Type_: Array of [Phase1IntegrityAlgorithmsRequestListValue](aws-properties-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase1LifetimeSeconds`

The lifetime for phase 1 of the IKE negotiation, in seconds.

Constraints: A value between 900 and 28,800.

Default: `28800`

_Required_: No

_Type_: Integer

_Minimum_: `900`

_Maximum_: `28800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase2DHGroupNumbers`

One or more Diffie-Hellman group numbers that are permitted for the VPN tunnel for
phase 2 IKE negotiations.

Valid values: `2` \| `5` \| `14` \| `15` \|
`16` \| `17` \| `18` \| `19` \|
`20` \| `21` \| `22` \| `23` \|
`24`

_Required_: No

_Type_: Array of [Phase2DHGroupNumbersRequestListValue](aws-properties-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase2EncryptionAlgorithms`

One or more encryption algorithms that are permitted for the VPN tunnel for phase 2
IKE negotiations.

Valid values: `AES128` \| `AES256` \| `AES128-GCM-16` \|
`AES256-GCM-16`

_Required_: No

_Type_: Array of [Phase2EncryptionAlgorithmsRequestListValue](aws-properties-ec2-vpnconnection-phase2encryptionalgorithmsrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase2IntegrityAlgorithms`

One or more integrity algorithms that are permitted for the VPN tunnel for phase 2 IKE
negotiations.

Valid values: `SHA1` \| `SHA2-256` \| `SHA2-384` \|
`SHA2-512`

_Required_: No

_Type_: Array of [Phase2IntegrityAlgorithmsRequestListValue](aws-properties-ec2-vpnconnection-phase2integrityalgorithmsrequestlistvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phase2LifetimeSeconds`

The lifetime for phase 2 of the IKE negotiation, in seconds.

Constraints: A value between 900 and 3,600. The value must be less than the value for
`Phase1LifetimeSeconds`.

Default: `3600`

_Required_: No

_Type_: Integer

_Minimum_: `900`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreSharedKey`

The pre-shared key (PSK) to establish initial authentication between the virtual
private gateway and customer gateway.

Constraints: Allowed characters are alphanumeric characters, periods (.), and
underscores (\_). Must be between 8 and 64 characters in length and cannot start with
zero (0).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RekeyFuzzPercentage`

The percentage of the rekey window (determined by `RekeyMarginTimeSeconds`)
during which the rekey time is randomly selected.

Constraints: A value between 0 and 100.

Default: `100`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RekeyMarginTimeSeconds`

The margin time, in seconds, before the phase 2 lifetime expires, during which the
AWS side of the VPN connection performs an IKE rekey. The exact time
of the rekey is randomly selected based on the value for
`RekeyFuzzPercentage`.

Constraints: A value between 60 and half of `Phase2LifetimeSeconds`.

Default: `270`

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplayWindowSize`

The number of packets in an IKE replay window.

Constraints: A value between 64 and 2048.

Default: `1024`

_Required_: No

_Type_: Integer

_Minimum_: `64`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartupAction`

The action to take when the establishing the tunnel for the VPN connection. By
default, your customer gateway device must initiate the IKE negotiation and bring up the
tunnel. Specify `start` for AWS to initiate the IKE
negotiation.

Valid Values: `add` \| `start`

Default: `add`

_Required_: No

_Type_: String

_Allowed values_: `add | start`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TunnelInsideCidr`

The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be
unique across all VPN connections that use the same virtual private gateway.

Constraints: A size /30 CIDR block from the `169.254.0.0/16` range. The
following CIDR blocks are reserved and cannot be used:

- `169.254.0.0/30`

- `169.254.1.0/30`

- `169.254.2.0/30`

- `169.254.3.0/30`

- `169.254.4.0/30`

- `169.254.5.0/30`

- `169.254.169.252/30`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TunnelInsideIpv6Cidr`

The range of inside IPv6 addresses for the tunnel. Any specified CIDR blocks must be
unique across all VPN connections that use the same transit gateway.

Constraints: A size /126 CIDR block from the local `fd00::/8` range.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpnTunnelLogOptionsSpecification

AWS::EC2::VPNConnectionRoute

All content copied from https://docs.aws.amazon.com/.
