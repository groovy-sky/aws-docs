This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::Pool

Creates a new pool and associates the specified origination identity to the pool. A
pool can include one or more phone numbers and SenderIds that are associated with your
AWS account.

The new pool inherits its configuration from the specified origination identity. This
includes keywords, message type, opt-out list, two-way configuration, and self-managed
opt-out configuration. Deletion protection isn't inherited from the origination identity
and defaults to false.

If the origination identity is a phone number and is already associated with another
pool, an error is returned. A sender ID can be associated with multiple pools.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::Pool",
  "Properties" : {
      "DeletionProtectionEnabled" : Boolean,
      "MandatoryKeywords" : MandatoryKeywords,
      "OptionalKeywords" : [ OptionalKeyword, ... ],
      "OptOutListName" : String,
      "OriginationIdentities" : [ String, ... ],
      "SelfManagedOptOutsEnabled" : Boolean,
      "SharedRoutesEnabled" : Boolean,
      "Tags" : [ Tag, ... ],
      "TwoWay" : TwoWay
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::Pool
Properties:
  DeletionProtectionEnabled: Boolean
  MandatoryKeywords:
    MandatoryKeywords
  OptionalKeywords:
    - OptionalKeyword
  OptOutListName: String
  OriginationIdentities:
    - String
  SelfManagedOptOutsEnabled: Boolean
  SharedRoutesEnabled: Boolean
  Tags:
    - Tag
  TwoWay:
    TwoWay

```

## Properties

`DeletionProtectionEnabled`

When set to true the pool can't be deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MandatoryKeywords`

Creates or updates the pool's `MandatoryKeyword` configuration. For more information, see
[Keywords](../../../sms-voice/latest/userguide/keywords.md) in the AWS End User Messaging SMS User Guide.

_Required_: Yes

_Type_: [MandatoryKeywords](aws-properties-smsvoice-pool-mandatorykeywords.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionalKeywords`

Specifies any optional keywords to associate with the pool. For more information, see
[Keywords](../../../sms-voice/latest/userguide/keywords.md) in the AWS End User Messaging SMS User Guide.

_Required_: No

_Type_: Array of [OptionalKeyword](aws-properties-smsvoice-pool-optionalkeyword.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptOutListName`

The name of the OptOutList associated with the pool.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_:/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginationIdentities`

The list of origination identities to apply to the pool, either `PhoneNumberArn` or
`SenderIdArn`. For more information, see [Registrations](../../../sms-voice/latest/userguide/registrations.md)
in the AWS End User Messaging SMS User Guide.

###### Important

If you are using a shared AWS End User Messaging SMS resource then you must use the full Amazon Resource Name (ARN).

_Required_: Yes

_Type_: Array of String

_Maximum_: `256`

_Minimum_: `20 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelfManagedOptOutsEnabled`

When set to false, an end recipient sends a message that begins with HELP or STOP to
one of your dedicated numbers, AWS End User Messaging SMS automatically replies with a
customizable message and adds the end recipient to the OptOutList. When set to true
you're responsible for responding to HELP and STOP requests. You're also responsible for
tracking and honoring opt-out requests. For more information see [Self-managed opt-outs](../../../pinpoint/latest/userguide/settings-sms-managing.md#settings-account-sms-self-managed-opt-out)

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharedRoutesEnabled`

Allows you to enable shared routes on your pool.

By default, this is set to `False`. If you set this value to
`True`, your messages are sent using phone numbers or sender IDs
(depending on the country) that are shared with other users. In some
countries, such as the United States, senders aren't allowed to use shared routes and
must use a dedicated phone number or short code.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of tags (key and value pairs) associated with the pool.

_Required_: No

_Type_: Array of [Tag](aws-properties-smsvoice-pool-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TwoWay`

Describes the two-way SMS configuration for a phone number. For more information, see
[Two-way SMS messaging](../../../sms-voice/latest/userguide/two-way-sms.md) in the AWS End User Messaging SMS User Guide.

_Required_: No

_Type_: [TwoWay](aws-properties-smsvoice-pool-twoway.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `PoolId`.

### Fn::GetAtt

`Arn`

The Amazon Resource Name of the `Pool`.

`PoolId`

The unique identifier for the pool.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TwoWay

MandatoryKeyword

All content copied from https://docs.aws.amazon.com/.
