This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Plan

Information about the stages and on-call rotation teams associated with an escalation
plan or engagement plan.

###### Note

**Template example**: We recommend creating all Incident Manager `Contacts` resources using a single AWS CloudFormation template. For a
demonstration, see the examples for [AWS::SSMContacts::Contacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMContacts::Plan",
  "Properties" : {
      "ContactId" : String,
      "RotationIds" : [ String, ... ],
      "Stages" : [ Stage, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSMContacts::Plan
Properties:
  ContactId: String
  RotationIds:
    - String
  Stages:
    - Stage

```

## Properties

`ContactId`

The Amazon Resource Name (ARN) of the contact.

_Required_: No

_Type_: String

_Pattern_: `arn:[-\w+=\/,.@]+:[-\w+=\/,.@]+:[-\w+=\/,.@]*:[0-9]+:([\w+=\/,.@:-]+)*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RotationIds`

The Amazon Resource Names (ARNs) of the on-call rotations associated with the
plan.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stages`

A list of stages that the escalation plan or engagement plan uses to engage contacts and
contact methods.

_Required_: No

_Type_: Array of [Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmcontacts-plan-stage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `Plan` resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SSMContacts::ContactChannel

ChannelTargetInfo
