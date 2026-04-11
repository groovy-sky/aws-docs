This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Domain DomainEntry

Describes a domain recordset entry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "IsAlias" : Boolean,
  "Name" : String,
  "Target" : String,
  "Type" : String
}

```

### YAML

```yaml

  Id: String
  IsAlias: Boolean
  Name: String
  Target: String
  Type: String

```

## Properties

`Id`

The ID of the domain recordset entry.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsAlias`

When `true`, specifies whether the domain entry is an alias used by the Lightsail load balancer, Lightsail container service, Lightsail content delivery network (CDN) distribution, or another AWS
resource. You can include an alias (A type) record in your request, which points to the DNS
name of a load balancer, container service, CDN distribution, or other AWS
resource and routes traffic to that resource.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the domain.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The target IP address ( `192.0.2.0`), or AWS name server
( `ns-111.awsdns-22.com.`).

For Lightsail load balancers, the value looks like
`ab1234c56789c6b86aba6fb203d443bc-123456789.us-east-2.elb.amazonaws.com`. For
Lightsail distributions, the value looks like `exampled1182ne.cloudfront.net`.
For Lightsail container services, the value looks like
`container-service-1.example23scljs.us-west-2.cs.amazonlightsail.com`. Be sure to
also set `isAlias` to `true` when setting up an A record for a
Lightsail load balancer, distribution, or container service.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of domain entry, such as address for IPv4 (A), address for IPv6 (AAAA), canonical
name (CNAME), mail exchanger (MX), name server (NS), start of authority (SOA), service locator
(SRV), or text (TXT).

The following domain entry types can be used:

- `A`

- `AAAA`

- `CNAME`

- `MX`

- `NS`

- `SOA`

- `SRV`

- `TXT`

_Required_: Yes

_Type_: String

_Allowed values_: `A | AAAA | CNAME | MX | NS | SOA | SRV | TXT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lightsail::Domain

Location

All content copied from https://docs.aws.amazon.com/.
