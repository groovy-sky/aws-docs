This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Permission Condition

The `Condition` property type specifies Property description not available. for an [AWS::QBusiness::Permission](aws-resource-qbusiness-permission.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionKey" : String,
  "ConditionOperator" : String,
  "ConditionValues" : [ String, ... ]
}

```

### YAML

```yaml

  ConditionKey: String
  ConditionOperator: String
  ConditionValues:
    - String

```

## Properties

`ConditionKey`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^aws:PrincipalTag/qbusiness-dataaccessor:[a-zA-Z]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConditionOperator`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `StringEquals`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConditionValues`

Property description not available.

_Required_: Yes

_Type_: Array of String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::QBusiness::Permission

AWS::QBusiness::Plugin
