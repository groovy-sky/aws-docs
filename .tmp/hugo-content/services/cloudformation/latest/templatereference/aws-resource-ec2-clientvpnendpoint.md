This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint

Specifies a Client VPN endpoint. A Client VPN endpoint is the resource you create and
configure to enable and manage client VPN sessions. It is the destination endpoint at which
all client VPN sessions are terminated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::ClientVpnEndpoint",
  "Properties" : {
      "AuthenticationOptions" : [ ClientAuthenticationRequest, ... ],
      "ClientCidrBlock" : String,
      "ClientConnectOptions" : ClientConnectOptions,
      "ClientLoginBannerOptions" : ClientLoginBannerOptions,
      "ClientRouteEnforcementOptions" : ClientRouteEnforcementOptions,
      "ConnectionLogOptions" : ConnectionLogOptions,
      "Description" : String,
      "DisconnectOnSessionTimeout" : Boolean,
      "DnsServers" : [ String, ... ],
      "EndpointIpAddressType" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SelfServicePortal" : String,
      "ServerCertificateArn" : String,
      "SessionTimeoutHours" : Integer,
      "SplitTunnel" : Boolean,
      "TagSpecifications" : [ TagSpecification, ... ],
      "TrafficIpAddressType" : String,
      "TransportProtocol" : String,
      "VpcId" : String,
      "VpnPort" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::EC2::ClientVpnEndpoint
Properties:
  AuthenticationOptions:
    - ClientAuthenticationRequest
  ClientCidrBlock: String
  ClientConnectOptions:
    ClientConnectOptions
  ClientLoginBannerOptions:
    ClientLoginBannerOptions
  ClientRouteEnforcementOptions:
    ClientRouteEnforcementOptions
  ConnectionLogOptions:
    ConnectionLogOptions
  Description: String
  DisconnectOnSessionTimeout: Boolean
  DnsServers:
    - String
  EndpointIpAddressType: String
  SecurityGroupIds:
    - String
  SelfServicePortal: String
  ServerCertificateArn: String
  SessionTimeoutHours: Integer
  SplitTunnel: Boolean
  TagSpecifications:
    - TagSpecification
  TrafficIpAddressType: String
  TransportProtocol: String
  VpcId: String
  VpnPort: Integer

```

## Properties

`AuthenticationOptions`

Information about the authentication method to be used to authenticate clients.

_Required_: Yes

_Type_: Array of [ClientAuthenticationRequest](aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientCidrBlock`

The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. Client CIDR range must have a size of at least /22 and must not be greater than /12.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientConnectOptions`

The options for managing connection authorization for new client connections.

_Required_: No

_Type_: [ClientConnectOptions](aws-properties-ec2-clientvpnendpoint-clientconnectoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientLoginBannerOptions`

Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is
established.

_Required_: No

_Type_: [ClientLoginBannerOptions](aws-properties-ec2-clientvpnendpoint-clientloginbanneroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientRouteEnforcementOptions`

Client route enforcement is a feature of the Client VPN service that helps enforce administrator defined routes on devices connected through the VPN. T
his feature helps improve your security posture by ensuring that network traffic originating from a connected client is not inadvertently sent outside the VPN tunnel.

Client route enforcement works by monitoring the route table of a connected device for routing policy changes to the VPN connection. If the feature detects any VPN routing policy modifications, it will automatically force an update to the route table,
reverting it back to the expected route configurations.

_Required_: No

_Type_: [ClientRouteEnforcementOptions](aws-properties-ec2-clientvpnendpoint-clientrouteenforcementoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionLogOptions`

Information about the client connection logging options.

If you enable client connection logging, data about client connections is sent to a
Cloudwatch Logs log stream. The following information is logged:

- Client connection requests

- Client connection results (successful and unsuccessful)

- Reasons for unsuccessful client connection requests

- Client connection termination time

_Required_: Yes

_Type_: [ConnectionLogOptions](aws-properties-ec2-clientvpnendpoint-connectionlogoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A brief description of the Client VPN endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisconnectOnSessionTimeout`

Indicates whether the client VPN session is disconnected after the maximum `sessionTimeoutHours` is reached. If `true`, users are prompted to reconnect client VPN. If `false`, client VPN attempts to reconnect automatically. The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsServers`

Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can
have up to two DNS servers. If no DNS server is specified, the DNS address configured on the device is used for the DNS server.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointIpAddressType`

The IP address type of the Client VPN endpoint. Possible values are `ipv4` for IPv4 addressing only, `ipv6` for IPv6 addressing only, or `dual-stack ` for both IPv4 and IPv6 addressing.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dual-stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The IDs of one or more security groups to apply to the target network. You must also specify the ID of the VPC that contains the security groups.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfServicePortal`

Specify whether to enable the self-service portal for the Client VPN endpoint.

Default Value: `enabled`

_Required_: No

_Type_: String

_Allowed values_: `enabled | disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerCertificateArn`

The ARN of the server certificate. For more information, see
the [AWS Certificate Manager User Guide](../../../acm/latest/userguide.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionTimeoutHours`

The maximum VPN session duration time in hours.

Valid values: `8 | 10 | 12 | 24`

Default value: `24`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SplitTunnel`

Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.

By default, split-tunnel on a VPN endpoint is disabled.

For information about split-tunnel VPN endpoints, see [Split-tunnel AWS Client VPN endpoint](../../../vpn/latest/clientvpn-admin/split-tunnel-vpn.md) in the
_AWS Client VPN Administrator Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagSpecifications`

The tags to apply to the Client VPN endpoint during creation.

_Required_: No

_Type_: Array of [TagSpecification](aws-properties-ec2-clientvpnendpoint-tagspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrafficIpAddressType`

The IP address type of the Client VPN endpoint. Possible values are either `ipv4` for IPv4 addressing only, `ipv6` for IPv6 addressing only, or `dual-stack` for both IPv4 and IPv6 addressing.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dual-stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransportProtocol`

The transport protocol to be used by the VPN session.

Default value: `udp`

_Required_: No

_Type_: String

_Allowed values_: `tcp | udp`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the VPC to associate with the Client VPN endpoint. If no security group IDs are specified in the request, the default security group for the VPC is applied.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpnPort`

The port number to assign to the Client VPN endpoint for TCP and UDP traffic.

Valid Values: `443` \| `1194`

Default Value: `443`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Client VPN endpoint ID. For example:
`cvpn-endpoint-1234567890abcdef0`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Client VPN endpoint

The following example creates a Client VPN endpoint that uses Active Directory
authentication and assigns client IP addresses from the `10.0.0.0/22` CIDR
range.

#### YAML

```yaml

myClientVpnEndpoint:
  Type: AWS::EC2::ClientVpnEndpoint
  Properties:
    AuthenticationOptions:
    - Type: "directory-service-authentication"
      ActiveDirectory:
        DirectoryId: d-926example
    ClientCidrBlock: "10.0.0.0/22"
    ConnectionLogOptions:
      Enabled: false
    Description: "My Client VPN Endpoint"
    DnsServers:
      - "11.11.0.1"
    SecurityGroupIds:
      - !Ref mySecurityGroup
    ServerCertificateArn: "arn:aws:acm:us-east-1:111122223333:certificate/12345678-1234-1234-1234-123456789012"
    TagSpecifications:
      - ResourceType: "client-vpn-endpoint"
        Tags:
        - Key: "Purpose"
          Value: "Production"
    TransportProtocol: "udp"
```

#### JSON

```json

"myClientVpnEndpoint": {
    "Type": "AWS::EC2::ClientVpnEndpoint",
    "Properties": {
        "AuthenticationOptions": [
            {
                "Type": "directory-service-authentication",
                "ActiveDirectory": {
                    "DirectoryId": "d-926example"
                }
            }
        ],
        "ClientCidrBlock": "10.0.0.0/22",
        "ConnectionLogOptions": {
            "Enabled": false
        },
        "Description": "My Client VPN Endpoint",
        "DnsServers": [
            "11.11.0.1"
        ],
        "SecurityGroupIds": [
            {
                "Ref": "mySecurityGroup"
            }
        ],
        "ServerCertificateArn": "arn:aws:acm:us-east-1:111122223333:certificate/12345678-1234-1234-1234-123456789012",
        "TagSpecifications": [
            {
                "ResourceType": "client-vpn-endpoint",
                "Tags": [
                    {
                        "Key": "Purpose",
                        "Value": "Production"
                    }
                ]
            }
        ],
        "TransportProtocol": "udp"
    }
}
```

## See also

- [Getting Started\
with Client VPN](../../../vpn/latest/clientvpn-admin/cvpn-getting-started.md) in the _AWS Client VPN_
_Administrator Guide_

- [Client VPN\
Endpoints](../../../vpn/latest/clientvpn-admin/cvpn-working-endpoints.md) in the _AWS Client VPN Administrator_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::ClientVpnAuthorizationRule

CertificateAuthenticationRequest

All content copied from https://docs.aws.amazon.com/.
