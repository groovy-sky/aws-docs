# DescribeStacks

Retrieves a list that describes one or more specified stacks, if the stack names are provided. Otherwise, all stacks in the account are described.

## Request Syntax

```nohighlight

{
   "Names": [ "string" ],
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Names](#API_DescribeStacks_RequestSyntax)**

The names of the stacks to describe.

Type: Array of strings

Length Constraints: Minimum length of 1.

Required: No

**[NextToken](#API_DescribeStacks_RequestSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If this value is null, it retrieves the first page.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Stacks": [
      {
         "AccessEndpoints": [
            {
               "EndpointType": "string",
               "VpceId": "string"
            }
         ],
         "ApplicationSettings": {
            "Enabled": boolean,
            "S3BucketName": "string",
            "SettingsGroup": "string"
         },
         "Arn": "string",
         "CreatedTime": number,
         "Description": "string",
         "DisplayName": "string",
         "EmbedHostDomains": [ "string" ],
         "FeedbackURL": "string",
         "Name": "string",
         "RedirectURL": "string",
         "StackErrors": [
            {
               "ErrorCode": "string",
               "ErrorMessage": "string"
            }
         ],
         "StorageConnectors": [
            {
               "ConnectorType": "string",
               "Domains": [ "string" ],
               "DomainsRequireAdminConsent": [ "string" ],
               "ResourceIdentifier": "string"
            }
         ],
         "StreamingExperienceSettings": {
            "PreferredProtocol": "string"
         },
         "UserSettings": [
            {
               "Action": "string",
               "MaximumLength": number,
               "Permission": "string"
            }
         ]
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_DescribeStacks_ResponseSyntax)**

The pagination token to use to retrieve the next page of results for this operation. If there are no more pages, this value is null.

Type: String

Length Constraints: Minimum length of 1.

**[Stacks](#API_DescribeStacks_ResponseSyntax)**

Information about the stacks.

Type: Array of [Stack](api-stack.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/describestacks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/describestacks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/describestacks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/describestacks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/describestacks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/describestacks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/describestacks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/describestacks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/describestacks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/describestacks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeSoftwareAssociations

DescribeThemeForStack

All content copied from https://docs.aws.amazon.com/.
