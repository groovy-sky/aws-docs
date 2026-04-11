This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnAuthorizationRule

Specifies an ingress authorization rule to add to a Client VPN endpoint. Ingress
authorization rules act as firewall rules that grant access to networks. You must configure
ingress authorization rules to enable clients to access resources in AWS
or on-premises networks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::ClientVpnAuthorizationRule",
  "Properties" : {
      "AccessGroupId" : String,
      "AuthorizeAllGroups" : Boolean,
      "ClientVpnEndpointId" : String,
      "Description" : String,
      "TargetNetworkCidr" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::ClientVpnAuthorizationRule
Properties:
  AccessGroupId: String
  AuthorizeAllGroups: Boolean
  ClientVpnEndpointId: String
  Description: String
  TargetNetworkCidr: String

```

## Properties

`AccessGroupId`

The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group. Required if `AuthorizeAllGroups` is `false` or not specified.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizeAllGroups`

Indicates whether to grant access to all clients. Specify `true` to grant all
clients who successfully establish a VPN connection access to the network. Must be set
to `true` if `AccessGroupId` is not specified.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientVpnEndpointId`

The ID of the Client VPN endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A brief description of the authorization rule.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetNetworkCidr`

The IPv4 address range, in CIDR notation, of the network for which access is being authorized.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Add an authorization rule to a client VPN endpoint

The following example adds an authorization rule that grants all users access to
the internet.

#### YAML

```yaml

myAuthRule:
  Type: "AWS::EC2::ClientVpnAuthorizationRule"
  Properties:
    ClientVpnEndpointId:
      Ref: myClientVpnEndpoint
    AuthorizeAllGroups: true
    TargetNetworkCidr: "0.0.0.0/0"
    Description: "myAuthRule"
```

#### JSON

```json

"myAuthRule": {
    "Type": "AWS::EC2::ClientVpnAuthorizationRule",
    "Properties": {
        "ClientVpnEndpointId": {
            "Ref": "myClientVpnEndpoint"
        },
        "AuthorizeAllGroups": true,
        "TargetNetworkCidr": "0.0.0.0/0",
        "Description": "myAuthRule"
    }
}
```

## See also

- [Getting Started with\
Client VPN](../../../vpn/latest/clientvpn-admin/cvpn-getting-started.md) in the _AWS Client VPN Administrator_
_Guide_

- [Authorization\
Rules](../../../vpn/latest/clientvpn-admin/cvpn-working-rules.md) in the _AWS Client VPN Administrator_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::ClientVpnEndpoint

All content copied from https://docs.aws.amazon.com/.
