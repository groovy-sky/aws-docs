# PrivateConnectionProvisioningState

Specifies the private connection provisioning state.

## Contents

**failureCause**

Specifies the private connection provisioning failure cause.

Type: String

Valid Values: `CONNECTOR_AUTHENTICATION | CONNECTOR_SERVER | INTERNAL_SERVER | ACCESS_DENIED | VALIDATION`

Required: No

**failureMessage**

Specifies the private connection provisioning failure reason.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\s\w/!@#+=.-]*`

Required: No

**status**

Specifies the private connection provisioning status.

Type: String

Valid Values: `FAILED | PENDING | CREATED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/privateconnectionprovisioningstate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/privateconnectionprovisioningstate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/privateconnectionprovisioningstate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PrefixConfig

Range
