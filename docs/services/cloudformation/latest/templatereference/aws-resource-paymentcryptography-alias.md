This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PaymentCryptography::Alias

Creates an _alias_, or a friendly name, for an AWS Payment Cryptography key. You can
use an alias to identify a key in the console and when you call cryptographic operations
such as [EncryptData](../../../../reference/payment-cryptography/latest/dataapireference/api-encryptdata.md) or [DecryptData](../../../../reference/payment-cryptography/latest/dataapireference/api-decryptdata.md).

You can associate the alias with any key in the same AWS Region. Each
alias is associated with only one key at a time, but a key can have multiple aliases. You
can't create an alias without a key. The alias must be unique in the account and AWS Region, but you can create another alias with the same name in a different
AWS Region.

To change the key that's associated with the alias, call [UpdateAlias](../../../../reference/payment-cryptography/latest/apireference/api-updatealias.md).
To delete the alias, call [DeleteAlias](../../../../reference/payment-cryptography/latest/apireference/api-deletealias.md).
These operations don't affect the underlying key. To get the alias that you created, call
[ListAliases](../../../../reference/payment-cryptography/latest/apireference/api-listaliases.md).

**Cross-account use**: This operation can't be used across different AWS accounts.

**Related operations:**

- [DeleteAlias](../../../../reference/payment-cryptography/latest/apireference/api-deletealias.md)

- [GetAlias](../../../../reference/payment-cryptography/latest/apireference/api-getalias.md)

- [ListAliases](../../../../reference/payment-cryptography/latest/apireference/api-listaliases.md)

- [UpdateAlias](../../../../reference/payment-cryptography/latest/apireference/api-updatealias.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PaymentCryptography::Alias",
  "Properties" : {
      "AliasName" : String,
      "KeyArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::PaymentCryptography::Alias
Properties:
  AliasName: String
  KeyArn: String

```

## Properties

`AliasName`

A friendly name that you can use to refer to a key. The value must begin with
`alias/`.

###### Important

Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in AWS CloudTrail logs and other output.

_Required_: Yes

_Type_: String

_Pattern_: `^alias/[a-zA-Z0-9/_-]+$`

_Minimum_: `7`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyArn`

The `KeyARN` of the key associated with the alias.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws:payment-cryptography:[a-z]{2}-[a-z]{1,16}-[0-9]+:[0-9]{12}:key/[0-9a-zA-Z]{16,64}$`

_Minimum_: `70`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Payment Cryptography

AWS::PaymentCryptography::Key

All content copied from https://docs.aws.amazon.com/.
