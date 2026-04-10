This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL SingleHeader

Inspect one of the headers in the web request, identified by name, for example,
`User-Agent` or `Referer`. The name isn't case sensitive.

You can filter and inspect all headers with the `FieldToMatch` setting
`Headers`.

This is used to indicate the web request component to inspect, in the [FieldToMatch](../userguide/aws-properties-wafv2-rulegroup-regexpatternsetreferencestatement.md#cfn-wafv2-rulegroup-regexpatternsetreferencestatement-fieldtomatch) specification.

Example JSON: `"SingleHeader": { "Name": "haystack" }`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the query header to inspect.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleGroupReferenceStatement

SingleQueryArgument

All content copied from https://docs.aws.amazon.com/.
