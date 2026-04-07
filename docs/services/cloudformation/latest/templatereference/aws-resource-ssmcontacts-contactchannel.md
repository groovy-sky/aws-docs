This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::ContactChannel

The `AWS::SSMContacts::ContactChannel` resource specifies a contact channel
as the method that Incident Manager uses to engage your contact.

###### Note

**Template example**: We recommend creating all Incident Manager `Contacts` resources using a single AWS CloudFormation template. For a
demonstration, see the examples for [AWS::SSMContacts::Contacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMContacts::ContactChannel",
  "Properties" : {
      "ChannelAddress" : String,
      "ChannelName" : String,
      "ChannelType" : String,
      "ContactId" : String,
      "DeferActivation" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::SSMContacts::ContactChannel
Properties:
  ChannelAddress: String
  ChannelName: String
  ChannelType: String
  ContactId: String
  DeferActivation: Boolean

```

## Properties

`ChannelAddress`

The details that Incident Manager uses when trying to engage the contact channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelName`

The name of the contact channel.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z 0-9_\-+'&amp;\uD83C-\uDBFF\uDC00-\uDFFF\u2000-\u3300]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelType`

The type of the contact channel. Incident Manager supports three contact methods:

- SMS

- VOICE

- EMAIL

_Required_: No

_Type_: String

_Allowed values_: `SMS | VOICE | EMAIL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContactId`

The Amazon Resource Name (ARN) of the contact you are adding the contact channel
to.

_Required_: No

_Type_: String

_Pattern_: `arn:[-\w+=\/,.@]+:[-\w+=\/,.@]+:[-\w+=\/,.@]*:[0-9]+:([\w+=\/,.@:-]+)*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeferActivation`

If you want to activate the channel at a later time, you can choose to defer activation.
Incident Manager can't engage your contact channel until it has been activated.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the resource, such as
`arn:aws:ssm-contacts:us-west-2:123456789012:contact-channel/contactalias/cec1bb12-34f5-6789-a1ee-e1ca2345d6f7`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `ContactChannel` resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Targets

AWS::SSMContacts::Plan
