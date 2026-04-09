# GetMultiRegionAccessPointPolicyStatus

###### Note

This operation is not supported by directory buckets.

Indicates whether the specified Multi-Region Access Point has an access control policy that allows public
access.

This action will always be routed to the US West (Oregon) Region. For more information
about the restrictions around working with Multi-Region Access Points, see [Multi-Region Access Point\
restrictions and limitations](../userguide/multiregionaccesspointrestrictions.md) in the _Amazon S3 User Guide_.

The following actions are related to
`GetMultiRegionAccessPointPolicyStatus`:

- [GetMultiRegionAccessPointPolicy](api-control-getmultiregionaccesspointpolicy.md)

- [PutMultiRegionAccessPointPolicy](api-control-putmultiregionaccesspointpolicy.md)

## Request Syntax

```nohighlight

GET /v20180820/mrap/instances/name+/policystatus HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetMultiRegionAccessPointPolicyStatus_RequestSyntax)**

Specifies the Multi-Region Access Point. The name of the Multi-Region Access Point is different from the alias. For more
information about the distinction between the name and the alias of an Multi-Region Access Point, see [Rules for naming Amazon S3 Multi-Region Access Points](../userguide/creatingmultiregionaccesspoints.md#multi-region-access-point-naming) in the
_Amazon S3 User Guide_.

Length Constraints: Maximum length of 50.

Pattern: `^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$`

Required: Yes

**[x-amz-account-id](#API_control_GetMultiRegionAccessPointPolicyStatus_RequestSyntax)**

The AWS account ID for the owner of the Multi-Region Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetMultiRegionAccessPointPolicyStatusResult>
   <Established>
      <IsPublic>boolean</IsPublic>
   </Established>
</GetMultiRegionAccessPointPolicyStatusResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetMultiRegionAccessPointPolicyStatusResult](#API_control_GetMultiRegionAccessPointPolicyStatus_ResponseSyntax)**

Root level tag for the GetMultiRegionAccessPointPolicyStatusResult parameters.

Required: Yes

**[Established](#API_control_GetMultiRegionAccessPointPolicyStatus_ResponseSyntax)**

Indicates whether this access point policy is public. For more information about how Amazon S3
evaluates policies to determine whether they are public, see [The Meaning of "Public"](../dev/access-control-block-public-access.md#access-control-block-public-access-policy-status) in the _Amazon S3 User Guide_.

Type: [PolicyStatus](api-control-policystatus.md) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getmultiregionaccesspointpolicystatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetMultiRegionAccessPointPolicy

GetMultiRegionAccessPointRoutes

All content copied from https://docs.aws.amazon.com/.
