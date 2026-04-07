This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution FunctionAssociation

A CloudFront function that is associated with a cache behavior in a CloudFront
distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventType" : String,
  "FunctionARN" : String
}

```

### YAML

```yaml

  EventType: String
  FunctionARN: String

```

## Properties

`EventType`

The event type of the function, either `viewer-request` or
`viewer-response`. You cannot use origin-facing event types
( `origin-request` and `origin-response`) with a CloudFront
function.

_Required_: No

_Type_: String

_Allowed values_: `viewer-request | viewer-response | origin-request | origin-response`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionARN`

The Amazon Resource Name (ARN) of the function.

_Required_: No

_Type_: String

_Pattern_: `arn:aws:cloudfront::[0-9]{12}:function\/[a-zA-Z0-9-_]{1,64}`

_Minimum_: `0`

_Maximum_: `108`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ForwardedValues

GeoRestriction
