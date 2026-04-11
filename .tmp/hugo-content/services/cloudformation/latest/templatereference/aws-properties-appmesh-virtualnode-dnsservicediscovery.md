This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode DnsServiceDiscovery

An object that represents the DNS service discovery information for your virtual
node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hostname" : String,
  "IpPreference" : String,
  "ResponseType" : String
}

```

### YAML

```yaml

  Hostname: String
  IpPreference: String
  ResponseType: String

```

## Properties

`Hostname`

Specifies the DNS service discovery hostname for the virtual node.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpPreference`

The preferred IP version that this virtual node uses. Setting the IP preference on the
virtual node only overrides the IP preference set for the mesh on this specific
node.

_Required_: No

_Type_: String

_Allowed values_: `IPv6_PREFERRED | IPv4_PREFERRED | IPv4_ONLY | IPv6_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseType`

Specifies the DNS response type for the virtual node.

_Required_: No

_Type_: String

_Allowed values_: `LOADBALANCER | ENDPOINTS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientTlsCertificate

Duration

All content copied from https://docs.aws.amazon.com/.
