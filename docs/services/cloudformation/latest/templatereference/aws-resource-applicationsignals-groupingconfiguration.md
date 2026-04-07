This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::GroupingConfiguration

A structure that contains the complete grouping configuration for an account, including all defined grouping attributes and metadata about when it was last updated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApplicationSignals::GroupingConfiguration",
  "Properties" : {
      "GroupingAttributeDefinitions" : [ GroupingAttributeDefinition, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApplicationSignals::GroupingConfiguration
Properties:
  GroupingAttributeDefinitions:
    - GroupingAttributeDefinition

```

## Properties

`GroupingAttributeDefinitions`

An array of grouping attribute definitions that specify how services should be grouped based on various attributes and source keys.

_Required_: Yes

_Type_: Array of [GroupingAttributeDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationsignals-groupingconfiguration-groupingattributedefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the 12 digit AWS Account ID for the account.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The 12 digit AWS Account ID for the account.

`UpdatedAt`

The timestamp when this grouping configuration was last updated. When used in a raw HTTP Query API, it is formatted as epoch time in seconds.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApplicationSignals::Discovery

GroupingAttributeDefinition
