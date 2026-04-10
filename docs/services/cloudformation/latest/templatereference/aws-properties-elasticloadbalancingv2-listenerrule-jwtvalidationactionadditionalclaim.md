This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule JwtValidationActionAdditionalClaim

Information about an additional claim to validate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Format" : String,
  "Name" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Format: String
  Name: String
  Values:
    - String

```

## Properties

`Format`

The format of the claim value.

_Required_: Yes

_Type_: String

_Allowed values_: `single-string | string-array | space-separated-values`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the claim. You can't specify `exp`, `iss`,
`nbf`, or `iat` because we validate them by default.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The claim value. The maximum size of the list is 10.
Each value can be up to 256 characters in length.
If the format is `space-separated-values`, the values
can't include spaces.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpRequestMethodConfig

JwtValidationConfig

All content copied from https://docs.aws.amazon.com/.
