This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule

Creates a new case rule. In the Amazon Connect admin website, case rules are known as _case field conditions_. For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cases::CaseRule",
  "Properties" : {
      "Description" : String,
      "DomainId" : String,
      "Name" : String,
      "Rule" : CaseRuleDetails,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cases::CaseRule
Properties:
  Description: String
  DomainId: String
  Name: String
  Rule:
    CaseRuleDetails
  Tags:
    - Tag

```

## Properties

`Description`

Description of a case rule.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainId`

Unique identifier of a Cases domain.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: Updates are not supported.

`Name`

Name of the case rule.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[\S]$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rule`

Represents what rule type should take place, under what conditions.

_Required_: Yes

_Type_: [CaseRuleDetails](aws-properties-cases-caserule-caseruledetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-cases-caserule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the case rule. For example:

`arn:aws:cases:us-west-2:123456789012:domain/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111/case-rule/a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CaseRuleArn`

The Amazon Resource Name (ARN) of the case rule.

`CaseRuleId`

Unique identifier of a case rule.

`CreatedTime`

Timestamp when the resource was created.

`LastModifiedTime`

Timestamp when the resource was created or last modified.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Connect Cases

BooleanCondition

All content copied from https://docs.aws.amazon.com/.
