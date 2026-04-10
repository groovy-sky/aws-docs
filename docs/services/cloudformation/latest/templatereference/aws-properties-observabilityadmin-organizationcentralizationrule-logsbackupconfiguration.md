This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule LogsBackupConfiguration

Configuration for backing up centralized log data to a secondary region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "Region" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  Region: String

```

## Properties

`KmsKeyArn`

KMS Key ARN belonging to the primary destination account and backup region, to encrypt
newly created central log groups in the backup destination.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws([a-z0-9\-]+)?:([a-zA-Z0-9\-]+):([a-z0-9\-]+)?:([0-9]{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

Logs specific backup destination region within the primary destination account to which
log data should be centralized.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogGroupNameConfiguration

LogsEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
