This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::InboundExternalLink LinkAttributes

The `LinkAttributes` property type specifies Property description not available. for an [AWS::RTBFabric::InboundExternalLink](aws-resource-rtbfabric-inboundexternallink.md).

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

_Type_: Array of [ResponderErrorMaskingForHttpCode](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rtbfabric-inboundexternallink-respondererrormaskingforhttpcode.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LinkApplicationLogSampling

LinkLogSettings
