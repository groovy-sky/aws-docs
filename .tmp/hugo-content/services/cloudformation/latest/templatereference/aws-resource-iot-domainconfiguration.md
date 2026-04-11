This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::DomainConfiguration

Specifies a domain configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::DomainConfiguration",
  "Properties" : {
      "ApplicationProtocol" : String,
      "AuthenticationType" : String,
      "AuthorizerConfig" : AuthorizerConfig,
      "ClientCertificateConfig" : ClientCertificateConfig,
      "DomainConfigurationName" : String,
      "DomainConfigurationStatus" : String,
      "DomainName" : String,
      "ServerCertificateArns" : [ String, ... ],
      "ServerCertificateConfig" : ServerCertificateConfig,
      "ServiceType" : String,
      "Tags" : [ Tag, ... ],
      "TlsConfig" : TlsConfig,
      "ValidationCertificateArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::DomainConfiguration
Properties:
  ApplicationProtocol: String
  AuthenticationType: String
  AuthorizerConfig:
    AuthorizerConfig
  ClientCertificateConfig:
    ClientCertificateConfig
  DomainConfigurationName: String
  DomainConfigurationStatus: String
  DomainName: String
  ServerCertificateArns:
    - String
  ServerCertificateConfig:
    ServerCertificateConfig
  ServiceType: String
  Tags:
    - Tag
  TlsConfig:
    TlsConfig
  ValidationCertificateArn: String

```

## Properties

`ApplicationProtocol`

An enumerated string that speciﬁes the application-layer protocol.

_Required_: No

_Type_: String

_Allowed values_: `SECURE_MQTT | MQTT_WSS | HTTPS | DEFAULT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

An enumerated string that speciﬁes the authentication type.

_Required_: No

_Type_: String

_Allowed values_: `AWS_X509 | CUSTOM_AUTH | AWS_SIGV4 | CUSTOM_AUTH_X509 | DEFAULT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerConfig`

An object that specifies the authorization service for a domain.

_Required_: No

_Type_: [AuthorizerConfig](aws-properties-iot-domainconfiguration-authorizerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCertificateConfig`

An object that speciﬁes the client certificate conﬁguration for a domain.

_Required_: No

_Type_: [ClientCertificateConfig](aws-properties-iot-domainconfiguration-clientcertificateconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainConfigurationName`

The name of the domain configuration. This value must be unique to a region.

_Required_: No

_Type_: String

_Pattern_: `^[\w.-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainConfigurationStatus`

The status to which the domain configuration should be updated.

Valid values: `ENABLED` \| `DISABLED`

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The name of the domain.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerCertificateArns`

The ARNs of the certificates that AWS IoT passes to the device during the TLS handshake. Currently you can specify only one certificate ARN.
This value is not required for AWS-managed domains.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `2048 | 1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerCertificateConfig`

The server certificate configuration.

For more information, see [Configurable\
endpoints](../../../iot/latest/developerguide/iot-custom-endpoints-configurable.md) from the AWS IoT Core Developer Guide.

_Required_: No

_Type_: [ServerCertificateConfig](aws-properties-iot-domainconfiguration-servercertificateconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceType`

The type of service delivered by the endpoint.

###### Note

AWS IoT Core currently supports only the `DATA` service type.

_Required_: No

_Type_: String

_Allowed values_: `DATA | CREDENTIAL_PROVIDER | JOBS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata which can be used to manage the domain configuration.

###### Note

For URI Request parameters use format: ...key1=value1&key2=value2...

For the CLI command-line parameter use format: &&tags
"key1=value1&key2=value2..."

For the cli-input-json file use format: "tags":
"key1=value1&key2=value2..."

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-domainconfiguration-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlsConfig`

An object that specifies the TLS configuration for a domain.

_Required_: No

_Type_: [TlsConfig](aws-properties-iot-domainconfiguration-tlsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationCertificateArn`

The certificate used to validate the server certificate and prove domain name ownership. This certificate must be signed by a public certificate authority.
This value is not required for AWS-managed domains.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-cn|-us-gov|-iso-b|-iso)?:acm:[a-z]{2}-(gov-|iso-|isob-)?[a-z]{4,9}-\d{1}:\d{12}:certificate/[a-zA-Z0-9/-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the domain configuration name. For example:

`{ "Ref": "MyDomainConfiguration" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the domain configuration.

`DomainType`

The type of service delivered by the domain.

`ServerCertificates`

The ARNs of the certificates that AWS IoT passes to the device during the
TLS handshake. Currently you can specify only one certificate ARN. This value is not
required for AWS-managed domains.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AuthorizerConfig

All content copied from https://docs.aws.amazon.com/.
