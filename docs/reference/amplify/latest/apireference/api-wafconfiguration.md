# WafConfiguration

Describes the Firewall configuration for a hosted Amplify application.
Firewall support enables you to protect your web applications with a direct integration
with AWS WAF. For more information about using AWS WAF protections
for an Amplify application, see [Firewall support for hosted\
sites](../../../../services/amplify/latest/userguide/waf-integration.md) in the _Amplify User Guide_.

## Contents

**statusReason**

The reason for the current status of the Firewall configuration.

Type: String

Length Constraints: Maximum length of 1000.

Required: No

**wafStatus**

The status of the process to associate or disassociate a web ACL to an Amplify app.

Type: String

Valid Values: `ASSOCIATING | ASSOCIATION_FAILED | ASSOCIATION_SUCCESS | DISASSOCIATING | DISASSOCIATION_FAILED`

Required: No

**webAclArn**

The Amazon Resource Name (ARN) for the web ACL associated with an Amplify app.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 512.

Pattern: `^arn:aws:wafv2:.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/wafconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/wafconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/wafconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SubDomainSetting

Webhook
