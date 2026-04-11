# CreateRepositoryCreationTemplate

Creates a repository creation template. This template is used to define the settings
for repositories created by Amazon ECR on your behalf. For example, repositories created
through pull through cache actions. For more information, see [Private\
repository creation templates](../../../../services/amazonecr/latest/userguide/repository-creation-templates.md) in the
_Amazon Elastic Container Registry User Guide_.

## Request Syntax

```nohighlight

{
   "appliedFor": [ "string" ],
   "customRoleArn": "string",
   "description": "string",
   "encryptionConfiguration": {
      "encryptionType": "string",
      "kmsKey": "string"
   },
   "imageTagMutability": "string",
   "imageTagMutabilityExclusionFilters": [
      {
         "filter": "string",
         "filterType": "string"
      }
   ],
   "lifecyclePolicy": "string",
   "prefix": "string",
   "repositoryPolicy": "string",
   "resourceTags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[appliedFor](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

A list of enumerable strings representing the Amazon ECR repository creation scenarios that
this template will apply towards. The supported scenarios are
`PULL_THROUGH_CACHE`, `REPLICATION`, and `CREATE_ON_PUSH`

Type: Array of strings

Valid Values: `REPLICATION | PULL_THROUGH_CACHE | CREATE_ON_PUSH`

Required: Yes

**[customRoleArn](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The ARN of the role to be assumed by Amazon ECR. This role must be in the same account as
the registry that you are configuring. Amazon ECR will assume your supplied role when the
customRoleArn is specified. When this field isn't specified, Amazon ECR will use the
service-linked role for the repository creation template.

Type: String

Length Constraints: Maximum length of 2048.

Required: No

**[description](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

A description for the repository creation template.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**[encryptionConfiguration](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The encryption configuration to use for repositories created using the
template.

Type: [EncryptionConfigurationForRepositoryCreationTemplate](api-encryptionconfigurationforrepositorycreationtemplate.md) object

Required: No

**[imageTagMutability](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The tag mutability setting for the repository. If this parameter is omitted, the
default setting of `MUTABLE` will be used which will allow image tags to be
overwritten. If `IMMUTABLE` is specified, all image tags within the
repository will be immutable which will prevent them from being overwritten.

Type: String

Valid Values: `MUTABLE | IMMUTABLE | IMMUTABLE_WITH_EXCLUSION | MUTABLE_WITH_EXCLUSION`

Required: No

**[imageTagMutabilityExclusionFilters](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

A list of filters that specify which image tags should be excluded from the repository
creation template's image tag mutability setting.

Type: Array of [ImageTagMutabilityExclusionFilter](api-imagetagmutabilityexclusionfilter.md) objects

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Required: No

**[lifecyclePolicy](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The lifecycle policy to use for repositories created using the template.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 30720.

Required: No

**[prefix](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The repository namespace prefix to associate with the template. All repositories
created using this namespace prefix will have the settings defined in this template
applied. For example, a prefix of `prod` would apply to all repositories
beginning with `prod/`. Similarly, a prefix of `prod/team` would
apply to all repositories beginning with `prod/team/`.

To apply a template to all repositories in your registry that don't have an associated
creation template, you can use `ROOT` as the prefix.

###### Important

There is always an assumed `/` applied to the end of the prefix. If you
specify `ecr-public` as the prefix, Amazon ECR treats that as
`ecr-public/`. When using a pull through cache rule, the repository
prefix you specify during rule creation is what you should specify as your
repository creation template prefix as well.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: Yes

**[repositoryPolicy](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The repository policy to apply to repositories created using the template. A
repository policy is a permissions policy associated with a repository to control access
permissions.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10240.

Required: No

**[resourceTags](#API_CreateRepositoryCreationTemplate_RequestSyntax)**

The metadata to apply to the repository to help you categorize and organize. Each tag
consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Syntax

```nohighlight

