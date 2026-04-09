# ListGroupResources

Returns a list of Amazon resource names (ARNs) of the resources that are members of a specified resource
group.

**Minimum permissions**

To run this command, you must have the following permissions:

- `resource-groups:ListGroupResources`

- `cloudformation:DescribeStacks`

- `cloudformation:ListStackResources`

- `tag:GetResources`

## Request Syntax

```nohighlight

POST /list-group-resources HTTP/1.1
Content-type: application/json

{
   "Filters": [
      {
         "Name": "string",
         "Values": [ "string" ]
      }
   ],
   "Group": "string",
   "GroupName": "string",
   "MaxResults": number,
   "NextToken": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[Filters](#API_ListGroupResources_RequestSyntax)**

Filters, formatted as [ResourceFilter](api-resourcefilter.md) objects, that you want to apply
to a `ListGroupResources` operation. Filters the results to include only
those of the specified resource types.

- `resource-type` \- Filter resources by their type. Specify up to
five resource types in the format `AWS::ServiceCode::ResourceType`.
For example, `AWS::EC2::Instance`, or `AWS::S3::Bucket`.

When you specify a `resource-type` filter for
`ListGroupResources`, AWS Resource Groups validates your filter resource types
against the types that are defined in the query associated with the group. For example,
if a group contains only S3 buckets because its query specifies only that resource type,
but your `resource-type` filter includes EC2 instances, AWS Resource Groups does not
filter for EC2 instances. In this case, a `ListGroupResources` request
returns a `BadRequestException` error with a message similar to the
following:

`The resource types specified as filters in the request are not
            valid.`

The error includes a list of resource types that failed the validation because they
are not part of the query associated with the group. This validation doesn't occur when
the group query specifies `AWS::AllSupported`, because a group based on such
a query can contain any of the allowed resource types for the query type (tag-based or
Amazon CloudFront stack-based queries).

Type: Array of [ResourceFilter](api-resourcefilter.md) objects

Required: No

**[Group](#API_ListGroupResources_RequestSyntax)**

The name or the Amazon resource name (ARN) of the resource group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[a-zA-Z0-9_\.-]{1,300}|[a-zA-Z0-9_\.-]{1,150}/[a-z0-9]{26}|arn:aws(-[a-z]+)*:resource-groups:[a-z]{2}(-[a-z]+)+-\d{1}:[0-9]{12}:group/([a-zA-Z0-9_\.-]{1,300}|[a-zA-Z0-9_\.-]{1,150}/[a-z0-9]{26})`

Required: No

