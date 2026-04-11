This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition RepositoryCredentials

The repository credentials for private registry authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialsParameter" : String
}

```

### YAML

```yaml

  CredentialsParameter: String

```

## Properties

`CredentialsParameter`

The Amazon Resource Name (ARN) of the secret containing the private repository
credentials.

###### Note

When you use the Amazon ECS API, AWS CLI, or AWS SDK, if the secret
exists in the same Region as the task that you're launching then you can use either
the full ARN or the name of the secret. When you use the AWS Management
Console, you must specify the full ARN of the secret.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MountPoint

RestartPolicy

All content copied from https://docs.aws.amazon.com/.
