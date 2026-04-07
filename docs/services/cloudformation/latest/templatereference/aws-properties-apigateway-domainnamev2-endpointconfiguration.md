This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DomainNameV2 EndpointConfiguration

The endpoint configuration to indicate the types of endpoints an API (RestApi) or its custom domain name (DomainName) has and the IP address types that can invoke it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IpAddressType" : String,
  "Types" : [ String, ... ]
}

```

### YAML

```yaml

  IpAddressType: String
  Types:
    - String

```

## Properties

`IpAddressType`

The IP address types that can invoke an API (RestApi) or a DomainName. Use `ipv4` to allow only IPv4 addresses to
invoke an API or DomainName, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke an API or a DomainName. For the
`PRIVATE` endpoint type, only `dualstack` is supported.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | dualstack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Types`

A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is `"EDGE"`. For a regional API and its custom domain name, the endpoint type is `REGIONAL`. For a private API, the endpoint type is `PRIVATE`.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::DomainNameV2

Tag
