This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::EventBridgeRuleTemplateGroup

The `AWS::MediaLive::EventBridgeRuleTemplateGroup` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::EventBridgeRuleTemplateGroup",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::EventBridgeRuleTemplateGroup
Properties:
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`Description`

A resource's optional description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A resource's name. Names must be unique within the scope of a resource type in a specific region.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\s]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

An eventbridge rule template group's ARN (Amazon Resource Name)

`CreatedAt`

The date and time of resource creation.

`Id`

An eventbridge rule template group's id. Amazon Web Services provided template groups have ids that start with <code>\`aws-\`</code>.

`Identifier`

Property description not available.

`ModifiedAt`

The date and time of latest resource modification.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventBridgeRuleTemplateTarget

AWS::MediaLive::Input
