# UpsolverDestinationProperties

The properties that are applied when Upsolver is used as a destination.

## Contents

**bucketName**

The Upsolver Amazon S3 bucket name in which Amazon AppFlow places the
transferred data.

Type: String

Length Constraints: Minimum length of 16. Maximum length of 63.

Pattern: `^(upsolver-appflow)\S*`

Required: Yes

**s3OutputFormatConfig**

The configuration that determines how data is formatted when Upsolver is used as the flow
destination.

Type: [UpsolverS3OutputFormatConfig](api-upsolvers3outputformatconfig.md) object

Required: Yes

**bucketPrefix**

The object key for the destination Upsolver Amazon S3 bucket in which Amazon AppFlow places the files.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/upsolverdestinationproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/upsolverdestinationproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/upsolverdestinationproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TriggerProperties

UpsolverMetadata

All content copied from https://docs.aws.amazon.com/.
