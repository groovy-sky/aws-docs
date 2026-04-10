This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index JwtTokenTypeConfiguration

Provides the configuration information for the JWT token type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClaimRegex" : String,
  "GroupAttributeField" : String,
  "Issuer" : String,
  "KeyLocation" : String,
  "SecretManagerArn" : String,
  "URL" : String,
  "UserNameAttributeField" : String
}

```

### YAML

```yaml

  ClaimRegex: String
  GroupAttributeField: String
  Issuer: String
  KeyLocation: String
  SecretManagerArn: String
  URL: String
  UserNameAttributeField: String

```

## Properties

`ClaimRegex`

The regular expression that identifies the claim.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupAttributeField`

The group attribute field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The issuer of the token.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `65`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyLocation`

The location of the key.

_Required_: Yes

_Type_: String

_Allowed values_: `URL | SECRET_MANAGER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretManagerArn`

The Amazon Resource Name (arn) of the secret.

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URL`

The signing key URL.

_Required_: No

_Type_: String

_Pattern_: `^(https?|ftp|file):\/\/([^\s]*)`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserNameAttributeField`

The user name attribute field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JsonTokenTypeConfiguration

Relevance

All content copied from https://docs.aws.amazon.com/.
