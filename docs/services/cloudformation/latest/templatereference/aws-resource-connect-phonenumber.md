This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::PhoneNumber

Claims a phone number to the specified Amazon Connect instance or traffic
distribution group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::PhoneNumber",
  "Properties" : {
      "CountryCode" : String,
      "Description" : String,
      "Prefix" : String,
      "SourcePhoneNumberArn" : String,
      "Tags" : [ Tag, ... ],
      "TargetArn" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::PhoneNumber
Properties:
  CountryCode: String
  Description: String
  Prefix: String
  SourcePhoneNumberArn: String
  Tags:
    - Tag
  TargetArn: String
  Type: String

```

## Properties

`CountryCode`

The ISO country code.

_Required_: No

_Type_: String

_Pattern_: `^[A-Z]{2}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the phone number.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix of the phone number. If provided, it must contain `+` as part of
the country code.

_Pattern_: `^\\+[0-9]{1,15}`

_Required_: No

_Type_: String

_Pattern_: `^\+[0-9]{1,15}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePhoneNumberArn`

The claimed phone number ARN that was previously imported from the external service, such as AWS
End User Messaging. If it is from AWS End User Messaging, it looks like the ARN of the phone number
that was imported from AWS End User Messaging.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource. For example, {
"tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-phonenumber-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The Amazon Resource Name (ARN) for Amazon Connect instances or traffic
distribution group that phone numbers are claimed to.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:(instance|traffic-distribution-group)/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of phone number.

_Required_: No

_Type_: String

_Pattern_: `TOLL_FREE|DID|UIFN|SHARED|THIRD_PARTY_DID|THIRD_PARTY_TF|SHORT_CODE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the phone number. For example:

`{ "Ref": "myPhoneNumber" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Address`

The phone number, in E.164 format.

`PhoneNumberArn`

The Amazon Resource Name (ARN) of the phone number.

## Examples

### Specify a phone number resource

The following example specifies a phone number resource for an Amazon Connect instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
 Description: Specifies a phone number for Amazon Connect instance
 Resources:
   PhoneNumber:
     Type: 'AWS::Connect::PhoneNumber'
     Properties:
       TargetArn: arn:aws:connect:region-name:aws-account-id:instance/instance-arn
       Description: phone number created using cfn
       Type: DID
       CountryCode: US
       Tags:
         - Key: testkey
           Value: testValue
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
