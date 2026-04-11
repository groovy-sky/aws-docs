This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AIOps::InvestigationGroup CrossAccountConfiguration

This structure contains information about the cross-account configuration in the account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceRoleArn" : String
}

```

### YAML

```yaml

  SourceRoleArn: String

```

## Properties

`SourceRoleArn`

The ARN of an existing role which will be used to do investigations on your behalf.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChatbotNotificationChannel

EncryptionConfigMap

All content copied from https://docs.aws.amazon.com/.
