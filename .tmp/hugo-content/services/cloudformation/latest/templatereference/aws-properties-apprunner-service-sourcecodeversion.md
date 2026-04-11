This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service SourceCodeVersion

Identifies a version of code that AWS App Runner refers to within a source code repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Type: String
  Value: String

```

## Properties

`Type`

The type of version identifier.

For a git-based repository, branches represent versions.

_Required_: Yes

_Type_: String

_Allowed values_: `BRANCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A source code version.

For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceObservabilityConfiguration

SourceConfiguration

All content copied from https://docs.aws.amazon.com/.
