This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster Provider

Identifies the AWS Key Management Service (AWS KMS) key used to encrypt the secrets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyArn" : String
}

```

### YAML

```yaml

  KeyArn: String

```

## Properties

`KeyArn`

Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric and created in
the same AWS Region as the cluster. If the KMS key was created in a different
account, the [IAM principal](../../../iam/latest/userguide/id-roles-terms-and-concepts.md) must have access to the KMS key. For more information,
see [Allowing\
users in other accounts to use a KMS key](../../../kms/latest/developerguide/key-policy-modifying-external-accounts.md) in the _AWS Key Management Service_
_Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutpostConfig

RemoteNetworkConfig

All content copied from https://docs.aws.amazon.com/.
