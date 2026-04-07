This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::SenderId

Request a new sender ID that doesn't require registration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::SenderId",
  "Properties" : {
      "DeletionProtectionEnabled" : Boolean,
      "IsoCountryCode" : String,
      "SenderId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::SenderId
Properties:
  DeletionProtectionEnabled: Boolean
  IsoCountryCode: String
  SenderId: String
  Tags:
    - Tag

```

## Properties

`DeletionProtectionEnabled`

By default this is set to false. When set to true the sender ID can't be deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsoCountryCode`

The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z]{2}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SenderId`

The sender ID string to request.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z0-9_-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of tags (key and value pairs) to associate with the sender ID.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-senderid-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `IsoCountryCode|SenderId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name of the `SenderId`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SMSVOICE::ResourcePolicy

Tag
