This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Cluster ManagedStorageConfiguration

The managed storage configuration for the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FargateEphemeralStorageKmsKeyId" : String,
  "KmsKeyId" : String
}

```

### YAML

```yaml

  FargateEphemeralStorageKmsKeyId: String
  KmsKeyId: String

```

## Properties

`FargateEphemeralStorageKmsKeyId`

Specify the AWS Key Management Service key ID for Fargate ephemeral storage.

When you specify a `fargateEphemeralStorageKmsKeyId`, AWS
Fargate uses the key to encrypt data at rest in ephemeral storage. For more information
about Fargate ephemeral storage encryption, see [Customer managed keys for AWS Fargate ephemeral\
storage for Amazon ECS](../../../amazonecs/latest/developerguide/fargate-storage-encryption.md) in the _Amazon Elastic Container Service_
_Developer Guide_.

The key must be a single Region key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

Specify a AWS Key Management Service key ID to encrypt Amazon ECS managed storage.

When you specify a `kmsKeyId`, Amazon ECS uses the key to encrypt data
volumes managed by Amazon ECS that are attached to tasks in the cluster. The following
data volumes are managed by Amazon ECS: Amazon EBS. For more information about
encryption of Amazon EBS volumes attached to Amazon ECS tasks, see [Encrypt data stored in Amazon EBS volumes for Amazon ECS](../../../amazonecs/latest/developerguide/ebs-kms-encryption.md) in the
_Amazon Elastic Container Service Developer Guide_.

The key must be a single Region key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecuteCommandLogConfiguration

ServiceConnectDefaults

All content copied from https://docs.aws.amazon.com/.
