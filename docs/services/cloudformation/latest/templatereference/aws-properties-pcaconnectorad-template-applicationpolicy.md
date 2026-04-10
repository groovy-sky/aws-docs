This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::Template ApplicationPolicy

Application policies describe what the certificate can be used for.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PolicyObjectIdentifier" : String,
  "PolicyType" : String
}

```

### YAML

```yaml

  PolicyObjectIdentifier: String
  PolicyType: String

```

## Properties

`PolicyObjectIdentifier`

The object identifier (OID) of an application policy.

_Required_: No

_Type_: String

_Pattern_: `^([0-2])\.([0-9]|([0-3][0-9]))(\.([0-9]+)){0,126}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyType`

The type of application policy

_Required_: No

_Type_: String

_Allowed values_: `ALL_APPLICATION_POLICIES | ANY_PURPOSE | ATTESTATION_IDENTITY_KEY_CERTIFICATE | CERTIFICATE_REQUEST_AGENT | CLIENT_AUTHENTICATION | CODE_SIGNING | CTL_USAGE | DIGITAL_RIGHTS | DIRECTORY_SERVICE_EMAIL_REPLICATION | DISALLOWED_LIST | DNS_SERVER_TRUST | DOCUMENT_ENCRYPTION | DOCUMENT_SIGNING | DYNAMIC_CODE_GENERATOR | EARLY_LAUNCH_ANTIMALWARE_DRIVER | EMBEDDED_WINDOWS_SYSTEM_COMPONENT_VERIFICATION | ENCLAVE | ENCRYPTING_FILE_SYSTEM | ENDORSEMENT_KEY_CERTIFICATE | FILE_RECOVERY | HAL_EXTENSION | IP_SECURITY_END_SYSTEM | IP_SECURITY_IKE_INTERMEDIATE | IP_SECURITY_TUNNEL_TERMINATION | IP_SECURITY_USER | ISOLATED_USER_MODE | KDC_AUTHENTICATION | KERNEL_MODE_CODE_SIGNING | KEY_PACK_LICENSES | KEY_RECOVERY | KEY_RECOVERY_AGENT | LICENSE_SERVER_VERIFICATION | LIFETIME_SIGNING | MICROSOFT_PUBLISHER | MICROSOFT_TIME_STAMPING | MICROSOFT_TRUST_LIST_SIGNING | OCSP_SIGNING | OEM_WINDOWS_SYSTEM_COMPONENT_VERIFICATION | PLATFORM_CERTIFICATE | PREVIEW_BUILD_SIGNING | PRIVATE_KEY_ARCHIVAL | PROTECTED_PROCESS_LIGHT_VERIFICATION | PROTECTED_PROCESS_VERIFICATION | QUALIFIED_SUBORDINATION | REVOKED_LIST_SIGNER | ROOT_PROGRAM_AUTO_UPDATE_CA_REVOCATION | ROOT_PROGRAM_AUTO_UPDATE_END_REVOCATION | ROOT_PROGRAM_NO_OSCP_FAILOVER_TO_CRL | ROOT_LIST_SIGNER | SECURE_EMAIL | SERVER_AUTHENTICATION | SMART_CARD_LOGIN | SPC_ENCRYPTED_DIGEST_RETRY_COUNT | SPC_RELAXED_PE_MARKER_CHECK | TIME_STAMPING | WINDOWS_HARDWARE_DRIVER_ATTESTED_VERIFICATION | WINDOWS_HARDWARE_DRIVER_EXTENDED_VERIFICATION | WINDOWS_HARDWARE_DRIVER_VERIFICATION | WINDOWS_HELLO_RECOVERY_KEY_ENCRYPTION | WINDOWS_KITS_COMPONENT | WINDOWS_RT_VERIFICATION | WINDOWS_SOFTWARE_EXTENSION_VERIFICATION | WINDOWS_STORE | WINDOWS_SYSTEM_COMPONENT_VERIFICATION | WINDOWS_TCB_COMPONENT | WINDOWS_THIRD_PARTY_APPLICATION_COMPONENT | WINDOWS_UPDATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationPolicies

CertificateValidity

All content copied from https://docs.aws.amazon.com/.
