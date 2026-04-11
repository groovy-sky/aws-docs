This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL GeoMatchStatement

A rule statement that labels web requests by country and region and that matches against web requests based on country code. A geo match rule labels every request that it inspects regardless of whether it finds a match.

- To manage requests only by country, you can use this statement by itself and specify the countries that you want to match against in the `CountryCodes` array.

- Otherwise, configure your geo match rule with Count action so that it only labels requests. Then, add one or more label match rules to run after the geo match rule and configure them to match against the geographic labels and handle the requests as needed.

AWS WAF labels requests using the alpha-2 country and region codes from the International Organization for Standardization (ISO) 3166 standard. AWS WAF determines the codes using either the IP address in the web request origin or, if you specify it, the address in the geo match `ForwardedIPConfig`.

If you use the web request origin, the label formats are `awswaf:clientip:geo:region:<ISO country code>-<ISO region code>` and `awswaf:clientip:geo:country:<ISO country code>`.

If you use a forwarded IP address, the label formats are `awswaf:forwardedip:geo:region:<ISO country code>-<ISO region code>` and `awswaf:forwardedip:geo:country:<ISO country code>`.

For additional details, see [Geographic match rule statement](../../../waf/latest/developerguide/waf-rule-statement-type-geo-match.md) in the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CountryCodes" : [ String, ... ],
  "ForwardedIPConfig" : ForwardedIPConfiguration
}

```

### YAML

```yaml

  CountryCodes:
    - String
  ForwardedIPConfig:
    ForwardedIPConfiguration

```

## Properties

`CountryCodes`

An array of two-character country codes that you want to match against, for example, `[ "US", "CN" ]`, from
the alpha-2 country ISO codes of the ISO 3166 international standard.

When you use a geo match statement just for the region and country labels that it adds to requests, you still have to supply a country code for the rule to evaluate. In this case, you configure the rule to only count matching requests, but it will still generate logging and count metrics for any matches. You can reduce the logging and metrics that the rule produces by specifying a country that's unlikely to be a source of traffic to your site.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedIPConfig`

The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.

###### Note

If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.

_Required_: No

_Type_: [ForwardedIPConfiguration](aws-properties-wafv2-webacl-forwardedipconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ForwardedIPConfiguration

HeaderMatchPattern

All content copied from https://docs.aws.amazon.com/.
