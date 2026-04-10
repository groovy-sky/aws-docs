This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup HeaderOrder

Inspect a string containing the list of the request's header names, ordered as they appear in the web request
that AWS WAF receives for inspection.
AWS WAF generates the string and then uses that as the field to match component in its inspection.
AWS WAF separates the header names in the string using colons and no added spaces, for example `host:user-agent:accept:authorization:referer`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OversizeHandling" : String
}

```

### YAML

```yaml

  OversizeHandling: String

```

## Properties

`OversizeHandling`

What AWS WAF should do if the headers determined by your match scope are more numerous or larger than AWS WAF can inspect.
AWS WAF does not support inspecting the entire contents of request headers
when they exceed 8 KB (8192 bytes) or 200 total headers. The underlying host service forwards a maximum of 200 headers
and at most 8 KB of header contents to AWS WAF.

The options for oversize handling are the following:

- `CONTINUE` \- Inspect the available headers normally, according to the rule inspection criteria.

- `MATCH` \- Treat the web request as matching the rule statement. AWS WAF
applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule
statement.

_Required_: Yes

_Type_: String

_Allowed values_: `CONTINUE | MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeaderMatchPattern

Headers

All content copied from https://docs.aws.amazon.com/.
