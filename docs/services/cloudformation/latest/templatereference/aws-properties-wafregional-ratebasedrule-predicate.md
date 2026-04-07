This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::RateBasedRule Predicate

Specifies the `ByteMatchSet`, `IPSet`, `SqlInjectionMatchSet`, `XssMatchSet`, `RegexMatchSet`, `GeoMatchSet`, and `SizeConstraintSet` objects
that you want to add to a `Rule` and, for each object, indicates whether you want to negate the settings, for example, requests that do
NOT originate from the IP address 192.0.2.44.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataId" : String,
  "Negated" : Boolean,
  "Type" : String
}

```

### YAML

```yaml

  DataId: String
  Negated: Boolean
  Type: String

```

## Properties

`DataId`

A unique identifier for a predicate in a `Rule`, such as `ByteMatchSetId` or `IPSetId`.
The ID is returned by the corresponding `Create` or `List` command.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Negated`

Set `Negated` to `False` if you want AWS WAF to allow, block, or count requests based on the settings in the
specified `ByteMatchSet`, `IPSet`, `SqlInjectionMatchSet`, `XssMatchSet`, `RegexMatchSet`, `GeoMatchSet`, or `SizeConstraintSet`.
For example, if an `IPSet` includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.

Set `Negated` to `True` if you want AWS WAF to allow or block a request based on the negation
of the settings in the `ByteMatchSet`, `IPSet`, `SqlInjectionMatchSet`, `XssMatchSet`, `RegexMatchSet`, `GeoMatchSet`, or `SizeConstraintSet` >.
For example, if an `IPSet` includes the IP address `192.0.2.44`, AWS WAF will allow, block, or count requests based on
all IP addresses _except_ `192.0.2.44`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of predicate in a `Rule`, such as `ByteMatch` or `IPSet`.

_Required_: Yes

_Type_: String

_Allowed values_: `IPMatch | ByteMatch | SqlInjectionMatch | GeoMatch | SizeConstraint | XssMatch | RegexMatch`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::WAFRegional::RateBasedRule

AWS::WAFRegional::RegexPatternSet
