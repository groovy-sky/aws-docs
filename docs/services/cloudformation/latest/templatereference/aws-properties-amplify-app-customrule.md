This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::App CustomRule

The CustomRule property type allows you to specify redirects, rewrites, and reverse
proxies. Redirects enable a web app to reroute navigation from one URL to another.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String,
  "Source" : String,
  "Status" : String,
  "Target" : String
}

```

### YAML

```yaml

  Condition: String
  Source: String
  Status: String
  Target: String

```

## Properties

`Condition`

The condition for a URL rewrite or redirect rule, such as a country code.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source pattern for a URL rewrite or redirect rule.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status code for a URL rewrite or redirect rule.

200

Represents a 200 rewrite rule.

301

Represents a 301 (moved pemanently) redirect rule. This and all future requests
should be directed to the target URL.

302

Represents a 302 temporary redirect rule.

404

Represents a 404 redirect rule.

404-200

Represents a 404 rewrite rule.

_Required_: No

_Type_: String

_Pattern_: `.{3,7}`

_Minimum_: `3`

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The target pattern for a URL rewrite or redirect rule.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CacheConfig

EnvironmentVariable
