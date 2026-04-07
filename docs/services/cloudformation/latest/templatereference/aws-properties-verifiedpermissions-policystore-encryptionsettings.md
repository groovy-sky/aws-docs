This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::PolicyStore EncryptionSettings

A structure that contains the encryption configuration for the policy store and child resources.

This data type is used as a request parameter in the [CreatePolicyStore](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreatePolicyStore.html) operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Default" : Json,
  "KmsEncryptionSettings" : KmsEncryptionSettings
}

```

### YAML

```yaml

  Default: Json
  KmsEncryptionSettings:
    KmsEncryptionSettings

```

## Properties

`Default`

This is the default encryption setting. The policy store uses an AWS owned key for encrypting data.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsEncryptionSettings`

The AWS KMS encryption settings for this policy store to encrypt data with.
It will contain the customer-managed KMS key, and a user-defined encryption context.

_Required_: No

_Type_: [KmsEncryptionSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-verifiedpermissions-policystore-kmsencryptionsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeletionProtection

EncryptionState
