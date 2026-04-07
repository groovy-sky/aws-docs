This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::PhoneNumber

Request an origination phone number for use in your account. For more information on
phone number request see [Request a phone number](https://docs.aws.amazon.com/sms-voice/latest/userguide/phone-numbers-request.html) in the _AWS End User Messaging SMS_
_User Guide_.

###### Note

Registering phone numbers is not supported by AWS CloudFormation. You can import phone numbers and sender IDs that are automatically provisioned at registration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::PhoneNumber",
  "Properties" : {
      "DeletionProtectionEnabled" : Boolean,
      "IsoCountryCode" : String,
      "MandatoryKeywords" : MandatoryKeywords,
      "NumberCapabilities" : [ String, ... ],
      "NumberType" : String,
      "OptionalKeywords" : [ OptionalKeyword, ... ],
      "OptOutListName" : String,
      "SelfManagedOptOutsEnabled" : Boolean,
      "Tags" : [ Tag, ... ],
      "TwoWay" : TwoWay
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::PhoneNumber
Properties:
  DeletionProtectionEnabled: Boolean
  IsoCountryCode: String
  MandatoryKeywords:
    MandatoryKeywords
  NumberCapabilities:
    - String
  NumberType: String
  OptionalKeywords:
    - OptionalKeyword
  OptOutListName: String
  SelfManagedOptOutsEnabled: Boolean
  Tags:
    - Tag
  TwoWay:
    TwoWay

```

## Properties

`DeletionProtectionEnabled`

By default this is set to false. When set to true the phone number can't be
deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsoCountryCode`

The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z]{2}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MandatoryKeywords`

Creates or updates a `MandatoryKeyword` configuration on an origination phone number For more information, see
[Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the AWS End User Messaging SMS User Guide.

_Required_: Yes

_Type_: [MandatoryKeywords](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-phonenumber-mandatorykeywords.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberCapabilities`

Indicates if the phone number will be used for text messages, voice messages, or both.

_Required_: Yes

_Type_: Array of String

_Allowed values_: `SMS | VOICE | MMS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NumberType`

The type of phone number to request.

###### Note

The `ShortCode` number type is not supported in AWS CloudFormation.

_Required_: Yes

_Type_: String

_Allowed values_: `LONG_CODE | TOLL_FREE | TEN_DLC | SIMULATOR`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OptionalKeywords`

A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that
an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number
receives a message that begins with a keyword, AWS End User Messaging SMS responds with a customizable message. Optional keywords are differentiated from
mandatory keywords. For more information, see
[Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the AWS End User Messaging SMS User Guide.

_Required_: No

_Type_: Array of [OptionalKeyword](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-phonenumber-optionalkeyword.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptOutListName`

The name of the OptOutList associated with the phone number.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_:/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedOptOutsEnabled`

When set to false and an end recipient sends a message that begins with HELP or STOP to
one of your dedicated numbers, AWS End User Messaging SMS automatically replies with a
customizable message and adds the end recipient to the OptOutList. When set to true
you're responsible for responding to HELP and STOP requests. You're also responsible for
tracking and honoring opt-out request. For more information see [Self-managed opt-outs](https://docs.aws.amazon.com/sms-voice/latest/userguide/opt-out-list-self-managed.html)

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of tags (key and value pairs) to associate with the requested phone number.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-phonenumber-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TwoWay`

Describes the two-way SMS configuration for a phone number. For more information, see
[Two-way SMS messaging](https://docs.aws.amazon.com/sms-voice/latest/userguide/two-way-sms.html) in the AWS End User Messaging SMS User Guide.

_Required_: No

_Type_: [TwoWay](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-phonenumber-twoway.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `PhoneNumberId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The `PhoneNumber`'s Amazon Resource Name (ARN)

`PhoneNumber`

The phone number in E.164 format.

`PhoneNumberId`

The unique identifier for the phone number.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

MandatoryKeyword
