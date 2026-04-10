This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL CountAction

Specifies that AWS WAF should count the request. Optionally defines
additional custom handling for the request.

This is used in the context of other settings, for example to specify values for a rule
action or a web ACL default action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomRequestHandling" : CustomRequestHandling
}

```

### YAML

```yaml

  CustomRequestHandling:
    CustomRequestHandling

```

## Properties

`CustomRequestHandling`

Defines custom handling for the web request.

For information about customizing web requests and responses,
see [Customizing web requests and responses in AWS WAF](../../../waf/latest/developerguide/waf-custom-request-response.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [CustomRequestHandling](aws-properties-wafv2-webacl-customrequesthandling.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set an count action

The following shows an example count action specification.

#### YAML

```yaml

Action:
  Count: {}

```

#### JSON

```json

"Action": {
  "Count": {}
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Cookies

CustomHTTPHeader

All content copied from https://docs.aws.amazon.com/.
