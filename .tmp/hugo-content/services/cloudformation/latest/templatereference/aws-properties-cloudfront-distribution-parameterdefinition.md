This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution ParameterDefinition

A list of parameter values to add to the resource. A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : Definition,
  "Name" : String
}

```

### YAML

```yaml

  Definition:
    Definition
  Name: String

```

## Properties

`Definition`

The value that you assigned to the parameter.

_Required_: Yes

_Type_: [Definition](aws-properties-cloudfront-distribution-definition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the parameter.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9-_]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginShield

Restrictions

All content copied from https://docs.aws.amazon.com/.
