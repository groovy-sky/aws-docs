# APISchema

Contains details about the OpenAPI schema for a custom plugin. For more information,
see [custom plugin OpenAPI schemas](../qbusiness-ug/custom-plugin.md#plugins-api-schema). You can either include
the schema directly in the payload field or you can upload it to an S3 bucket and
specify the S3 bucket location in the `s3` field.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**payload**

The JSON or YAML-formatted payload defining the OpenAPI schema for a custom plugin.

Type: String

Required: No

**s3**

Contains details about the S3 object containing the OpenAPI schema for a custom
plugin. The schema could be in either JSON or YAML format.

Type: [S3](api-s3.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/apischema.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/apischema.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/apischema.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionSummary

Application

All content copied from https://docs.aws.amazon.com/.
