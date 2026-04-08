Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)

## EndpointProviderV2        in package    - [Aws](package-aws.md)

Given a service's Ruleset and client-provided input parameters, provides
either an object reflecting the properties of a resolved endpoint,
or throws an error.

### Table of Contents  [header link](class-aws-endpointv2-endpointproviderv2-toc.md)

#### Methods  [header link](class-aws-endpointv2-endpointproviderv2-toc-methods.md)

[\_\_construct()](class-aws-endpointv2-endpointproviderv2-method-construct.md)
: mixed [getRuleset()](class-aws-endpointv2-endpointproviderv2-method-getruleset.md)
: [Ruleset](class-aws-endpointv2-ruleset-ruleset.md)[resolveEndpoint()](class-aws-endpointv2-endpointproviderv2-method-resolveendpoint.md)
: [RulesetEndpoint](class-aws-endpointv2-ruleset-rulesetendpoint.md)Given a Ruleset and input parameters, determines the correct endpoint
or an error to be thrown for a given request.

### Methods  [header link](class-aws-endpointv2-endpointproviderv2-methods.md)

#### \_\_construct()  [header link](class-aws-endpointv2-endpointproviderv2-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $ruleset, array<string|int, mixed> $partitions) : mixed`

##### Parameters

$ruleset
: array<string\|int, mixed>$partitions
: array<string\|int, mixed>

#### getRuleset()  [header link](class-aws-endpointv2-endpointproviderv2-method-getruleset.md)

`
    public
                    getRuleset() : Ruleset`

##### Return values

[Ruleset](class-aws-endpointv2-ruleset-ruleset.md)

#### resolveEndpoint()  [header link](class-aws-endpointv2-endpointproviderv2-method-resolveendpoint.md)

Given a Ruleset and input parameters, determines the correct endpoint
or an error to be thrown for a given request.

`
    public
                    resolveEndpoint(array<string|int, mixed> $inputParameters) : RulesetEndpoint`

##### Parameters

$inputParameters
: array<string\|int, mixed>

##### Tags  [header link](class-aws-endpointv2-endpointproviderv2-method-resolveendpoint-tags.md)

throws[UnresolvedEndpointException](class-aws-exception-unresolvedendpointexception.md)

##### Return values

[RulesetEndpoint](class-aws-endpointv2-ruleset-rulesetendpoint.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-endpointv2-endpointproviderv2-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointv2-endpointproviderv2-method-construct.md)
  - [getRuleset()](class-aws-endpointv2-endpointproviderv2-method-getruleset.md)
  - [resolveEndpoint()](class-aws-endpointv2-endpointproviderv2-method-resolveendpoint.md)

[Back To Top](class-aws-endpointv2-endpointproviderv2-top.md)
