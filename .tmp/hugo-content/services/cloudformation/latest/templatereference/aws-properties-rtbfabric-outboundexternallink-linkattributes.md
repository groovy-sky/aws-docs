This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::OutboundExternalLink LinkAttributes

The `LinkAttributes` property type specifies Property description not available. for an [AWS::RTBFabric::OutboundExternalLink](aws-resource-rtbfabric-outboundexternallink.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerProvidedId" : String,
  "ResponderErrorMasking" : [ ResponderErrorMaskingForHttpCode, ... ]
}

```

### YAML

```yaml

  CustomerProvidedId: String
  ResponderErrorMasking:
    - ResponderErrorMaskingForHttpCode

```

## Properties

`CustomerProvidedId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ResponderErrorMasking`

Property description not available.

_Required_: No

_Type_: Array of [ResponderErrorMaskingForHttpCode](aws-properties-rtbfabric-outboundexternallink-respondererrormaskingforhttpcode.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LinkApplicationLogSampling

LinkLogSettings

All content copied from https://docs.aws.amazon.com/.
