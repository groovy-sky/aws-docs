# ListResolverQueryLogConfigs

Lists information about the specified query logging configurations. Each configuration defines where you want Resolver to save
DNS query logs and specifies the VPCs that you want to log queries for.

## Request Syntax

```nohighlight

{
   "Filters": [
      {
         "Name": "string",
         "Values": [ "string" ]
      }
   ],
   "MaxResults": number,
   "NextToken": "string",
   "SortBy": "string",
   "SortOrder": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Filters](#API_route53resolver_ListResolverQueryLogConfigs_RequestSyntax)**

An optional specification to return a subset of query logging configurations.

###### Note

If you submit a second or subsequent `ListResolverQueryLogConfigs` request and specify the `NextToken` parameter,
you must use the same values for `Filters`, if any, as in the previous request.

Type: Array of [Filter](api-route53resolver-filter.md) objects

Required: No

**[MaxResults](#API_route53resolver_ListResolverQueryLogConfigs_RequestSyntax)**

The maximum number of query logging configurations that you want to return in the response to a `ListResolverQueryLogConfigs` request.
If you don't specify a value for `MaxResults`, Resolver returns up to 100 query logging configurations.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_route53resolver_ListResolverQueryLogConfigs_RequestSyntax)**

For the first `ListResolverQueryLogConfigs` request, omit this value.

If there are more than `MaxResults` query logging configurations that match the values that you specify for `Filters`,
you can submit another `ListResolverQueryLogConfigs` request to get the next group of configurations. In the next request, specify the value of
`NextToken` from the previous response.

Type: String

Required: No

**[SortBy](#API_route53resolver_ListResolverQueryLogConfigs_RequestSyntax)**

The element that you want Resolver to sort query logging configurations by.

###### Note

If you submit a second or subsequent `ListResolverQueryLogConfigs` request and specify the `NextToken` parameter,
you must use the same value for `SortBy`, if any, as in the previous request.

Valid values include the following elements:

- `Arn`: The ARN of the query logging configuration

- `AssociationCount`: The number of VPCs that are associated with the specified configuration

- `CreationTime`: The date and time that Resolver returned when the configuration was created

- `CreatorRequestId`: The value that was specified for `CreatorRequestId` when the configuration was created

- `DestinationArn`: The location that logs are sent to

- `Id`: The ID of the configuration

- `Name`: The name of the configuration

- `OwnerId`: The AWS account number of the account that created the configuration

- `ShareStatus`: Whether the configuration is shared with other AWS accounts or shared with the current account by
another AWS account. Sharing is configured through AWS Resource Access Manager (AWS RAM).

- `Status`: The current status of the configuration. Valid values include the following:

- `CREATING`: Resolver is creating the query logging configuration.

- `CREATED`: The query logging configuration was successfully created.
Resolver is logging queries that originate in the specified VPC.

- `DELETING`: Resolver is deleting this query logging configuration.

- `FAILED`: Resolver either couldn't create or couldn't delete the query logging configuration.
Here are two common causes:

- The specified destination (for example, an Amazon S3 bucket) was deleted.

- Permissions don't allow sending logs to the destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[SortOrder](#API_route53resolver_ListResolverQueryLogConfigs_RequestSyntax)**

If you specified a value for `SortBy`, the order that you want query logging configurations to be listed in,
`ASCENDING` or `DESCENDING`.

###### Note

If you submit a second or subsequent `ListResolverQueryLogConfigs` request and specify the `NextToken` parameter,
you must use the same value for `SortOrder`, if any, as in the previous request.

Type: String

Valid Values: `ASCENDING | DESCENDING`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "ResolverQueryLogConfigs": [
      {
         "Arn": "string",
         "AssociationCount": number,
         "CreationTime": "string",
         "CreatorRequestId": "string",
         "DestinationArn": "string",
         "Id": "string",
         "Name": "string",
         "OwnerId": "string",
         "ShareStatus": "string",
         "Status": "string"
      }
   ],
   "TotalCount": number,
   "TotalFilteredCount": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_route53resolver_ListResolverQueryLogConfigs_ResponseSyntax)**

If there are more than `MaxResults` query logging configurations, you can submit another `ListResolverQueryLogConfigs` request
to get the next group of configurations. In the next request, specify the value of `NextToken` from the previous response.

Type: String

**[ResolverQueryLogConfigs](#API_route53resolver_ListResolverQueryLogConfigs_ResponseSyntax)**

A list that contains one `ResolverQueryLogConfig` element for each query logging configuration that matches the
values that you specified for `Filter`.

Type: Array of [ResolverQueryLogConfig](api-route53resolver-resolverquerylogconfig.md) objects

**[TotalCount](#API_route53resolver_ListResolverQueryLogConfigs_ResponseSyntax)**

The total number of query logging configurations that were created by the current account in the specified Region. This count can differ from the
number of query logging configurations that are returned in a `ListResolverQueryLogConfigs` response, depending on the values that you specify
in the request.

Type: Integer

**[TotalFilteredCount](#API_route53resolver_ListResolverQueryLogConfigs_ResponseSyntax)**

The total number of query logging configurations that were created by the current account in the specified Region and that match the filters
that were specified in the `ListResolverQueryLogConfigs` request. For the total number of query logging configurations that were created by the
current account in the specified Region, see `TotalCount`.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified Resolver operation.

This error can also be thrown when a customer has reached the 5120 character limit for a
resource policy for CloudWatch Logs.

HTTP Status Code: 400

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**InvalidNextTokenException**

The value that you specified for `NextToken` in a `List` request isn't valid.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters in this request are not valid.

**FieldName**

For an `InvalidParameterException` error, the name of the parameter that's invalid.

HTTP Status Code: 400

**InvalidRequestException**

The request is invalid.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### ListResolverQueryLogConfigs Example

This example illustrates one usage of ListResolverQueryLogConfigs.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 283
X-Amz-Target: Route53Resolver.ListResolverQueryLogConfigs
X-Amz-Date: 20200415T191344Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "Filters": [
        {
            "Name": "ShareStatus",
            "Values": "SHARED_BY_ME"
        }
    ],
    "MaxResults": 10,
    "SortBy": "DestinationArn",
    "SortOrder": "ASCENDING"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 19:13:44 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 531
x-amzn-RequestId: 08afd081-9d67-4281-a277-b3880example
Connection: keep-alive

{
    "MaxResults": 10,
    "ResolverQueryLogConfigs": [
        {
            "Arn": "arn:aws:route53resolver:us-east-1:111122223333:resolver-query-log-config/rqlc-8ca61fe7cexample",
            "AssociationCount": "1",
            "CreationTime": "20200415T191604Z",
            "CreatorRequestId": "ramirezd-20200415T191001Z",
            "DestinationArn": "arn:aws:s3:::amzn-s3-demo-bucket/development/",
            "Id": "rqlc-8ca61fe7c055415a",
            "Name": "MyQueryLog",
            "OwnerId": "111122223333",
            "ShareStatus": "SHARED_BY_ME",
            "Status": "CREATING"
        }
    ],
    "SortBy": "DestinationArn",
    "SortOrder": "ASCENDING",
    "TotalCount": 10,
    "TotalFilteredCount": 1
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/listresolverquerylogconfigs.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListResolverQueryLogConfigAssociations

ListResolverRuleAssociations
