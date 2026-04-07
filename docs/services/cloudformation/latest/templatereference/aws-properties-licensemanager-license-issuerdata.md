This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LicenseManager::License IssuerData

Details associated with the issuer of a license.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SignKey" : String
}

```

### YAML

```yaml

  Name: String
  SignKey: String

```

## Properties

`Name`

Issuer name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignKey`

Asymmetric KMS key from AWS Key Management Service. The KMS key must have a key usage of sign and verify,
and support the RSASSA-PSS SHA-256 signing algorithm.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Entitlement

Metadata
