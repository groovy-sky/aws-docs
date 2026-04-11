This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service CodeConfiguration

Describes the configuration that AWS App Runner uses to build and run an App Runner service from a source code repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeConfigurationValues" : CodeConfigurationValues,
  "ConfigurationSource" : String
}

```

### YAML

```yaml

  CodeConfigurationValues:
    CodeConfigurationValues
  ConfigurationSource: String

```

## Properties

`CodeConfigurationValues`

The basic configuration for building and running the App Runner service. Use it to quickly launch an App Runner service without providing a
`apprunner.yaml` file in the source code repository (or ignoring the file if it exists).

_Required_: No

_Type_: [CodeConfigurationValues](aws-properties-apprunner-service-codeconfigurationvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationSource`

The source of the App Runner configuration. Values are interpreted as follows:

- `REPOSITORY` – App Runner reads configuration values from the `apprunner.yaml` file in the source code repository and
ignores `CodeConfigurationValues`.

- `API` – App Runner uses configuration values provided in `CodeConfigurationValues` and ignores the
`apprunner.yaml` file in the source code repository.

_Required_: Yes

_Type_: String

_Allowed values_: `REPOSITORY | API`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthenticationConfiguration

CodeConfigurationValues

All content copied from https://docs.aws.amazon.com/.
