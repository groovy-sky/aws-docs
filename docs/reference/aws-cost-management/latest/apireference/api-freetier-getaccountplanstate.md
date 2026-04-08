# GetAccountPlanState

This returns all of the information related to the state of the account plan related to Free Tier.

## Response Syntax

```nohighlight

{
   "accountId": "string",
   "accountPlanExpirationDate": "string",
   "accountPlanRemainingCredits": {
      "amount": number,
      "unit": "string"
   },
   "accountPlanStatus": "string",
   "accountPlanType": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[accountId](#API_freetier_GetAccountPlanState_ResponseSyntax)**

A unique identifier that identifies the account.

Type: String

Pattern: `[0-9]{12}`

**[accountPlanExpirationDate](#API_freetier_GetAccountPlanState_ResponseSyntax)**

The timestamp for when the current account plan expires.

Type: Timestamp

**[accountPlanRemainingCredits](#API_freetier_GetAccountPlanState_ResponseSyntax)**

The amount of credits remaining for the account.

Type: [MonetaryAmount](api-freetier-monetaryamount.md) object

**[accountPlanStatus](#API_freetier_GetAccountPlanState_ResponseSyntax)**

The current status for the account plan.

Type: String

Valid Values: `NOT_STARTED | ACTIVE | EXPIRED`

**[accountPlanType](#API_freetier_GetAccountPlanState_ResponseSyntax)**

The plan type for the account.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/freetier-2023-09-07/getaccountplanstate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/freetier-2023-09-07/getaccountplanstate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetAccountActivity

GetFreeTierUsage
