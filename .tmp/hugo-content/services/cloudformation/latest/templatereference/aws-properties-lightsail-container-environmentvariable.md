This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container EnvironmentVariable

`EnvironmentVariable` is a property of the [Container](../userguide/aws-properties-lightsail-container-container.md) property. It describes the environment variables of a container on a container service which are key-value parameters that
provide dynamic configuration of the application or script run by the container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : String,
  "Variable" : String
}

```

### YAML

```yaml

  Value: String
  Variable: String

```

## Properties

`Value`

The environment variable value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variable`

The environment variable key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcrImagePullerRole

HealthCheckConfig

All content copied from https://docs.aws.amazon.com/.
