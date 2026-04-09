# ImageDetail

An object that describes an image returned by a [DescribeImages](api-describeimages.md)
operation.

## Contents

**artifactMediaType**

The artifact media type of the image.

Type: String

Required: No

**imageDigest**

The `sha256` digest of the image manifest.

Type: String

Required: No

**imageManifestMediaType**

The media type of the image manifest.

Type: String

Required: No

**imagePushedAt**

The date and time, expressed in standard JavaScript date format, at which the current
image was pushed to the repository.

Type: Timestamp

Required: No

**imageScanFindingsSummary**

A summary of the last completed image scan.

Type: [ImageScanFindingsSummary](api-imagescanfindingssummary.md) object

Required: No

**imageScanStatus**

The current state of the scan.

Type: [ImageScanStatus](api-imagescanstatus.md) object

Required: No

**imageSizeInBytes**

The size, in bytes, of the image in the repository.

If the image is a manifest list, this will be the max size of all manifests in the
list.

###### Note

Starting with Docker version 1.9, the Docker client compresses image layers before
pushing them to a V2 Docker registry. The output of the `docker images`
command shows the uncompressed image size. Therefore, Docker might return a larger
image than the image shown in the AWS Management Console.

Type: Long

Required: No

**imageStatus**

The current status of the image.

Type: String

Valid Values: `ACTIVE | ARCHIVED | ACTIVATING`

Required: No

**imageTags**

The list of tags associated with this image.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 300.

Required: No

**lastActivatedAt**

The date and time, expressed in standard JavaScript date format, when the image was last restored from Amazon ECR archive to Amazon ECR standard.

Type: Timestamp

Required: No

**lastArchivedAt**

The date and time, expressed in standard JavaScript date format, when the image was last transitioned to Amazon ECR archive.

Type: Timestamp

Required: No

**lastRecordedPullTime**

The date and time, expressed in standard JavaScript date format, when Amazon ECR recorded
the last image pull.

###### Note

Amazon ECR refreshes the last image pull timestamp at least once every 24 hours. For
example, if you pull an image once a day then the `lastRecordedPullTime`
timestamp will indicate the exact time that the image was last pulled. However, if
you pull an image once an hour, because Amazon ECR refreshes the
`lastRecordedPullTime` timestamp at least once every 24 hours, the
result may not be the exact time that the image was last pulled.

Type: Timestamp

Required: No

**registryId**

The AWS account ID associated with the registry to which this image belongs.

Type: String

Pattern: `[0-9]{12}`

Required: No

**repositoryName**

The name of the repository to which this image belongs.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: No

**subjectManifestDigest**

The digest of the subject manifest for images that are referrers.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagedetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagedetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagedetail.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Image

ImageFailure

All content copied from https://docs.aws.amazon.com/.
