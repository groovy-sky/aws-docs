# AwsEcrContainerImageDetails

The image details of the Amazon ECR container image.

## Contents

**architecture**

The architecture of the Amazon ECR container image.

Type: String

Required: No

**author**

The image author of the Amazon ECR container image.

Type: String

Required: No

**imageHash**

The image hash of the Amazon ECR container image.

Type: String

Required: No

**imageTags**

The image tags attached to the Amazon ECR container image.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 300.

Required: No

**inUseCount**

The number of Amazon ECS or Amazon EKS clusters currently running the
image.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lastInUseAt**

The most recent date and time a cluster was running the image.

Type: Timestamp

Required: No

**platform**

The platform of the Amazon ECR container image.

Type: String

Required: No

**pushedAt**

The date and time the Amazon ECR container image was pushed.

Type: Timestamp

Required: No

**registry**

The registry the Amazon ECR container image belongs to.

Type: String

Pattern: `[0-9]{12}`

Required: No

**repositoryName**

The name of the repository the Amazon ECR container image resides in.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/awsecrcontainerimagedetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/awsecrcontainerimagedetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/awsecrcontainerimagedetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationData

CvssScore

All content copied from https://docs.aws.amazon.com/.
