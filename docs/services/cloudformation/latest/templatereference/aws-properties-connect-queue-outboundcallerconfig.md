This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Queue OutboundCallerConfig

The outbound caller ID name, number, and outbound whisper flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutboundCallerIdName" : String,
  "OutboundCallerIdNumberArn" : String,
  "OutboundFlowArn" : String
}

```

### YAML

```yaml

  OutboundCallerIdName: String
  OutboundCallerIdNumberArn: String
  OutboundFlowArn: String

```

## Properties

`OutboundCallerIdName`

The caller ID name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundCallerIdNumberArn`

The Amazon Resource Name (ARN) of the outbound caller ID number.

###### Note

Only use the phone number ARN format that doesn't contain `instance` in
the path, for example,
`arn:aws:connect:us-east-1:1234567890:phone-number/uuid`. This is the
same ARN format that is returned when you create a phone number using CloudFormation, or when you call the [ListPhoneNumbersV2](../../../../reference/connect/latest/apireference/api-listphonenumbersv2.md) API.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:phone-number/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundFlowArn`

The Amazon Resource Name (ARN) of the outbound flow.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmailAddress

OutboundEmailConfig

All content copied from https://docs.aws.amazon.com/.
