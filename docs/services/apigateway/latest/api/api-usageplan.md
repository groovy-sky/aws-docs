# UsagePlan

Represents a usage plan used to specify who can assess associated API stages. Optionally, target request rate and quota limits can be set.
In some cases clients can exceed the targets that you set. Don’t rely on usage plans to control costs.
Consider using [AWS Budgets](../../../cost-management/latest/userguide/budgets-managing-costs.md) to monitor costs
and [AWS WAF](../../../waf/latest/developerguide/waf-chapter.md) to manage API requests.

## Contents

**apiStages**

The associated API stages of a usage plan.

Type: Array of [ApiStage](api-apistage.md) objects

Required: No

**description**

The description of a usage plan.

Type: String

Required: No

**id**

The identifier of a UsagePlan resource.

Type: String

Required: No

**name**

The name of a usage plan.

Type: String

Required: No

**productCode**

The AWS Marketplace product identifier to associate with the usage plan as a SaaS product on the AWS Marketplace.

Type: String

Required: No

**quota**

The target maximum number of permitted requests per a given unit time interval.

Type: [QuotaSettings](api-quotasettings.md) object

Required: No

**tags**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

Required: No

**throttle**

A map containing method level throttling information for API stage in a usage plan.

Type: [ThrottleSettings](api-throttlesettings.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/usageplan.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/usageplan.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/usageplan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TlsConfig

UsagePlanKey

All content copied from https://docs.aws.amazon.com/.