{
   "registryId": "string",
   "repositoryCreationTemplate": {
      "appliedFor": [ "string" ],
      "createdAt": number,
      "customRoleArn": "string",
      "description": "string",
      "encryptionConfiguration": {
         "encryptionType": "string",
         "kmsKey": "string"
      },
      "imageTagMutability": "string",
      "imageTagMutabilityExclusionFilters": [
         {
            "filter": "string",
            "filterType": "string"
         }
      ],
      "lifecyclePolicy": "string",
      "prefix": "string",
      "repositoryPolicy": "string",
      "resourceTags": [
         {
            "Key": "string",
            "Value": "string"
         }
      ],
      "updatedAt": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[registryId](#API_CreateRepositoryCreationTemplate_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryCreationTemplate](#API_CreateRepositoryCreationTemplate_ResponseSyntax)**

The details of the repository creation template associated with the request.

Type: [RepositoryCreationTemplate](api-repositorycreationtemplate.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**LimitExceededException**

The operation did not succeed because it would have exceeded a service limit for your
account. For more information, see [Amazon ECR service quotas](../../../../services/amazonecr/latest/userguide/service-quotas.md) in
the Amazon Elastic Container Registry User Guide.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**TemplateAlreadyExistsException**

The repository creation template already exists. Specify a unique prefix and try
again.

HTTP Status Code: 400

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## Examples

In the following example or examples, the Authorization header contents
( `AUTHPARAMS`) must be replaced with an AWS Signature Version 4
signature. For more information about creating these signatures, see [Signature\
Version 4 Signing Process](../../../../general/latest/gr/signature-version-4.md) in the _AWS General_
_Reference_.

You only need to learn how to sign HTTP requests if you intend to manually
create them. When you use the [AWS Command Line Interface (AWS CLI)](http://aws.amazon.com/cli) or one of the [AWS SDKs](http://aws.amazon.com/tools) to make requests to AWS, these tools automatically sign the
requests for you with the access key that you specify when you configure the tools.
When you use these tools, you don't need to learn how to sign requests
yourself.

### Example

This example creates a repository creation template in the default registry
for an account.

#### Sample Request

```

POST / HTTP/1.1
Host: ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 1134
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.CreateRepositoryCreationTemplate
X-Amz-Date: 20231216T195356Z
User-Agent: aws-cli/1.11.22 Python/2.7.12 Darwin/16.3.0 botocore/1.4.79
Content-Type: application/x-amz-json-1.1
Authorization: AUTHPARAMS

{
 "prefix": "eng/test",
 "description": "Repos for testing images",
 "encryptionConfiguration":
  {"encryptionType": "AES256"},
 "resourceTags":
  [
    {"Key": "environment",
     "Value": "test"}
  ],
 "imageTagMutability": "MUTABLE",
 "repositoryPolicy": "{\r\n  \"Version\": \"2012-10-17\",\r\n  \"Statement\": [\r\n    {\r\n      \"Sid\": \"LambdaECRPullPolicy\",\r\n      \"Effect\": \"Allow\",\r\n      \"Principal\": {\r\n        \"Service\": \"lambda.amazonaws.com\"\r\n      },\r\n      \"Action\": \"ecr:BatchGetImage\"\r\n    }\r\n  ]\r\n}",
 "lifecyclePolicy": "{\r\n    \"rules\": [\r\n        {\r\n            \"rulePriority\": 1,\r\n            \"description\": \"Expire images older than 14 days\",\r\n            \"selection\": {\r\n                \"tagStatus\": \"untagged\",\r\n                \"countType\": \"sinceImagePushed\",\r\n                \"countUnit\": \"days\",\r\n                \"countNumber\": 14\r\n            },\r\n            \"action\": {\r\n                \"type\": \"expire\"\r\n            }\r\n        }\r\n    ]\r\n}",
 "appliedFor":
  ["REPLICATION", "PULL_THROUGH_CACHE", "CREATE_ON_PUSH"]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Sat, 16 Dec 2023 19:53:56 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 963
Connection: keep-alive
x-amzn-RequestId: 60dc1ea1-c3c9-11e6-aa04-25c3a5fb1b54

{
 "registryId": "012345678901",
 "repositoryCreationTemplate":
  {
   "appliedFor":
    ["REPLICATION", "PULL_THROUGH_CACHE", "CREATE_ON_PUSH"],
   "description": "Repos for testing images",
   "encryptionConfiguration":
    {
        "encryptionType": "AES256"
    },
   "imageTagMutability": "MUTABLE",
   "lifecyclePolicy": "{\r\n    \"rules\": [\r\n        {\r\n            \"rulePriority\": 1,\r\n            \"description\": \"Expire images older than 14 days\",\r\n            \"selection\": {\r\n                \"tagStatus\": \"untagged\",\r\n                \"countType\": \"sinceImagePushed\",\r\n                \"countUnit\": \"days\",\r\n                \"countNumber\": 14\r\n            },\r\n            \"action\": {\r\n                \"type\": \"expire\"\r\n            }\r\n        }\r\n    ]\r\n}",
   "prefix": "eng/test",
   "repositoryPolicy": "{\n  \"Version\" : \"2012-10-17\",\n  \"Statement\" : [ {\n    \"Sid\" : \"LambdaECRPullPolicy\",\n    \"Effect\" : \"Allow\",\n    \"Principal\" : {\n      \"Service\" : \"lambda.amazonaws.com\"\n    },\n    \"Action\" : \"ecr:BatchGetImage\"\n  } ]\n}",
   "resourceTags":
    [
      {"Key": "environment",
       "Value": "test"}
    ],
   "createdAt": "2023-12-16T17:29:02-07:00",
   "updatedAt": "2023-12-16T17:29:02-07:00"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/createrepositorycreationtemplate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/createrepositorycreationtemplate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateRepository

DeleteLifecyclePolicy

All content copied from https://docs.aws.amazon.com/.
