# ObjectLambdaAccessPointAlias

The alias of an Object Lambda Access Point. For more information, see [How to use a\
bucket-style alias for your S3 bucket Object Lambda Access Point](../userguide/olap-use.md#ol-access-points-alias).

## Contents

**Status**

The status of the Object Lambda Access Point alias. If the status is `PROVISIONING`, the Object Lambda Access Point
is provisioning the alias and the alias is not ready for use yet. If the status is
`READY`, the Object Lambda Access Point alias is successfully provisioned and ready for
use.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 16.

Valid Values: `PROVISIONING | READY`

Required: No

**Value**

The alias value of the Object Lambda Access Point.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `^[0-9a-z\\-]{3,63}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/objectlambdaaccesspointalias.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/objectlambdaaccesspointalias.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/objectlambdaaccesspointalias.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectLambdaAccessPoint

ObjectLambdaConfiguration

All content copied from https://docs.aws.amazon.com/.
