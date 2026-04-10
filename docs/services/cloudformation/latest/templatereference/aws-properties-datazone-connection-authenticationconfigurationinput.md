This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection AuthenticationConfigurationInput

The authentication configuration of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationType" : String,
  "BasicAuthenticationCredentials" : BasicAuthenticationCredentials,
  "CustomAuthenticationCredentials" : {Key: Value, ...},
  "KmsKeyArn" : String,
  "OAuth2Properties" : OAuth2Properties,
  "SecretArn" : String
}

```

### YAML

```yaml

  AuthenticationType: String
  BasicAuthenticationCredentials:
    BasicAuthenticationCredentials
  CustomAuthenticationCredentials:
    Key: Value
  KmsKeyArn: String
  OAuth2Properties:
    OAuth2Properties
  SecretArn: String

```

## Properties

`AuthenticationType`

The authentication type of a connection.

_Required_: No

_Type_: String

_Allowed values_: `BASIC | OAUTH2 | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthenticationCredentials`

The basic authentication credentials of a connection.

_Required_: No

_Type_: [BasicAuthenticationCredentials](aws-properties-datazone-connection-basicauthenticationcredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAuthenticationCredentials`

The custom authentication credentials of a connection.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The KMS key ARN of a connection.

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[a-z0-9-]*:kms:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2Properties`

The oAuth2 properties of a connection.

_Required_: No

_Type_: [OAuth2Properties](aws-properties-datazone-connection-oauth2properties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The secret ARN of a connection.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso(-[bef])?))?:secretsmanager:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AthenaPropertiesInput

AuthorizationCodeProperties

All content copied from https://docs.aws.amazon.com/.
