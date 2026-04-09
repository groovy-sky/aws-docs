# ListVpcConnectors

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Returns a list of AWS App Runner VPC connectors in your AWS account.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_ListVpcConnectors_RequestSyntax)**

The maximum number of results to include in each response (result page). It's used for a paginated request.

If you don't specify `MaxResults`, the request retrieves all available results in a single response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_ListVpcConnectors_RequestSyntax)**

A token from a previous result page. It's used for a paginated request. The request retrieves the next result page. All other parameter values must be
identical to the ones that are specified in the initial request.

If you don't specify `NextToken`, the request retrieves the first result page.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `.*`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "VpcConnectors": [
      {
         "CreatedAt": number,
         "DeletedAt": number,
         "SecurityGroups": [ "string" ],
         "Status": "string",
         "Subnets": [ "string" ],
         "VpcConnectorArn": "string",
         "VpcConnectorName": "string",
         "VpcConnectorRevision": number
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListVpcConnectors_ResponseSyntax)**

The token that you can pass in a subsequent request to get the next result page. It's returned in a paginated request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `.*`

**[VpcConnectors](#API_ListVpcConnectors_ResponseSyntax)**

A list of information records for VPC connectors. In a paginated request, the request returns up to `MaxResults` records for each
call.

Type: Array of [VpcConnector](api-vpcconnector.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

## Examples

### Paginated listing of App Runner VPC connectors

This example illustrates how to list all App Runner VPC connectors in your AWS account. Up to five VPC connectors are listed
in each response.

In this example, the response includes one result and there aren't additional ones, so no `NextToken` is returned.

#### Sample Request

```json

$ aws apprunner list-vpc-connectors --cli-input-json "`cat`"
{
  "MaxResults": 5
}
```

#### Sample Response

```json

{
  "VpcConnectors": [
    {
      "VpcConnectorArn": "arn:aws:apprunner:us-east-1:123456789012:vpcconnector/my-vpc-connector/1/3f2eb10e2c494674952026f646844e3d",
      "VpcConnectorName": "my-vpc-connector",
      "VpcConnectorRevision": 1,
      "Subnets": ["subnet-123", "subnet-456"],
      "SecurityGroups": ["sg-123", "sg-456"],
      "Status": "ACTIVE",
      "CreatedAt": "2021-08-18T23:36:45.374Z"
    }
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/listvpcconnectors.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/listvpcconnectors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForResource

ListVpcIngressConnections

All content copied from https://docs.aws.amazon.com/.