**[GroupName](#API_ListGroupResources_RequestSyntax)**

###### Important

_**Deprecated - don't use this parameter. Use the**_
_**`Group` request field instead.**_

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `[a-zA-Z0-9_\.-]{1,300}|[a-zA-Z0-9_\.-]{1,150}/[a-z0-9]{26}`

Required: No

**[MaxResults](#API_ListGroupResources_RequestSyntax)**

The total number of results that you want included on each page of the
response. If you do not include this parameter, it defaults to a value that is specific to the
operation. If additional items exist beyond the maximum you specify, the `NextToken`
response element is present and has a value (is not null). Include that value as the
`NextToken` request parameter in the next call to the operation to get the next part
of the results. Note that the service might return fewer results than the maximum even when there
are more results available. You should check `NextToken` after every operation to
ensure that you receive all of the results.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[NextToken](#API_ListGroupResources_RequestSyntax)**

The parameter for receiving additional results if you receive a
`NextToken` response in a previous request. A `NextToken` response
indicates that more output is available. Set this parameter to the value provided by a previous
call's `NextToken` response to indicate where the output should continue from.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 8192.

Pattern: `^[a-zA-Z0-9+/]*={0,2}$`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "NextToken": "string",
   "QueryErrors": [
      {
         "ErrorCode": "string",
         "Message": "string"
      }
   ],
   "ResourceIdentifiers": [
      {
         "ResourceArn": "string",
         "ResourceType": "string"
      }
   ],
   "Resources": [
      {
         "Identifier": {
            "ResourceArn": "string",
            "ResourceType": "string"
         },
         "Status": {
            "Name": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListGroupResources_ResponseSyntax)**

If present, indicates that more output is available than is
included in the current response. Use this value in the `NextToken` request parameter
in a subsequent call to the operation to get the next part of the output. You should repeat this
until the `NextToken` response element comes back as `null`.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 8192.

Pattern: `^[a-zA-Z0-9+/]*={0,2}$`

**[QueryErrors](#API_ListGroupResources_ResponseSyntax)**

A list of `QueryError` objects. Each error contains an
`ErrorCode` and `Message`. Possible values for
ErrorCode are `CLOUDFORMATION_STACK_INACTIVE`, `CLOUDFORMATION_STACK_NOT_EXISTING`,
`CLOUDFORMATION_STACK_UNASSUMABLE_ROLE` and `RESOURCE_TYPE_NOT_SUPPORTED`.

Type: Array of [QueryError](api-queryerror.md) objects

**[ResourceIdentifiers](#API_ListGroupResources_ResponseSyntax)**

###### Important

**_Deprecated - don't use this parameter. Use the_**
**_`Resources` response field_**
**_instead._**

Type: Array of [ResourceIdentifier](api-resourceidentifier.md) objects

**[Resources](#API_ListGroupResources_ResponseSyntax)**

An array of resources from which you can determine each resource's identity, type, and
group membership status.

Type: Array of [ListGroupResourcesItem](api-listgroupresourcesitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The request includes one or more parameters that violate validation rules.

HTTP Status Code: 400

**ForbiddenException**

The caller isn't authorized to make the request. Check permissions.

HTTP Status Code: 403

**InternalServerErrorException**

An internal error occurred while processing the request. Try again later.

HTTP Status Code: 500

**MethodNotAllowedException**

The request uses an HTTP method that isn't allowed for the specified resource.

HTTP Status Code: 405

**NotFoundException**

One or more of the specified resources don't exist.

HTTP Status Code: 404

**TooManyRequestsException**

You've exceeded throttling limits by making too many requests in a period of
time.

HTTP Status Code: 429

**UnauthorizedException**

The request was rejected because it doesn't have valid credentials for the target
resource.

HTTP Status Code: 401

## Examples

### Example

This example illustrates one usage of ListGroupResources.

#### Sample Request

```

POST /list-group-resources HTTP/1.1
Host: resource-groups.us-west-2.amazonaws.com
Accept-Encoding: identity
User-Agent: aws-cli/2.2.40 Python/3.8.8 Windows/10 exe/AMD64 prompt/off command/resource-groups.list-group-resources
X-Amz-Date: 20220114T205755Z
X-Amz-Security-Token: <SECURITY-TOKEN>
Authorization: AWS4-HMAC-SHA256 Credential=<ACCESS-KEY>/20220113/us-west-2/resource-groups/aws4_request,SignedHeaders=host;x-amz-date;x-amz-security-token,Signature=<SIGV4-SIGNATURE>
Content-Length: 21

{
    "Group": "CRPGroup"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Fri, 14 Jan 2022 20:57:55 GMT
Content-Type: application/json
Content-Length: 654
x-amzn-RequestId: <VARIES>
x-amz-apigw-id: <VARIES>
X-Amzn-Trace-Id: Root=<VARIES>
Connection: keep-alive

{
    "Resources":[
        {
            "Identifier":{
                "ResourceArn":"arn:aws:ec2:us-west-2:123456789012:capacity-reservation/cr-0070b00d13EXAMPLE",
                "ResourceType":"AWS::EC2::CapacityReservation"
            }
        },
        {
            "Identifier":{
                "ResourceArn":"arn:aws:ec2:us-west-2:123456789012:capacity-reservation/cr-061abec820EXAMPLE",
                "ResourceType":"AWS::EC2::CapacityReservation"
            }
        }
    ],
    "QueryErrors":[]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for Python](../../../../services/goto/boto3/resource-groups-2017-11-27/listgroupresources.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/listgroupresources.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListGroupingStatuses

ListGroups

All content copied from https://docs.aws.amazon.com/.
