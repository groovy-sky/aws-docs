This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource KeyPairCredentials

The combination of username, private key and passphrase that are used as credentials.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyPairUsername" : String,
  "PrivateKey" : String,
  "PrivateKeyPassphrase" : String
}

```

### YAML

```yaml

  KeyPairUsername: String
  PrivateKey: String
  PrivateKeyPassphrase: String

```

## Properties

`KeyPairUsername`

Username

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKey`

PrivateKey

_Required_: Yes

_Type_: String

_Pattern_: `^-{5}BEGIN (ENCRYPTED )?PRIVATE KEY-{5}\u000D?\u000A([A-Za-z0-9/+]{64}\u000D?\u000A)*[A-Za-z0-9/+]{1,64}={0,2}\u000D?\u000A-{5}END (ENCRYPTED )?PRIVATE KEY-{5}(\u000D?\u000A)?$`

_Minimum_: `1600`

_Maximum_: `8000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKeyPassphrase`

PrivateKeyPassphrase

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdentityCenterConfiguration

ManifestFileLocation

All content copied from https://docs.aws.amazon.com/.
