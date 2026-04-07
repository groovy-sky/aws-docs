This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AssistantAssociation

Specifies an association between an Amazon Connect Wisdom assistant and another
resource. Currently, the only supported association is with a knowledge base. An
assistant can have only a single association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AssistantAssociation",
  "Properties" : {
      "AssistantId" : String,
      "Association" : AssociationData,
      "AssociationType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AssistantAssociation
Properties:
  AssistantId: String
  Association:
    AssociationData
  AssociationType: String
  Tags:
    - Tag

```

## Properties

`AssistantId`

The identifier of the Wisdom assistant.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Association`

The identifier of the associated resource.

_Required_: Yes

_Type_: [AssociationData](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-assistantassociation-associationdata.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssociationType`

The type of association.

_Required_: Yes

_Type_: String

_Allowed values_: `KNOWLEDGE_BASE | EXTERNAL_BEDROCK_KNOWLEDGE_BASE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-assistantassociation-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the association ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssistantArn`

The Amazon Resource Name (ARN) of the Wisdom assistant.

`AssistantAssociationArn`

The Amazon Resource Name (ARN) of the assistant association.

`AssistantAssociationId`

The ID of the association.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AssociationData
