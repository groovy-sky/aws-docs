This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup Tag

A tag associated with an AWS resource. Tags are key:value pairs that you can use to
categorize and manage your resources, for purposes like billing or other management.
Typically, the tag key represents a category, such as "environment", and the tag value
represents a specific value within that category, such as "test," "development," or
"production". Or you might set the tag key to "customer" and the value to the customer name
or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a
resource.

You can tag the AWS resources that you manage through AWS WAF: web ACLs, rule
groups, IP sets, and regex pattern sets. You can't manage or view tags through the AWS WAF
console.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

Part of the key:value pair that defines a tag. You can use a tag key to describe a
category of information, such as "customer." Tag keys are case-sensitive.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Part of the key:value pair that defines a tag. You can use a tag value to describe a
specific value within a category, such as "companyA" or "companyB." Tag values are
case-sensitive.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Statement

TextTransformation
