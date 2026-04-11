This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy ServerTimingHeadersConfig

A configuration for enabling the `Server-Timing` header in HTTP responses
sent from CloudFront.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "SamplingRate" : Number
}

```

### YAML

```yaml

  Enabled: Boolean
  SamplingRate: Number

```

## Properties

`Enabled`

A Boolean that determines whether CloudFront adds the `Server-Timing` header to
HTTP responses that it sends in response to requests that match a cache behavior that's
associated with this response headers policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamplingRate`

A number 0–100 (inclusive) that specifies the percentage of responses that you want
CloudFront to add the `Server-Timing` header to. When you set the sampling rate to
100, CloudFront adds the `Server-Timing` header to the HTTP response for every
request that matches the cache behavior that this response headers policy is attached
to. When you set it to 50, CloudFront adds the header to 50% of the responses for requests
that match the cache behavior. You can set the sampling rate to any number 0–100 with up
to four decimal places.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityHeadersConfig

StrictTransportSecurity

All content copied from https://docs.aws.amazon.com/.
