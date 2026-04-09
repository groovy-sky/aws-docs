# UpdateSubscription

Updates the pricing tier for an Amazon Q Business subscription. Upgrades are instant.
Downgrades apply at the start of the next month. Subscription tier determines feature
access for the user. For more information on subscriptions and pricing tiers, see [Amazon Q Business\
pricing](https://aws.amazon.com/q/business/pricing).

## Request Syntax

```nohighlight

PUT /applications/applicationId/subscriptions/subscriptionId HTTP/1.1
Content-type: application/json

{
   "type": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_UpdateSubscription_RequestSyntax)**

The identifier of the Amazon Q Business application where the subscription update should
take effect.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[subscriptionId](#API_UpdateSubscription_RequestSyntax)**

The identifier of the Amazon Q Business subscription to be updated.

Length Constraints: Minimum length of 0. Maximum length of 1224.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[type](#API_UpdateSubscription_RequestSyntax)**

The type of the Amazon Q Business subscription to be updated.

Type: String

Valid Values: `Q_LITE | Q_BUSINESS`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "currentSubscription": {
      "type": "string"
   },
   "nextSubscription": {
      "type": "string"
   },
   "subscriptionArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[currentSubscription](#API_UpdateSubscription_ResponseSyntax)**

The type of your current Amazon Q Business subscription.

Type: [SubscriptionDetails](api-subscriptiondetails.md) object

**[nextSubscription](#API_UpdateSubscription_ResponseSyntax)**

The type of the Amazon Q Business subscription for the next month.

Type: [SubscriptionDetails](api-subscriptiondetails.md) object

**[subscriptionArn](#API_UpdateSubscription_ResponseSyntax)**

The Amazon Resource Name (ARN) of the Amazon Q Business subscription that was
updated.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**ConflictException**

You are trying to perform an action that conflicts with the current status of your
resource. Fix any inconsistencies with your resources and try again.

**message**

The message describing a `ConflictException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 409

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/updatesubscription.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/updatesubscription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateRetriever

UpdateUser

All content copied from https://docs.aws.amazon.com/.
