---
title: "DescribeImageReplicationStatus"
---

# DescribeImageReplicationStatus

Returns the replication status for a specified image.

## Request Syntax

```nohighlight

{
   "imageId": {
      "imageDigest": "string",
      "imageTag": "string"
   },
   "registryId": "string",
   "repositoryName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[imageId](#API_DescribeImageReplicationStatus_RequestSyntax)**

An object with identifying information for an image in an Amazon ECR repository.

Type: [ImageIdentifier](api-imageidentifier.md) object

Required: Yes

**[registryId](#API_DescribeImageReplicationStatus_RequestSyntax)**

The AWS account ID associated with the registry. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_DescribeImageReplicationStatus_RequestSyntax)**

The name of the repository that the image is in.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "imageId": {
      "imageDigest": "string",
      "imageTag": "string"
   },
   "replicationStatuses": [
      {
         "failureCode": "string",
         "region": "string",
         "registryId": "string",
         "status": "string"
      }
   ],
   "repositoryName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[imageId](#API_DescribeImageReplicationStatus_ResponseSyntax)**

An object with identifying information for an image in an Amazon ECR repository.

Type: [ImageIdentifier](api-imageidentifier.md) object

**[replicationStatuses](#API_DescribeImageReplicationStatus_ResponseSyntax)**

The replication status details for the images in the specified repository.

Type: Array of [ImageReplicationStatus](api-imagereplicationstatus.md) objects

**[repositoryName](#API_DescribeImageReplicationStatus_ResponseSyntax)**

The repository name associated with the request.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ImageNotFoundException**

The image requested does not exist in the specified repository.

HTTP Status Code: 400

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**RepositoryNotFoundException**

The specified repository could not be found. Check the spelling of the specified
repository and ensure that you are performing operations on the correct registry.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ecr-2015-09-21/DescribeImageReplicationStatus)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/DescribeImageReplicationStatus)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeregisterPullTimeUpdateExclusion

DescribeImages

All content copied from https://docs.aws.amazon.com/.
