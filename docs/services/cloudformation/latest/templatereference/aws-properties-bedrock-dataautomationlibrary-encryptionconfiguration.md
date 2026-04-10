This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationLibrary EncryptionConfiguration

Encryption settings for an invocation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsEncryptionContext" : {Key: Value, ...},
  "KmsKeyId" : String
}

```

### YAML

```yaml

  KmsEncryptionContext:
    Key: Value
  KmsKeyId: String

```

## Properties

`KmsEncryptionContext`

Name-value pairs to include as an encryption context.

_Required_: No

_Type_: Object of String

_Pattern_: `^.*\S.*$`

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

A KMS key ID to use for encryption.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9:_/+=,@.-]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::DataAutomationLibrary

EntityTypeInfo

All content copied from https://docs.aws.amazon.com/.
