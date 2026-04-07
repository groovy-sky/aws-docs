This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecretsManager::Secret GenerateSecretString

Generates a random password. We recommend that you specify the maximum length and include
every character type that the system you are generating a password for can support.

**Required permissions:** `secretsmanager:GetRandomPassword`. For more information, see [IAM policy actions for Secrets Manager](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awssecretsmanager.html#awssecretsmanager-actions-as-permissions) and [Authentication and access control\
in Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeCharacters" : String,
  "ExcludeLowercase" : Boolean,
  "ExcludeNumbers" : Boolean,
  "ExcludePunctuation" : Boolean,
  "ExcludeUppercase" : Boolean,
  "GenerateStringKey" : String,
  "IncludeSpace" : Boolean,
  "PasswordLength" : Integer,
  "RequireEachIncludedType" : Boolean,
  "SecretStringTemplate" : String
}

```

### YAML

```yaml

  ExcludeCharacters: String
  ExcludeLowercase: Boolean
  ExcludeNumbers: Boolean
  ExcludePunctuation: Boolean
  ExcludeUppercase: Boolean
  GenerateStringKey:
    String
  IncludeSpace: Boolean
  PasswordLength: Integer
  RequireEachIncludedType: Boolean
  SecretStringTemplate:
    String

```

## Properties

`ExcludeCharacters`

A string of the characters that you don't want in the password.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeLowercase`

Specifies whether to exclude lowercase letters from the password. If
you don't include this switch, the password can contain lowercase letters.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeNumbers`

Specifies whether to exclude numbers from the password. If you don't
include this switch, the password can contain numbers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludePunctuation`

Specifies whether to exclude the following punctuation characters from the password:
``! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~``.
If you don't include this switch, the password can contain punctuation.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeUppercase`

Specifies whether to exclude uppercase letters from the password. If you
don't include this switch, the password can contain uppercase letters.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerateStringKey`

The JSON key name for the key/value pair, where the value is the generated password. This
pair is added to the JSON structure specified by the `SecretStringTemplate`
parameter. If you specify this parameter, then you must also specify
`SecretStringTemplate`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSpace`

Specifies whether to include the space character. If you
include this switch, the password can contain space characters.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PasswordLength`

The length of the password. If you don't include this parameter, the
default length is 32 characters.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireEachIncludedType`

Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation.
If you don't include this switch, the password contains at least one of every character type.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretStringTemplate`

A template that the generated string must match. When you make a change to this property, a new secret version is created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [GetRandomPassword](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetRandomPassword.html) in the AWS Secrets Manager API Reference

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecretsManager::Secret

ReplicaRegion
