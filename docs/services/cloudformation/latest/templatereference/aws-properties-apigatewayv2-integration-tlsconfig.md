This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Integration TlsConfig

The `TlsConfig` property specifies the TLS configuration for a private
integration. If you specify a TLS configuration, private integration traffic uses the
HTTPS protocol. Supported only for HTTP APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServerNameToVerify" : String
}

```

### YAML

```yaml

  ServerNameToVerify: String

```

## Properties

`ServerNameToVerify`

If you specify a server name, API Gateway uses it to verify the hostname on
the integration's certificate. The server name is also included in the TLS
handshake to support Server Name Indication (SNI) or virtual hosting.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResponseParameterMap

AWS::ApiGatewayV2::IntegrationResponse
