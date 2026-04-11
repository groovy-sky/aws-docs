This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service CodeConfigurationValues

Describes the basic configuration needed for building and running an AWS App Runner service. This type doesn't support the full set of possible
configuration options. Fur full configuration capabilities, use a `apprunner.yaml` file in the source code repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BuildCommand" : String,
  "Port" : String,
  "Runtime" : String,
  "RuntimeEnvironmentSecrets" : [ KeyValuePair, ... ],
  "RuntimeEnvironmentVariables" : [ KeyValuePair, ... ],
  "StartCommand" : String
}

```

### YAML

```yaml

  BuildCommand: String
  Port: String
  Runtime: String
  RuntimeEnvironmentSecrets:
    - KeyValuePair
  RuntimeEnvironmentVariables:
    - KeyValuePair
  StartCommand: String

```

## Properties

`BuildCommand`

The command App Runner runs to build your
application.

_Required_: No

_Type_: String

_Pattern_: `[^\x0a\x0d]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port that your application listens to in the container.

Default: `8080`

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Runtime`

A runtime environment type for building and running an App Runner service.
It represents a
programming language runtime.

_Required_: Yes

_Type_: String

_Allowed values_: `PYTHON_3 | NODEJS_12 | NODEJS_14 | CORRETTO_8 | CORRETTO_11 | NODEJS_16 | GO_1 | DOTNET_6 | PHP_81 | RUBY_31 | PYTHON_311 | NODEJS_18 | NODEJS_22`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeEnvironmentSecrets`

An array of key-value pairs representing the secrets and parameters that get referenced to your service as an environment variable. The supported
values are either the full Amazon Resource Name (ARN) of the AWS Secrets Manager secret or the full ARN of the parameter in the AWS Systems Manager
Parameter Store.

###### Note

- If the AWS Systems Manager Parameter Store parameter exists in the same AWS Region as the service that you're launching, you can use
either the full ARN or name of the secret. If the parameter exists in a different Region, then the full ARN must be specified.

- Currently, cross account referencing of AWS Systems Manager Parameter Store parameter is not supported.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-apprunner-service-keyvaluepair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeEnvironmentVariables`

The environment variables that are available to your running AWS App Runner service. An array of key-value pairs.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-apprunner-service-keyvaluepair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartCommand`

The command App Runner runs to start your
application.

_Required_: No

_Type_: String

_Pattern_: `[^\x0a\x0d]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeConfiguration

CodeRepository

All content copied from https://docs.aws.amazon.com/.
