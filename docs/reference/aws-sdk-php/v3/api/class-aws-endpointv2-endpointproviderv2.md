Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)

## EndpointProviderV2        in package    - [Aws](package-aws.md)

Given a service's Ruleset and client-provided input parameters, provides
either an object reflecting the properties of a resolved endpoint,
or throws an error.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#method___construct)
: mixed [getRuleset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#method_getRuleset)
: [Ruleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html)[resolveEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#method_resolveEndpoint)
: [RulesetEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.RulesetEndpoint.html)Given a Ruleset and input parameters, determines the correct endpoint
or an error to be thrown for a given request.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $ruleset, array<string|int, mixed> $partitions) : mixed`

##### Parameters

$ruleset
: array<string\|int, mixed>$partitions
: array<string\|int, mixed>

#### getRuleset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#method_getRuleset)

`
    public
                    getRuleset() : Ruleset`

##### Return values

[Ruleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html)

#### resolveEndpoint()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#method_resolveEndpoint)

Given a Ruleset and input parameters, determines the correct endpoint
or an error to be thrown for a given request.

`
    public
                    resolveEndpoint(array<string|int, mixed> $inputParameters) : RulesetEndpoint`

##### Parameters

$inputParameters
: array<string\|int, mixed>

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html\#method_resolveEndpoint\#tags)

throws[UnresolvedEndpointException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.UnresolvedEndpointException.html)

##### Return values

[RulesetEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.RulesetEndpoint.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#method___construct)
  - [getRuleset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#method_getRuleset)
  - [resolveEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#method_resolveEndpoint)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.EndpointProviderV2.html#top)
