# DeregistrationPolicy

The deregistration policy for the data that's stored in AWS Audit Manager. You can
use this attribute to determine how your data is handled when you [deregister Audit Manager](api-deregisteraccount.md).

By default, Audit Manager retains evidence data for two years from the time of its
creation. Other Audit Manager resources (including assessments, custom controls, and
custom frameworks) remain in Audit Manager indefinitely, and are available if you
[re-register Audit Manager](api-registeraccount.md) in the future. For more
information about data retention, see [Data\
Protection](../../../../services/audit-manager/latest/userguide/data-protection.md) in the _AWS Audit Manager User Guide_.

###### Important

If you choose to delete all data, this action permanently deletes all evidence data
in your account within seven days. It also deletes all of the Audit Manager
resources that you created, including assessments, custom controls, and custom
frameworks. Your data will not be available if you re-register Audit Manager in the
future.

## Contents

**deleteResources**

Specifies which Audit Manager data will be deleted when you deregister Audit Manager.

- If you set the value to `ALL`, all of your data is deleted within seven
days of deregistration.

- If you set the value to `DEFAULT`, none of your data is deleted at the
time of deregistration. However, keep in mind that the Audit Manager data
retention policy still applies. As a result, any evidence data will be deleted two
years after its creation date. Your other Audit Manager resources will continue
to exist indefinitely.

Type: String

Valid Values: `ALL | DEFAULT`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/deregistrationpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/deregistrationpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/deregistrationpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DelegationMetadata

Evidence

All content copied from https://docs.aws.amazon.com/.
