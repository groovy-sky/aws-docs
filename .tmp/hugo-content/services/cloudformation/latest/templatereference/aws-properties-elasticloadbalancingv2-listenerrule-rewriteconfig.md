This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule RewriteConfig

Information about a rewrite transform. This transform matches a pattern and replaces it with the specified string.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Regex" : String,
  "Replace" : String
}

```

### YAML

```yaml

  Regex: String
  Replace: String

```

## Properties

`Regex`

The regular expression to match in the input string. The maximum length of the string is 1,024 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Replace`

The replacement string to use when rewriting the matched input. The maximum length of the string is 1,024 characters.
You can specify capture groups in the regular expression (for example, $1 and $2).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedirectConfig

RewriteConfigObject

All content copied from https://docs.aws.amazon.com/.
