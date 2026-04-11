# ImageSigningStatus

The signing status for an image. Each status corresponds to a signing profile.

## Contents

**failureCode**

The failure code, which is only present if `status`
is `FAILED`.

Type: String

Required: No

**failureReason**

A description of why signing the image failed. This field is only
present if `status` is `FAILED`.

Type: String

Required: No

**signingProfileArn**

The ARN of the AWS Signer signing profile used to sign the image.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `^arn:aws(-[a-z]+)*:signer:[a-z0-9-]+:[0-9]{12}:\/signing-profiles\/[a-zA-Z0-9_]{2,}$`

Required: No

**status**

The image's signing status. Possible values are:

- `IN_PROGRESS` \- Signing is currently in progress.

- `COMPLETE` \- The signature was successfully generated.

- `FAILED` \- Signing failed. See
`failureCode` and `failureReason` for details.

Type: String

Valid Values: `IN_PROGRESS | COMPLETE | FAILED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagesigningstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagesigningstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagesigningstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageScanStatus

ImageTagMutabilityExclusionFilter

All content copied from https://docs.aws.amazon.com/.
