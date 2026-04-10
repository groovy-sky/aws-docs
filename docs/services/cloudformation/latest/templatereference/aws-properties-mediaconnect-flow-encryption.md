This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow Encryption

Encryption information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Algorithm" : String,
  "ConstantInitializationVector" : String,
  "DeviceId" : String,
  "KeyType" : String,
  "Region" : String,
  "ResourceId" : String,
  "RoleArn" : String,
  "SecretArn" : String,
  "Url" : String
}

```

### YAML

```yaml

  Algorithm: String
  ConstantInitializationVector: String
  DeviceId: String
  KeyType: String
  Region: String
  ResourceId: String
  RoleArn: String
  SecretArn: String
  Url: String

```

## Properties

`Algorithm`

The type of algorithm that is used for static key encryption (such as aes128, aes192, or
aes256). If you are using SPEKE or SRT-password encryption, this property must be left blank.

_Required_: No

_Type_: String

_Allowed values_: `aes128 | aes192 | aes256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConstantInitializationVector`

A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceId`

The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyType`

The type of key that is used for the encryption. If you don't specify a
`keyType` value, the service uses the default setting
( `static-key`). Valid key types are: `static-key`, `speke`, and `srt-password`.

_Required_: No

_Type_: String

_Allowed values_: `speke | static-key | srt-password`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that you created during setup (when you set up MediaConnect as a trusted entity).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::[0-9]{12}:role/[a-zA-Z0-9_+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):secretsmanager:[a-z0-9-]+:[0-9]{12}:secret:[a-zA-Z0-9/_+=.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncodingConfig

FailoverConfig

All content copied from https://docs.aws.amazon.com/.
