This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Resolver AppSyncRuntime

Describes a runtime used by an AWS AppSync resolver or AWS AppSync function. Specifies the name and version of the runtime to use.
Note that if a runtime is specified, code must also be specified.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "RuntimeVersion" : String
}

```

### YAML

```yaml

  Name: String
  RuntimeVersion: String

```

## Properties

`Name`

The `name` of the runtime to use. Currently, the only allowed value is
`APPSYNC_JS`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeVersion`

The `version` of the runtime to use. Currently, the only allowed version is
`1.0.0`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppSync::Resolver

CachingConfig

All content copied from https://docs.aws.amazon.com/.
