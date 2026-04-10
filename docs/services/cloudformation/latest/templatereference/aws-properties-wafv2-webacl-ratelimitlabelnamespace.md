This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RateLimitLabelNamespace

Specifies a label namespace to use as an aggregate key for a rate-based rule. Each distinct fully qualified label name that has the specified label namespace contributes to the aggregation instance. If you use just one label namespace as your custom key, then each label name fully defines an aggregation instance.

This uses only labels that have been added to the request by rules that are evaluated before this rate-based rule in the web ACL.

For information about label namespaces and names, see
[Label syntax and naming requirements](../../../waf/latest/developerguide/waf-rule-label-requirements.md) in the _AWS WAF Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Namespace" : String
}

```

### YAML

```yaml

  Namespace: String

```

## Properties

`Namespace`

The namespace to use for aggregation.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_:-]{1,1024}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RateLimitJA4Fingerprint

RateLimitQueryArgument

All content copied from https://docs.aws.amazon.com/.
