This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector IAMConnectionMetadata

Authentication metadata for IAM-based connections, used for first-party AWS service integrations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String
}

```

### YAML

```yaml

  RoleArn: String

```

## Properties

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role to assume for authentication with AWS services. This IAM role should be in the same account as Quick Sight.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientCredentialsGrantMetadata

NoneConnectionMetadata

All content copied from https://docs.aws.amazon.com/.
