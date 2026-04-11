# ConnectorRuntimeSetting

Contains information about the connector runtime settings that are required for flow
execution.

## Contents

**connectorSuppliedValueOptions**

Contains default values for the connector runtime setting that are supplied by the
connector.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**dataType**

Data type of the connector runtime setting.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**description**

A description about the connector runtime setting.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `[\s\w/!@#+=.-]*`

Required: No

**isRequired**

Indicates whether this connector runtime setting is required.

Type: Boolean

Required: No

**key**

Contains value information about the connector runtime setting.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**label**

A label used for connector runtime setting.

Type: String

Length Constraints: Maximum length of 128.

Pattern: `.*`

Required: No

**scope**

Indicates the scope of the connector runtime setting.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectorruntimesetting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectorruntimesetting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectorruntimesetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorProvisioningConfig

CustomAuthConfig

All content copied from https://docs.aws.amazon.com/.
