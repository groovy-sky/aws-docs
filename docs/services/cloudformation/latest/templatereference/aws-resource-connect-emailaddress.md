This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EmailAddress

Create new email address in the specified Amazon Connect instance. For more information about email
addresses, see [Create email\
addresses](https://docs.aws.amazon.com/connect/latest/adminguide/create-email-address1.html) in the Amazon Connect Administrator Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::EmailAddress",
  "Properties" : {
      "AliasConfigurations" : [ AliasConfiguration, ... ],
      "Description" : String,
      "DisplayName" : String,
      "EmailAddress" : String,
      "InstanceArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::EmailAddress
Properties:
  AliasConfigurations:
    - AliasConfiguration
  Description: String
  DisplayName: String
  EmailAddress: String
  InstanceArn: String
  Tags:
    - Tag

```

## Properties

`AliasConfigurations`

A list of alias configurations for this email address, showing which email addresses forward to this primary
address. Each configuration contains the email address ID of an alias that forwards emails to this address.

_Required_: No

_Type_: Array of [AliasConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-emailaddress-aliasconfiguration.html)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the email address.

_Required_: No

_Type_: String

_Pattern_: `(^[\S].*[\S]$)|(^[\S]$)`

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of email address.

_Required_: No

_Type_: String

_Pattern_: `(^[\S].*[\S]$)|(^[\S]$)`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailAddress`

The email address, including the domain.

_Required_: Yes

_Type_: String

_Pattern_: `([^\s@]+@[^\s@]+\.[^\s@]+)`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov):connect:[a-z]{2}-[a-z]+-[0-9]{1}:[0-9]{1,20}:instance/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-emailaddress-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the email address. For example:

`{ "Ref": "myEmailAddress" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EmailAddressArn`

The Amazon Resource Name (ARN) of the email address.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Value

AliasConfiguration
