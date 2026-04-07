# FailedConfiguration

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

A configuration ID paired with an error message.

## Contents

**configurationId**

The unique identifier of the configuration the failed to delete.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: No

**errorMessage**

A descriptive message indicating why the associated configuration failed to delete.

Type: String

Required: No

**errorStatusCode**

The integer error code associated with the error message.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/discovery-2015-11-01/FailedConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/discovery-2015-11-01/FailedConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/discovery-2015-11-01/FailedConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExportPreferences

Filter
