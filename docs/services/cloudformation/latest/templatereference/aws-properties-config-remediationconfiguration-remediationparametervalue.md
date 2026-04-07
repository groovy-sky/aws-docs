This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::RemediationConfiguration RemediationParameterValue

The value is either a dynamic (resource) value or a static value. You must select either a dynamic value or a static value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceValue" : String,
  "StaticValue" : [ String, ... ]
}

```

### YAML

```yaml

  ResourceValue: String
  StaticValue:
    - String

```

## Properties

`ResourceValue`

The value is dynamic and changes at run-time.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticValue`

The value is static and does not change at run-time.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExecutionControls

SsmControls
