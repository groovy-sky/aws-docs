This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition Secret

An object that represents the secret to expose to your container. Secrets can be exposed to
a container in the following ways:

- To inject sensitive data into your containers as environment variables, use the
`secrets` container definition parameter.

- To reference sensitive information in the log configuration of a container, use the
`secretOptions` container definition parameter.

For more information, see [Specifying sensitive data](../../../batch/latest/userguide/specifying-sensitive-data.md) in the
_AWS Batch User Guide_.

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

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueFrom`

The secret to expose to the container. The supported values are either the full Amazon Resource Name (ARN) of
the AWS Secrets Manager secret or the full ARN of the parameter in the AWS Systems Manager Parameter Store.

###### Note

If the AWS Systems Manager Parameter Store parameter exists in the same Region as the job you're
launching, then you can use either the full Amazon Resource Name (ARN) or name of the parameter. If the parameter
exists in a different Region, then the full ARN must be specified.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuntimePlatform

TaskContainerDependency

All content copied from https://docs.aws.amazon.com/.
