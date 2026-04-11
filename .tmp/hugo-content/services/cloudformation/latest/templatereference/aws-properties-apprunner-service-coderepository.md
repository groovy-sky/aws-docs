This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service CodeRepository

Describes a source code repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeConfiguration" : CodeConfiguration,
  "RepositoryUrl" : String,
  "SourceCodeVersion" : SourceCodeVersion,
  "SourceDirectory" : String
}

```

### YAML

```yaml

  CodeConfiguration:
    CodeConfiguration
  RepositoryUrl: String
  SourceCodeVersion:
    SourceCodeVersion
  SourceDirectory: String

```

## Properties

`CodeConfiguration`

Configuration for building and running the service from a source code repository.

###### Note

`CodeConfiguration` is required only for `CreateService` request.

_Required_: No

_Type_: [CodeConfiguration](aws-properties-apprunner-service-codeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryUrl`

The location of the repository that contains the source code.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceCodeVersion`

The version that should be used within the source code repository.

_Required_: Yes

_Type_: [SourceCodeVersion](aws-properties-apprunner-service-sourcecodeversion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceDirectory`

The path of the directory that stores source code and configuration files. The build and start commands also execute from here. The path is absolute
from root and, if not specified, defaults to the repository root.

_Required_: No

_Type_: String

_Pattern_: `[^\x00]+`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeConfigurationValues

EgressConfiguration

All content copied from https://docs.aws.amazon.com/.
