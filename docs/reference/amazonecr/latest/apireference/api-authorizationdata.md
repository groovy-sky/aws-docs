# AuthorizationData

An object representing authorization data for an Amazon ECR registry.

## Contents

**authorizationToken**

A base64-encoded string that contains authorization data for the specified Amazon ECR
registry. When the string is decoded, it is presented in the format
`user:password` for private registry authentication using `docker
                login`.

Type: String

Pattern: `^\S+$`

Required: No

**expiresAt**

The Unix time in seconds and milliseconds when the authorization token expires.
Authorization tokens are valid for 12 hours.

Type: Timestamp

Required: No

**proxyEndpoint**

The registry URL to use for this authorization token in a `docker login`
command. The Amazon ECR registry URL format is
`https://aws_account_id.dkr.ecr.region.amazonaws.com`. For example,
`https://012345678910.dkr.ecr.us-east-1.amazonaws.com`..

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/authorizationdata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/authorizationdata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/authorizationdata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Attribute

AwsEcrContainerImageDetails

All content copied from https://docs.aws.amazon.com/.
