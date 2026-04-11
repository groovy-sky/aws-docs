This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project GitSubmodulesConfig

`GitSubmodulesConfig` is a property of the [AWS CodeBuild Project Source](../userguide/aws-properties-codebuild-project-source.md)
property type that specifies information about the Git submodules configuration for the build project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FetchSubmodules" : Boolean
}

```

### YAML

```yaml

  FetchSubmodules: Boolean

```

## Properties

`FetchSubmodules`

Set to true to fetch Git submodules for your AWS CodeBuild build project.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [GitSubmodulesConfig](../../../../reference/codebuild/latest/apireference/api-gitsubmodulesconfig.md) in the _AWS CodeBuild API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentVariable

LogsConfig

All content copied from https://docs.aws.amazon.com/.
