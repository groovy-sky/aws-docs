This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Service UrlEndpointInput

The configuration for the URL endpoint type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthUrl" : String,
  "Url" : String
}

```

### YAML

```yaml

  HealthUrl: String
  Url: String

```

## Properties

`HealthUrl`

The health check URL of the URL endpoint type. If the URL is a public endpoint, the
`HealthUrl` must also be a public endpoint. If the URL is a private endpoint
inside a virtual private cloud (VPC), the health URL must also be a private endpoint, and the
host must be the same as the URL.

_Required_: No

_Type_: String

_Pattern_: `^https?://[-a-zA-Z0-9+\x38@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\x38@#/%=~_|]$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Url`

The URL to route traffic to. The URL must be an [rfc3986-formatted URL](https://datatracker.ietf.org/doc/html/rfc3986). If the
host is a domain name, the name must be resolvable over the public internet. If the scheme is
`https`, the top level domain of the host must be listed in the [IANA root zone database](https://www.iana.org/domains/root/db).

_Required_: Yes

_Type_: String

_Pattern_: `^https?://[-a-zA-Z0-9+\x38@#/%?=~_|!:,.;]*[-a-zA-Z0-9+\x38@#/%=~_|]$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
