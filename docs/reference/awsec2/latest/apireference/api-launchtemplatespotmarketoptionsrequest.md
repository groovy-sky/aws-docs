# LaunchTemplateSpotMarketOptionsRequest

The options for Spot Instances.

## Contents

**BlockDurationMinutes**

Deprecated.

Type: Integer

Required: No

**InstanceInterruptionBehavior**

The behavior when a Spot Instance is interrupted. The default is
`terminate`.

Type: String

Valid Values: `hibernate | stop | terminate`

Required: No

**MaxPrice**

The maximum hourly price you're willing to pay for a Spot Instance. We do not
recommend using this parameter because it can lead to increased interruptions. If you do
not specify this parameter, you will pay the current Spot price. If you do specify this
parameter, it must be more than USD $0.001. Specifying a value below USD $0.001 will
result in an `InvalidParameterValue` error message when the launch template
is used to launch an instance.

###### Important

If you specify a maximum price, your Spot Instances will be interrupted more
frequently than if you do not specify this parameter.

Type: String

Required: No

**SpotInstanceType**

The Spot Instance request type.

Type: String

Valid Values: `one-time | persistent`

Required: No

**ValidUntil**

The end date of the request, in UTC format
( _YYYY-MM-DD_ T _HH:MM:SS_ Z). Supported only for
persistent requests.

- For a persistent request, the request remains active until the
`ValidUntil` date and time is reached. Otherwise, the request
remains active until you cancel it.

- For a one-time request, `ValidUntil` is not supported. The request
remains active until all instances launch or you cancel the request.

Default: 7 days from the current date

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatespotmarketoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatespotmarketoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatespotmarketoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateSpotMarketOptions

LaunchTemplateTagSpecification
