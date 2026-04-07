This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::OptOutList

Creates a new opt-out list.

If the opt-out list name already exists, an error is returned.

An opt-out list is a list of phone numbers that are opted out, meaning you can't send
SMS or voice messages to them. If end user replies with the keyword "STOP," an entry for
the phone number is added to the opt-out list. In addition to STOP, your recipients can
use any supported opt-out keyword, such as CANCEL or OPTOUT. For a list of supported
opt-out keywords, see [SMS opt out](https://docs.aws.amazon.com/sms-voice/latest/userguide/opt-out-list-keywords.html) in the AWS End User Messaging SMS User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::OptOutList",
  "Properties" : {
      "OptOutListName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::OptOutList
Properties:
  OptOutListName: String
  Tags:
    - Tag

```

## Properties

`OptOutListName`

The name of the OptOutList.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of tags (key and value pairs) to associate with the new OptOutList.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-optoutlist-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `OptOutListName`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the `OptOutList`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
