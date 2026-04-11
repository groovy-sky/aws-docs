# SourceKeyword

A keyword that relates to the control data source.

For manual evidence, this keyword indicates if the manual evidence is a file or
text.

For automated evidence, this keyword identifies a specific CloudTrail event,
AWS Config rule, Security Hub CSPM control, or AWS API name.

To learn more about the supported keywords that you can use when mapping a control data
source, see the following pages in the _AWS Audit Manager User_
_Guide_:

- [AWS Config rules supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-config.md)

- [AWS Security Hub CSPM controls supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-ash.md)

- [API calls\
supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-api.md)

- [AWS CloudTrail event names supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-cloudtrail.md)

## Contents

**keywordInputType**

The input method for the keyword.

- `SELECT_FROM_LIST` is used when mapping a data source for automated
evidence.

- When `keywordInputType` is `SELECT_FROM_LIST`, a
keyword must be selected to collect automated evidence. For example, this
keyword can be a CloudTrail event name, a rule name for AWS Config, a Security Hub CSPM control, or the name of an AWS API call.

- `UPLOAD_FILE` and `INPUT_TEXT` are only used when mapping a
data source for manual evidence.

- When `keywordInputType` is `UPLOAD_FILE`, a file must
be uploaded as manual evidence.

- When `keywordInputType` is `INPUT_TEXT`, text must be
entered as manual evidence.

Type: String

Valid Values: `SELECT_FROM_LIST | UPLOAD_FILE | INPUT_TEXT`

Required: No

**keywordValue**

The value of the keyword that's used when mapping a control data source. For example,
this can be a CloudTrail event name, a rule name for AWS Config, a
Security Hub CSPM control, or the name of an AWS API call.

If you’re mapping a data source to a rule in AWS Config, the
`keywordValue` that you specify depends on the type of rule:

- For [managed rules](../../../../services/config/latest/developerguide/evaluate-config-use-managed-rules.md), you can use the rule identifier as the
`keywordValue`. You can find the rule identifier from the [list of AWS Config managed rules](../../../../services/config/latest/developerguide/managed-rules-by-aws-config.md). For some
rules, the rule identifier is different from the rule name. For example, the rule
name `restricted-ssh` has the following rule identifier:
`INCOMING_SSH_DISABLED`. Make sure to use the rule identifier, not the
rule name.

Keyword example for managed rules:

- Managed rule name: [s3-bucket-acl-prohibited](../../../../services/config/latest/developerguide/s3-bucket-acl-prohibited.md)

`keywordValue`: `S3_BUCKET_ACL_PROHIBITED`

- For [custom rules](../../../../services/config/latest/developerguide/evaluate-config-develop-rules.md), you form the `keywordValue`
by adding the `Custom_` prefix to the rule name. This prefix distinguishes
the custom rule from a managed rule.

Keyword example for custom rules:

- Custom rule name: my-custom-config-rule

`keywordValue`: `Custom_my-custom-config-rule`

- For [service-linked rules](../../../../services/config/latest/developerguide/service-linked-awsconfig-rules.md), you form the
`keywordValue` by adding the `Custom_` prefix to the rule
name. In addition, you remove the suffix ID that appears at the end of the rule name.

Keyword examples for service-linked rules:

- Service-linked rule name:
CustomRuleForAccount-conformance-pack-szsm1uv0w

`keywordValue`:
`Custom_CustomRuleForAccount-conformance-pack`

- Service-linked rule name:
OrgConfigRule-s3-bucket-versioning-enabled-dbgzf8ba

`keywordValue`:
`Custom_OrgConfigRule-s3-bucket-versioning-enabled`

###### Important

The `keywordValue` is case sensitive. If you enter a value incorrectly, Audit Manager might not recognize the data source mapping. As a result, you might not
successfully collect evidence from that data source as intended.

Keep in mind the following requirements, depending on the data source type that
you're using.

1. For AWS Config:

- For managed rules, make sure that the `keywordValue` is the rule identifier in
`ALL_CAPS_WITH_UNDERSCORES`. For example,
`CLOUDWATCH_LOG_GROUP_ENCRYPTED`. For accuracy, we recommend
that you reference the list of [supported AWS Config managed rules](../../../../services/audit-manager/latest/userguide/control-data-sources-config.md).

- For custom rules, make sure that the `keywordValue` has the `Custom_`
prefix followed by the custom rule name. The format of the custom rule name
itself may vary. For accuracy, we recommend that you visit the [AWS Config console](https://console.aws.amazon.com/config) to
verify your custom rule name.

2. For Security Hub CSPM: The format varies for Security Hub CSPM control names.
    For accuracy, we recommend that you reference the list of [supported\
    AWS Security Hub CSPM controls](../../../../services/audit-manager/latest/userguide/control-data-sources-ash.md).

3. For AWS API calls: Make sure that the `keywordValue`
    is written as `serviceprefix_ActionName`. For example,
    `iam_ListGroups`. For accuracy, we recommend that you reference the
    list of [supported\
    API calls](../../../../services/audit-manager/latest/userguide/control-data-sources-api.md).

4. For CloudTrail: Make sure that the `keywordValue` is written
    as `serviceprefix_ActionName`. For example,
    `cloudtrail_StartLogging`. For accuracy, we recommend that you
    review the AWS service prefix and action names in the [Service Authorization Reference](../../../../services/service-authorization/latest/reference/reference-policies-actions-resources-contextkeys.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z_0-9-\s().:\/]+$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/sourcekeyword.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/sourcekeyword.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/sourcekeyword.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Settings

UpdateAssessmentFrameworkControlSet

All content copied from https://docs.aws.amazon.com/.
