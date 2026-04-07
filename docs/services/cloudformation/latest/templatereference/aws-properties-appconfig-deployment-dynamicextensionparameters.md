This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Deployment DynamicExtensionParameters

A map of dynamic extension parameter names to values to pass to associated extensions
with `PRE_START_DEPLOYMENT` actions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExtensionReference" : String,
  "ParameterName" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ExtensionReference: String
  ParameterName: String
  ParameterValue: String

```

## Properties

`ExtensionReference`

The ARN or ID of the extension for which you are inserting a dynamic parameter.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterName`

The parameter name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterValue`

The parameter value.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppConfig::Deployment

Tag
