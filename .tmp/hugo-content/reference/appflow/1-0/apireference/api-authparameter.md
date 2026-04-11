# AuthParameter

Information about required authentication parameters.

## Contents

**connectorSuppliedValues**

Contains default values for this authentication parameter that are supplied by the
connector.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**description**

A description about the authentication parameter.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `[\s\w/!@#+=.-]*`

Required: No

**isRequired**

Indicates whether this authentication parameter is required.

Type: Boolean

Required: No

**isSensitiveField**

Indicates whether this authentication parameter is a sensitive field.

Type: Boolean

Required: No

**key**

The authentication key required to authenticate with the connector.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**label**

Label used for authentication parameter.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/authparameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/authparameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/authparameter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthenticationConfig

BasicAuthCredentials

All content copied from https://docs.aws.amazon.com/.
