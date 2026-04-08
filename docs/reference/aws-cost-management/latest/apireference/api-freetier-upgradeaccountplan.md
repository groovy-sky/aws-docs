# UpgradeAccountPlan

The account plan type for the AWS account.

## Request Syntax

```nohighlight

{
   "accountPlanType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[accountPlanType](#API_freetier_UpgradeAccountPlan_RequestSyntax)**

The target account plan type. This makes it explicit about the change and latest value of the `accountPlanType`.

Type: String

Valid Values: `FREE | PAID`

Required: Yes

## Response Syntax

```nohighlight

{
   "accountId": "string",
   "accountPlanStatus": "string",
   "accountPlanType": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[accountId](#API_freetier_UpgradeAccountPlan_ResponseSyntax)**

A unique identifier that identifies the account.

Type: String

Pattern: `[0-9]{12}`

**[accountPlanStatus](#API_freetier_UpgradeAccountPlan_ResponseSyntax)**

This indicates the latest state of the account plan within its lifecycle.

Type: String

Valid Values: `NOT_STARTED | ACTIVE | EXPIRED`

**[accountPlanType](#API_freetier_UpgradeAccountPlan_ResponseSyntax)**

The type of plan for the account.

Type: String

Valid Values: `FREE | PAID`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient access to perform this action.

HTTP Status Code: 400

**InternalServerException**

An unexpected error occurred during the processing of your request.

HTTP Status Code: 500

**ResourceNotFoundException**

This exception is thrown when the requested resource cannot be found.

HTTP Status Code: 400

**ThrottlingException**

The request was denied due to request throttling.

HTTP Status Code: 400

**ValidationException**

The input fails to satisfy the constraints specified by an AWS service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for Python](../../../../services/goto/boto3/freetier-2023-09-07/upgradeaccountplan.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/freetier-2023-09-07/upgradeaccountplan.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListAccountActivities

AWS Invoicing
