This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Application IamIdentityCenterOptions

Configuration settings for IAM Identity Center in an OpenSearch application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "IamIdentityCenterInstanceArn" : String,
  "IamRoleForIdentityCenterApplicationArn" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  IamIdentityCenterInstanceArn: String
  IamRoleForIdentityCenterApplicationArn: String

```

## Properties

`Enabled`

Indicates whether IAM Identity Center is enabled for the OpenSearch application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamIdentityCenterInstanceArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleForIdentityCenterApplicationArn`

The Amazon Resource Name (ARN) of the IAM role assigned to the IAM Identity Center
application for the OpenSearch application.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws\-cn|aws\-us\-gov|aws\-iso|aws\-iso\-b):iam::[0-9]+:role\/.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSource

Tag

All content copied from https://docs.aws.amazon.com/.
