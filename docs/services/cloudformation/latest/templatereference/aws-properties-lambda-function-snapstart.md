This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function SnapStart

The function's [AWS Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplyOn" : String
}

```

### YAML

```yaml

  ApplyOn: String

```

## Properties

`ApplyOn`

Set `ApplyOn` to `PublishedVersions` to create a snapshot of the initialized execution environment when you publish a function version.

_Required_: Yes

_Type_: String

_Allowed values_: `PublishedVersions | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RuntimeManagementConfig

SnapStartResponse
