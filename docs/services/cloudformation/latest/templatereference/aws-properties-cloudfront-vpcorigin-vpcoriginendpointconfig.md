This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::VpcOrigin VpcOriginEndpointConfig

An Amazon CloudFront VPC origin endpoint configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "HTTPPort" : Integer,
  "HTTPSPort" : Integer,
  "Name" : String,
  "OriginProtocolPolicy" : String,
  "OriginSSLProtocols" : [ String, ... ]
}

```

### YAML

```yaml

  Arn: String
  HTTPPort: Integer
  HTTPSPort: Integer
  Name: String
  OriginProtocolPolicy: String
  OriginSSLProtocols:
    - String

```

## Properties

`Arn`

The ARN of the CloudFront VPC origin endpoint configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTPPort`

The HTTP port for the CloudFront VPC origin endpoint configuration. The default value is `80`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTPSPort`

The HTTPS port of the CloudFront VPC origin endpoint configuration. The default value is `443`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the CloudFront VPC origin endpoint configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginProtocolPolicy`

The origin protocol policy for the CloudFront VPC origin endpoint configuration.

_Required_: No

_Type_: String

_Allowed values_: `http-only | match-viewer | https-only`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginSSLProtocols`

Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin
over HTTPS. Valid values include `SSLv3`, `TLSv1`,
`TLSv1.1`, and `TLSv1.2`.

For more information, see [Minimum Origin SSL Protocol](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#DownloadDistValuesOriginSSLProtocols) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
