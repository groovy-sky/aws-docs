This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup EngineVersion

The Athena engine version for running queries, or the PySpark engine
version for running sessions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EffectiveEngineVersion" : String,
  "SelectedEngineVersion" : String
}

```

### YAML

```yaml

  EffectiveEngineVersion: String
  SelectedEngineVersion: String

```

## Properties

`EffectiveEngineVersion`

Read only. The engine version on which the query runs. If the user requests a valid
engine version other than Auto, the effective engine version is the same as the engine
version that the user requested. If the user requests Auto, the effective engine version
is chosen by Athena. When a request to update the engine version is made by
a `CreateWorkGroup` or `UpdateWorkGroup` operation, the
`EffectiveEngineVersion` field is ignored.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedEngineVersion`

The engine version requested by the user. Possible values are determined by the output
of `ListEngineVersions`, including AUTO. The default is AUTO.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EngineConfiguration

ManagedLoggingConfiguration
