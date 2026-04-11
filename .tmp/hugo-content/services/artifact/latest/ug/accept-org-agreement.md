# Accepting agreements for your organization in AWS Artifact

If you're the owner of the management account of an AWS Organizations organization, then you can
accept an agreement with AWS on behalf of all AWS accounts in your
organization.

###### Important

Before you accept an agreement, we recommend that you consult with your legal,
privacy, and compliance team.

AWS Organizations has two available feature sets: _consolidated billing_
_features_ and _all features_. To use AWS Artifact for your
organization, the organization that you belong to must be enabled for [all features](../../../organizations/latest/userguide/orgs-getting-started-concepts.md#feature-set). If your organization is configured only for consolidated
billing, see [Enabling all features in your organization](../../../organizations/latest/userguide/orgs-manage-org-support-all-features.md) in the
_AWS Organizations User Guide_.

To accept or terminate organization agreements, you must be signed in to the
management account with the correct AWS Artifact permissions. Users of member accounts that have
`organizations:DescribeOrganization` permissions can view the
organization agreements that are accepted on their behalf.

For more information, see [Managing accounts in\
an organization with AWS Organizations](../../../organizations/latest/userguide/orgs-manage-accounts.md) in the
_AWS Organizations User Guide_.

###### Required permissions

To accept an agreement, the owner of the management account must have the required
[permissions](example-iam-policies.md#example-policy-organizational-agreements).

For more information, see [Identity and access management in AWS Artifact](security-iam.md).

###### To accept an agreement for an organization

01. Open the AWS Artifact console at [https://console.aws.amazon.com/artifact/](https://console.aws.amazon.com/artifact).

02. On the AWS Artifact dashboard, choose **Agreements**.

03. Choose the **Organization agreements** tab.

04. Open the AWS Artifact console at [https://console.aws.amazon.com/artifact/](https://console.aws.amazon.com/artifact).

05. In the navigation pane, choose **Agreements**.

06. On the **Agreements** page, do either of the following:
    - To accept an agreement only for your account, choose the **Account agreements**
       tab.

    - To accept an agreement on behalf of your organization, choose the **Organization**
      **agreements** tab.
07. Select an agreement, and then choose **Download agreement**.

    The **Accept NDA to download report** dialog box appears.

08. Before you can download the agreement that you selected, you must first accept the terms of the AWS Artifact
     Nondisclosure Agreement (AWS Artifact NDA).
    1. In the **Accept NDA to download report** dialog box, review the AWS Artifact NDA.

    2. (Optional) To print a copy of the AWS Artifact NDA (or to save it as a PDF), choose **Print**
       **NDA**.

    3. Select **I have read and agree to all the**
       **terms of the NDA.**

    4. To accept the AWS Artifact NDA and to download a PDF of the agreement that you selected, choose **Accept NDA and**
       **download**.
09. In a PDF viewer, review the agreement PDF that you downloaded.

10. In the AWS Artifact console, with the agreement selected, choose **Accept agreement**.

11. In the **Accept agreement** dialog box, do the following:
    1. Review the agreement.

    2. Select **I agree to all of these terms and conditions**.

    3. Choose **Accept agreement**.
12. Choose **Accept** to accept the agreement for all existing
     and future accounts in your organization.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Terminating account
agreements

Terminating organization
agreements

All content copied from https://docs.aws.amazon.com/.
