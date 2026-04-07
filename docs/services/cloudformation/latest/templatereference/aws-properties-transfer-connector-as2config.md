This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Connector As2Config

A structure that contains the parameters for an AS2 connector object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AsyncMdnConfig" : ConnectorAsyncMdnConfig,
  "BasicAuthSecretId" : String,
  "Compression" : String,
  "EncryptionAlgorithm" : String,
  "LocalProfileId" : String,
  "MdnResponse" : String,
  "MdnSigningAlgorithm" : String,
  "MessageSubject" : String,
  "PartnerProfileId" : String,
  "PreserveContentType" : String,
  "SigningAlgorithm" : String
}

```

### YAML

```yaml

  AsyncMdnConfig:
    ConnectorAsyncMdnConfig
  BasicAuthSecretId: String
  Compression: String
  EncryptionAlgorithm: String
  LocalProfileId: String
  MdnResponse: String
  MdnSigningAlgorithm: String
  MessageSubject: String
  PartnerProfileId: String
  PreserveContentType: String
  SigningAlgorithm: String

```

## Properties

`AsyncMdnConfig`

Property description not available.

_Required_: No

_Type_: [ConnectorAsyncMdnConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-connector-connectorasyncmdnconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthSecretId`

Provides Basic authentication support to the AS2 Connectors API. To use Basic
authentication, you must provide the name or Amazon Resource Name (ARN) of a secret in
AWS Secrets Manager.

The default value for this parameter is `null`, which indicates that Basic
authentication is not enabled for the connector.

If the connector should use Basic authentication, the secret needs to be in the
following format:

`{ "Username": "user-name", "Password": "user-password" }`

Replace `user-name` and `user-password` with the credentials for
the actual user that is being authenticated.

Note the following:

- You are storing these credentials in Secrets Manager, _not passing_
_them directly_ into this API.

- If you are using the API, SDKs, or CloudFormation to configure your connector,
then you must create the secret before you can enable Basic authentication.
However, if you are using the AWS management console, you can
have the system create the secret for you.

If you have previously enabled Basic authentication for a connector, you can disable
it by using the `UpdateConnector` API call. For example, if you are using the
CLI, you can run the following command to remove Basic authentication:

`update-connector --connector-id my-connector-id --as2-config
                'BasicAuthSecretId=""'`

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compression`

Specifies whether the AS2 file is compressed.

_Required_: No

_Type_: String

_Allowed values_: `ZLIB | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionAlgorithm`

The algorithm that is used to encrypt the file.

Note the following:

- Do not use the `DES_EDE3_CBC` algorithm unless you must support a
legacy client that requires it, as it is a weak encryption algorithm.

- You can only specify `NONE` if the URL for your connector uses
HTTPS. Using HTTPS ensures that no traffic is sent in clear text.

_Required_: No

_Type_: String

_Allowed values_: `AES128_CBC | AES192_CBC | AES256_CBC | NONE | DES_EDE3_CBC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalProfileId`

A unique identifier for the AS2 local profile.

_Required_: No

_Type_: String

_Pattern_: `^p-([0-9a-f]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MdnResponse`

Used for outbound requests (from an AWS Transfer Family connector to a partner AS2 server) to determine whether
the partner response for transfers is synchronous or asynchronous. Specify either of the following values:

- `ASYNC`: The system expects an asynchronous MDN response, confirming that the file was transferred successfully (or not).

- `SYNC`: The system expects a synchronous MDN response, confirming that the file was transferred successfully (or not).

- `NONE`: Specifies that no MDN response is required.

_Required_: No

_Type_: String

_Allowed values_: `SYNC | ASYNC | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MdnSigningAlgorithm`

The signing algorithm for the MDN response.

###### Note

If set to DEFAULT (or not set at all), the value for `SigningAlgorithm`
is used.

_Required_: No

_Type_: String

_Allowed values_: `SHA256 | SHA384 | SHA512 | SHA1 | NONE | DEFAULT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageSubject`

Used as the `Subject` HTTP header attribute in AS2 messages that are being
sent with the connector.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\u007E\t]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartnerProfileId`

A unique identifier for the partner profile for the connector.

_Required_: No

_Type_: String

_Pattern_: `^p-([0-9a-f]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreserveContentType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningAlgorithm`

The algorithm that is used to sign the AS2 messages sent with the connector.

_Required_: No

_Type_: String

_Allowed values_: `SHA256 | SHA384 | SHA512 | SHA1 | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Transfer::Connector

ConnectorAsyncMdnConfig
