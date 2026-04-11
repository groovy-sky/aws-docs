This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection AuthenticationConfigurationInput

A structure containing the authentication configuration in the CreateConnection request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationType" : String,
  "BasicAuthenticationCredentials" : BasicAuthenticationCredentials,
  "CustomAuthenticationCredentials" : Json,
  "KmsKeyArn" : String,
  "OAuth2Properties" : OAuth2PropertiesInput,
  "SecretArn" : String
}

```

### YAML

```yaml

  AuthenticationType: String
  BasicAuthenticationCredentials:
    BasicAuthenticationCredentials
  CustomAuthenticationCredentials: Json
  KmsKeyArn: String
  OAuth2Properties:
    OAuth2PropertiesInput
  SecretArn: String

```

## Properties

`AuthenticationType`

A structure containing the authentication configuration in the CreateConnection request.

_Required_: Yes

_Type_: String

_Allowed values_: `BASIC | OAUTH2 | CUSTOM | IAM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthenticationCredentials`

The credentials used when the authentication type is basic authentication.

_Required_: No

_Type_: [BasicAuthenticationCredentials](aws-properties-glue-connection-basicauthenticationcredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAuthenticationCredentials`

The credentials used when the authentication type is custom authentication.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the KMS key used to encrypt the connection. Only taken an as input in the request and stored in the Secret Manager.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[a-z0-9-]*:kms:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2Properties`

The properties for OAuth2 authentication in the CreateConnection request.

_Required_: No

_Type_: [OAuth2PropertiesInput](aws-properties-glue-connection-oauth2propertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The secret manager ARN to store credentials in the CreateConnection request.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso(-[bef])?))?:secretsmanager:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::Connection

AuthorizationCodeProperties

All content copied from https://docs.aws.amazon.com/.
