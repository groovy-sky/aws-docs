This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::OrganizationConformancePack ConformancePackInputParameter

Input parameters in the form of key-value pairs for the conformance pack, both of which you define.
Keys can have a maximum character length of 255 characters, and values can have a maximum length of 4096 characters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterName" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ParameterName: String
  ParameterValue: String

```

## Properties

`ParameterName`

One part of a key-value pair.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

One part of a key-value pair.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Config::OrganizationConformancePack

AWS::Config::RemediationConfiguration
