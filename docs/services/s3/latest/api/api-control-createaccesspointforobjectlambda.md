# CreateAccessPointForObjectLambda

###### Note

This operation is not supported by directory buckets.

Creates an Object Lambda Access Point. For more information, see [Transforming objects with\
Object Lambda Access Points](../userguide/transforming-objects.md) in the _Amazon S3 User Guide_.

The following actions are related to
`CreateAccessPointForObjectLambda`:

- [DeleteAccessPointForObjectLambda](api-control-deleteaccesspointforobjectlambda.md)

- [GetAccessPointForObjectLambda](api-control-getaccesspointforobjectlambda.md)

- [ListAccessPointsForObjectLambda](api-control-listaccesspointsforobjectlambda.md)

## Request Syntax

```nohighlight

PUT /v20180820/accesspointforobjectlambda/name HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<CreateAccessPointForObjectLambdaRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <Configuration>
      <AllowedFeatures>
         <AllowedFeature>string</AllowedFeature>
      </AllowedFeatures>
      <CloudWatchMetricsEnabled>boolean</CloudWatchMetricsEnabled>
      <SupportingAccessPoint>string</SupportingAccessPoint>
      <TransformationConfigurations>
         <TransformationConfiguration>
            <Actions>
               <Action>string</Action>
            </Actions>
            <ContentTransformation>
               <AwsLambda>
                  <FunctionArn>string</FunctionArn>
                  <FunctionPayload>string</FunctionPayload>
               </AwsLambda>
            </ContentTransformation>
         </TransformationConfiguration>
      </TransformationConfigurations>
   </Configuration>
</CreateAccessPointForObjectLambdaRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_CreateAccessPointForObjectLambda_RequestSyntax)**

The name you want to assign to this Object Lambda Access Point.

Length Constraints: Minimum length of 3. Maximum length of 45.

Pattern: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

Required: Yes

**[x-amz-account-id](#API_control_CreateAccessPointForObjectLambda_RequestSyntax)**

The AWS account ID for owner of the specified Object Lambda Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[CreateAccessPointForObjectLambdaRequest](#API_control_CreateAccessPointForObjectLambda_RequestSyntax)**

Root level tag for the CreateAccessPointForObjectLambdaRequest parameters.

Required: Yes

**[Configuration](#API_control_CreateAccessPointForObjectLambda_RequestSyntax)**

Object Lambda Access Point configuration as a JSON document.

Type: [ObjectLambdaConfiguration](api-control-objectlambdaconfiguration.md) data type

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<CreateAccessPointForObjectLambdaResult>
   <ObjectLambdaAccessPointArn>string</ObjectLambdaAccessPointArn>
   <Alias>
      <Status>string</Status>
      <Value>string</Value>
   </Alias>
</CreateAccessPointForObjectLambdaResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[CreateAccessPointForObjectLambdaResult](#API_control_CreateAccessPointForObjectLambda_ResponseSyntax)**

Root level tag for the CreateAccessPointForObjectLambdaResult parameters.

Required: Yes

**[Alias](#API_control_CreateAccessPointForObjectLambda_ResponseSyntax)**

The alias of the Object Lambda Access Point.

Type: [ObjectLambdaAccessPointAlias](api-control-objectlambdaaccesspointalias.md) data type

**[ObjectLambdaAccessPointArn](#API_control_CreateAccessPointForObjectLambda_ResponseSyntax)**

Specifies the ARN for the Object Lambda Access Point.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:[^:]+:s3-object-lambda:[^:]*:\d{12}:accesspoint/.*`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/createaccesspointforobjectlambda.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/createaccesspointforobjectlambda.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAccessPoint

CreateBucket

All content copied from https://docs.aws.amazon.com/.
