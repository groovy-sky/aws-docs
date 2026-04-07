AWS Audit Manager will no longer be open to new customers starting
April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers
can continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](audit-manager-availability-change.md).

# Enabling evidence finder

You can enable the evidence finder feature in Audit Manager to search for evidence in your
AWS account. If you're a delegated administrator for Audit Manager, you can search for evidence
for all member accounts in your organization.

Follow these steps to learn how to enable evidence finder. Pay close attention to the
prerequisites, as you'll need specific permissions to create and manage an event data
store in CloudTrail Lake for this functionality.

## Prerequisites

### Required permissions to enable evidence finder

To enable evidence finder, you need permissions to create and manage an event
data store in CloudTrail Lake. To use the feature, you need permissions to perform
CloudTrail Lake queries. For an example permission policy that you can use, see [Example 3 (Permissions to enable evidence finder)](https://docs.aws.amazon.com/audit-manager/latest/userguide/security_iam_id-based-policy-examples.html#full-administrator-access-enable-evidence-finder).

If you need help with permissions, contact your AWS administrator. If you’re
an AWS administrator, you can copy the required permission statement and
[attach it to an IAM policy](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html#add-policies-console).

## Procedure

### Requesting to enable evidence finder

You can complete this task using the Audit Manager console, the AWS Command Line Interface (AWS CLI), or
the Audit Manager API.

###### Note

You must enable evidence finder in each AWS Region where you want to
search for evidence.

Audit Manager console

###### To request to enable evidence finder on the Audit Manager console

1. Open the AWS Audit Manager console at [https://console.aws.amazon.com/auditmanager/home](https://console.aws.amazon.com/auditmanager/home).

2. From the **Evidence finder** settings
    tab, go to the **Evidence finder**
    section.

3. Choose **Required permission policy**,
    then **View CloudTrail Lake permissions** to view
    the required evidence finder permissions. If you don't
    already have these permissions, you can copy this policy
    statement and [attach it to an IAM\
    policy](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html#add-policies-console).

4. Choose **Enable**.

5. In the pop-up window, choose **Request to**
**enable**.

AWS CLI

###### To request to enable evidence finder in the AWS CLI

Run the [update-settings](https://docs.aws.amazon.com/cli/latest/reference/auditmanager/update-settings.html) command with the
`--evidence-finder-enabled` parameter.

```

aws auditmanager update-settings --evidence-finder-enabled
```

Audit Manager API

###### To request to enable evidence finder using the API

Call the [UpdateSettings](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_UpdateSettings.html) operation and use the [evidenceFinderEnabled](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_UpdateSettings.html#auditmanager-UpdateSettings-request-evidenceFinderEnabled) parameter.

For more information, choose the previous links to read more in
the _Audit Manager API Reference_. This
includes information about how to use this operation and parameter
in one of the language-specific AWS SDKs.

## Next steps

After you've requested to enable evidence finder, you can check the status of your
request. For instructions, see [Confirming the status of evidence finder](https://docs.aws.amazon.com/audit-manager/latest/userguide/confirm-status-of-evidence-finder.html).

## Additional resources

- [Evidence finder](evidence-finder.md)

- [Troubleshooting evidence finder issues](https://docs.aws.amazon.com/audit-manager/latest/userguide/evidence-finder-issues.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configuring your Audit Manager notifications

Confirming the status of evidence finder
