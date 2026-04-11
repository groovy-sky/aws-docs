This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition Secret

An object representing the secret to expose to your container. Secrets can be exposed
to a container in the following ways:

- To inject sensitive data into your containers as environment variables, use
the `secrets` container definition parameter.

- To reference sensitive information in the log configuration of a container,
use the `secretOptions` container definition parameter.

For more information, see [Specifying\
sensitive data](../../../amazonecs/latest/developerguide/specifying-sensitive-data.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "ValueFrom" : String
}

```

### YAML

```yaml

  Name: String
  ValueFrom: String

```

## Properties

`Name`

The name of the secret.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValueFrom`

The secret to expose to the container. The supported values are either the full ARN of
the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter
Store.

For information about the require AWS Identity and Access Management permissions,
see [Required IAM permissions for Amazon ECS secrets](../../../amazonecs/latest/developerguide/specifying-sensitive-data-secrets.md#secrets-iam) (for Secrets Manager) or
[Required IAM permissions for Amazon ECS secrets](../../../amazonecs/latest/developerguide/specifying-sensitive-data-parameters.md) (for Systems Manager
Parameter store) in the _Amazon Elastic Container Service Developer_
_Guide_.

###### Note

If the SSM Parameter Store parameter exists in the same Region as the task you're
launching, then you can use either the full ARN or name of the parameter. If the
parameter exists in a different Region, then the full ARN must be specified.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestartPolicy

SystemControl

All content copied from https://docs.aws.amazon.com/.
