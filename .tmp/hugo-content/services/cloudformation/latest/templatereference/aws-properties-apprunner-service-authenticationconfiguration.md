This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service AuthenticationConfiguration

Describes resources needed to authenticate access to some source repositories. The specific resource depends on the repository provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessRoleArn" : String,
  "ConnectionArn" : String
}

```

### YAML

```yaml

  AccessRoleArn: String
  ConnectionArn: String

```

## Properties

`AccessRoleArn`

The Amazon Resource Name (ARN) of the IAM role that grants the App Runner service access to a source repository. It's required for ECR image repositories
(but not for ECR Public repositories).

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::[0-9]{12}:role/[\w+=,.@-]{1,64}`

_Minimum_: `29`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionArn`

The Amazon Resource Name (ARN) of the App Runner connection that enables the App Runner service to connect to a source repository. It's required for GitHub code
repositories.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppRunner::Service

CodeConfiguration

All content copied from https://docs.aws.amazon.com/.
