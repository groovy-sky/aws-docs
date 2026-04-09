# RegistryScanningRule

The details of a scanning rule for a private registry.

## Contents

**repositoryFilters**

The repository filters associated with the scanning configuration for a private
registry.

Type: Array of [ScanningRepositoryFilter](api-scanningrepositoryfilter.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: Yes

**scanFrequency**

The frequency that scans are performed at for a private registry. When the
`ENHANCED` scan type is specified, the supported scan frequencies are
`CONTINUOUS_SCAN` and `SCAN_ON_PUSH`. When the
`BASIC` scan type is specified, the `SCAN_ON_PUSH` scan
frequency is supported. If scan on push is not specified, then the `MANUAL`
scan frequency is set by default.

Type: String

Valid Values: `SCAN_ON_PUSH | CONTINUOUS_SCAN | MANUAL`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/registryscanningrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/registryscanningrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/registryscanningrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegistryScanningConfiguration

Remediation

All content copied from https://docs.aws.amazon.com/.
