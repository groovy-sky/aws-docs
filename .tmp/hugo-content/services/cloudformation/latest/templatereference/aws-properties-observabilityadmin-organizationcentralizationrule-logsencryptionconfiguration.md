This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule LogsEncryptionConfiguration

Configuration for encrypting centralized log groups. This configuration is only applied to
destination log groups for which the corresponding source log groups are encrypted using
Customer Managed KMS Keys.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionConflictResolutionStrategy" : String,
  "EncryptionStrategy" : String,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  EncryptionConflictResolutionStrategy: String
  EncryptionStrategy: String
  KmsKeyArn: String

```

## Properties

`EncryptionConflictResolutionStrategy`

Conflict resolution strategy for centralization if the encryption strategy is set to
CUSTOMER\_MANAGED and the destination log group is encrypted with an AWS\_OWNED KMS Key. ALLOW
lets centralization go through while SKIP prevents centralization into the destination log
group.

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | SKIP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionStrategy`

Configuration that determines the encryption strategy of the destination log groups.
CUSTOMER\_MANAGED uses the configured KmsKeyArn to encrypt newly created destination log
groups.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED | AWS_OWNED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

KMS Key ARN belonging to the primary destination account and region, to encrypt newly
created central log groups in the primary destination.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws([a-z0-9\-]+)?:([a-zA-Z0-9\-]+):([a-z0-9\-]+)?:([0-9]{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogsBackupConfiguration

SourceLogsConfiguration

All content copied from https://docs.aws.amazon.com/.
