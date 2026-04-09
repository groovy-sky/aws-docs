# GetAccessPointForObjectLambda

###### Note

This operation is not supported by directory buckets.

Returns configuration information about the specified Object Lambda Access Point

The following actions are related to `GetAccessPointForObjectLambda`:

- [CreateAccessPointForObjectLambda](api-control-createaccesspointforobjectlambda.md)

- [DeleteAccessPointForObjectLambda](api-control-deleteaccesspointforobjectlambda.md)

- [ListAccessPointsForObjectLambda](api-control-listaccesspointsforobjectlambda.md)

## Request Syntax

```nohighlight

GET /v20180820/accesspointforobjectlambda/name HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetAccessPointForObjectLambda_RequestSyntax)**

The name of the Object Lambda Access Point.

Length Constraints: Minimum length of 3. Maximum length of 45.

Pattern: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

Required: Yes

**[x-amz-account-id](#API_control_GetAccessPointForObjectLambda_RequestSyntax)**

The account ID for the account that owns the specified Object Lambda Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetAccessPointForObjectLambdaResult>
   <Name>string</Name>
   <PublicAccessBlockConfiguration>
      <BlockPublicAcls>boolean</BlockPublicAcls>
      <BlockPublicPolicy>boolean</BlockPublicPolicy>
      <IgnorePublicAcls>boolean</IgnorePublicAcls>
      <RestrictPublicBuckets>boolean</RestrictPublicBuckets>
   </PublicAccessBlockConfiguration>
   <CreationDate>timestamp</CreationDate>
   <Alias>
      <Status>string</Status>
      <Value>string</Value>
   </Alias>
</GetAccessPointForObjectLambdaResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetAccessPointForObjectLambdaResult](#API_control_GetAccessPointForObjectLambda_ResponseSyntax)**

Root level tag for the GetAccessPointForObjectLambdaResult parameters.

Required: Yes

**[Alias](#API_control_GetAccessPointForObjectLambda_ResponseSyntax)**

The alias of the Object Lambda Access Point.

Type: [ObjectLambdaAccessPointAlias](api-control-objectlambdaaccesspointalias.md) data type

**[CreationDate](#API_control_GetAccessPointForObjectLambda_ResponseSyntax)**

The date and time when the specified Object Lambda Access Point was created.

Type: Timestamp

**[Name](#API_control_GetAccessPointForObjectLambda_ResponseSyntax)**

The name of the Object Lambda Access Point.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 45.

Pattern: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

**[PublicAccessBlockConfiguration](#API_control_GetAccessPointForObjectLambda_ResponseSyntax)**

Configuration to block all public access. This setting is turned on and can not be
edited.

Type: [PublicAccessBlockConfiguration](api-control-publicaccessblockconfiguration.md) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getaccesspointforobjectlambda.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getaccesspointforobjectlambda.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAccessPointConfigurationForObjectLambda

GetAccessPointPolicy

All content copied from https://docs.aws.amazon.com/.
