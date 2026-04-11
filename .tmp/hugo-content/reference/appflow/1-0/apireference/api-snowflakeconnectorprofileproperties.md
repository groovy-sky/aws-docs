# SnowflakeConnectorProfileProperties

The connector-specific profile properties required when using Snowflake.

## Contents

**bucketName**

The name of the Amazon S3 bucket associated with Snowflake.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: Yes

**stage**

The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snowflake account. This is written in the following format: <
Database>< Schema><Stage Name>.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**warehouse**

The name of the Snowflake warehouse.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `[\s\w/!@#+=.-]*`

Required: Yes

**accountName**

The name of the account.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**bucketPrefix**

The bucket path that refers to the Amazon S3 bucket associated with Snowflake.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**privateLinkServiceName**

The Snowflake Private Link service name to be used for private data transfers.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `^$|com.amazonaws.vpce.[\w/!:@#.\-]+`

Required: No

**region**

The AWS Region of the Snowflake account.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/snowflakeconnectorprofileproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/snowflakeconnectorprofileproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/snowflakeconnectorprofileproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeConnectorProfileCredentials

SnowflakeDestinationProperties

All content copied from https://docs.aws.amazon.com/.
