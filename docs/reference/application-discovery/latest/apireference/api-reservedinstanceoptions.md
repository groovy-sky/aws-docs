# ReservedInstanceOptions

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Used to provide Reserved Instance preferences for the recommendation.

## Contents

**offeringClass**

The flexibility to change the instance types needed for your Reserved Instance.

Type: String

Valid Values: `STANDARD | CONVERTIBLE`

Required: Yes

**purchasingOption**

The payment plan to use for your Reserved Instance.

Type: String

Valid Values: `ALL_UPFRONT | PARTIAL_UPFRONT | NO_UPFRONT`

Required: Yes

**termLength**

The preferred duration of the Reserved Instance term.

Type: String

Valid Values: `ONE_YEAR | THREE_YEAR`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/reservedinstanceoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/reservedinstanceoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/reservedinstanceoptions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrderByElement

Tag

All content copied from https://docs.aws.amazon.com/.
