# Getting set up with Billing

Use this section to get started with the AWS Billing and Cost Management console. Prerequisites include signing up for AWS, setting up IAM users, and reviewing your AWS bills.

###### Topics

- [Learn more about Billing features](#billing-gs-features)

- [What do I do next?](#what-next)

- [Setting up your tax information](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-account-payment.html)

- [Customizing your Billing preferences](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-pref.html)

- [Customizing your AWS payment preferences](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-payment-method.html)

- [Setting up your India billing](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-account-payment-aispl.html)

- [Finding the seller of record](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/finding-the-seller-of-record.html)

- [Reviewing your monthly billing best practices](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/monthly-billing-checklist.html)

If you're new to AWS, create an AWS account. For more information, see [Getting Started with AWS](https://aws.amazon.com/getting-started).

### Sign up for an AWS account

If you do not have an AWS account, complete the following steps to create one.

###### To sign up for an AWS account

1. Open [https://portal.aws.amazon.com/billing/signup](https://portal.aws.amazon.com/billing/signup).

2. Follow the online instructions.

Part of the sign-up procedure involves receiving a phone call or text message and entering
    a verification code on the phone keypad.

When you sign up for an AWS account, an _AWS account root user_ is created. The root user has access to all AWS services
    and resources in the account. As a security best practice, assign administrative access to a user, and use only the root user to perform [tasks that require root user access](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#root-user-tasks).

AWS sends you a confirmation email after the sign-up process is
complete. At any time, you can view your current account activity and manage your account by
going to [https://aws.amazon.com/](https://aws.amazon.com/) and choosing **My**
**Account**.

### Create a user with administrative access

After you sign up for an AWS account, secure your AWS account root user, enable AWS IAM Identity Center, and create an administrative user so that you
don't use the root user for everyday tasks.

###### Secure your AWS account root user

1. Sign in to the [AWS Management Console](https://console.aws.amazon.com/) as the account owner by choosing **Root user** and entering your AWS account email address. On the next page, enter your password.

For help signing in by using root user, see [Signing in as the root user](https://docs.aws.amazon.com/signin/latest/userguide/console-sign-in-tutorials.html#introduction-to-root-user-sign-in-tutorial) in the _AWS Sign-In User Guide_.

2. Turn on multi-factor authentication (MFA) for your root user.

For instructions, see [Enable a virtual MFA device for your AWS account root user (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/enable-virt-mfa-for-root.html) in the _IAM User Guide_.

###### Create a user with administrative access

1. Enable IAM Identity Center.

For instructions, see [Enabling\
    AWS IAM Identity Center](https://docs.aws.amazon.com/singlesignon/latest/userguide/get-set-up-for-idc.html) in the
    _AWS IAM Identity Center User Guide_.

2. In IAM Identity Center, grant administrative access to a user.

For a tutorial about using the IAM Identity Center directory as your identity source, see [Configure user access with the default IAM Identity Center directory](https://docs.aws.amazon.com/singlesignon/latest/userguide/quick-start-default-idc.html) in the
    _AWS IAM Identity Center User Guide_.

###### Sign in as the user with administrative access

- To sign in with your IAM Identity Center user, use the sign-in URL that was sent to your email address when you created the IAM Identity Center user.

For help signing in using an IAM Identity Center user, see [Signing in to the AWS access portal](https://docs.aws.amazon.com/signin/latest/userguide/iam-id-center-sign-in-tutorial.html) in the _AWS Sign-In User Guide_.

###### Assign access to additional users

1. In IAM Identity Center, create a permission set that follows the best practice of applying least-privilege permissions.

For instructions, see [Create a permission set](https://docs.aws.amazon.com/singlesignon/latest/userguide/get-started-create-a-permission-set.html) in the _AWS IAM Identity Center User Guide_.

2. Assign users to a group, and then assign single sign-on access to the group.

For instructions, see [Add groups](https://docs.aws.amazon.com/singlesignon/latest/userguide/addgroups.html) in the _AWS IAM Identity Center User Guide_.

### Activating IAM access to the AWS Billing and Cost Management console

By default, IAM roles within an AWS account can't access the Billing and Cost Management console. This is
true even if the IAM user or role has IAM policies that grant access to specific
Billing features. The root user can allow IAM users and roles access to Billing and Cost Management console by
using the **Activate IAM access** setting.

###### To provide access to the Billing and Cost Management console

1. Sign in to the **Account** page in the Billing and Cost Management console at [https://console.aws.amazon.com/billing/home?#/account](https://console.aws.amazon.com/billing/home?).

2. Under **IAM user and role access to Billing**
**information**, choose **Edit**.

3. Select **Activate IAM access**.

4. Choose **Update**.

For more information about this feature, see [Activating access to the Billing and Cost Management console](control-access-billing.md#ControllingAccessWebsite-Activate).

Use features in the Billing and Cost Management console to view your current AWS charges and AWS
usage.

###### To open the Billing and Cost Management console and view your usage and charges

1. Sign into the AWS Management Console and open the Billing and Cost Management console at
    [https://console.aws.amazon.com/costmanagement/](https://console.aws.amazon.com/costmanagement).

2. Choose **Bills** to see details about your current charges.

3. Choose **Payments** to see your historical payment
    transactions.

4. Choose **AWS Cost and Usage Reports** to see reports that break down your
    costs.

For more information about setting up and using AWS Cost and Usage Reports, see the [AWS Cost and Usage Reports User Guide](https://docs.aws.amazon.com/cur/latest/userguide/what-is-cur.html).

AWS Billing closes the billing period at midnight on the last day of each month and calculates your bill. Most bills are ready for you to download by the seventh accounting
day of the month.

###### To download or print your bill

1. Sign into the AWS Management Console and open the Billing and Cost Management console at
    [https://console.aws.amazon.com/costmanagement/](https://console.aws.amazon.com/costmanagement).

2. On the navigation pane, choose **Bills**.

3. For **Date**, choose the month of the bill you want to work with.

4. Choose **Download CSV** to download a comma-separated variable file or choose
    **Print**.

### Adding or updating alternate contacts

Alternate contacts allows AWS to contact another person about issues with your
account, even if you're unavailable. The alternate contact doesn't have to be a specific
person. You could instead add an email distribution list if you have a team that manages
billing, operations and security related issues.

**Examples for alternate contacts**

AWS will reach out to each contact type in the following scenarios:

- **Billing** – When your monthly invoice is available, or your
payment method needs to be updated. If you enabled **Receive PDF**
**Invoice By Email** in your **Billing**
**preferences**, your alternate billing contact also receives the
PDF invoices. Notifications can be from AWS service teams.

- **Operations** – When your service is, or will be, temporarily
unavailable in one of more AWS Regions. Your contacts will also receive
any notification related to operations. Notifications can be from
AWS service teams

- **Security** – When you have notifications from the AWS
Security, AWS Trust and Safety, or AWS service teams. These
notifications might include security issues or potential abusive or
fraudulent activities on your AWS account. Notifications can be from
AWS service teams concerning security related topics associated with your
AWS account usage. Don't include sensitive information in the subject line
or full name fields since this might be used in email communications to
you.

For more information about managing your alternate account contacts, see [Alternate account contacts](https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-alternate.html) in the
_AWS Account Management_ Reference Guide.

## Learn more about Billing features

Understand the features available to you in the Billing and Cost Management console.

- **AWS Free Tier**: [Trying services using AWS Free Tier (before July 15, 2025)](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-free-tier.html)

- **Payments**: [Managing Your Payments](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-payments.html)

- **Viewing your bills**: [Understanding your bill](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/getting-viewing-bill.html)

- **AWS Cost Categories**: [Organizing costs using AWS Cost Categories](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-cost-categories.html)

- **Cost Allocation Tags**: [Organizing and tracking costs using AWS cost allocation tags](cost-alloc-tags.md)

- **AWS Purchase Orders**: [Managing your purchase orders](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-purchaseorders.html)

- **AWS Cost and Usage Reports**: [Using AWS Cost and Usage Reports](https://docs.aws.amazon.com/cur/latest/userguide/what-is-cur.html)

- **Using AWS CloudTrail**: [Logging Billing and Cost Management API calls with AWS CloudTrail](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/logging-using-cloudtrail.html)

- **Consolidated billing**: [Consolidating billing for AWS Organizations](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/consolidated-billing.html)

## What do I do next?

Now that you can view and pay your AWS bill, you're ready to use the features available to you. The rest of this guide helps you navigate your journey using the console.

### Optimize your spending using AWS Cost Management features

Use the AWS Cost Management features to budget and forecast costs so you can optimize your AWS spends and reduce your overall AWS bill. Combine and use the Billing and Cost Management console resources to manage your payments, while using AWS Cost Management features to optimize your future costs.

For more information about AWS Cost Management features, see the [AWS Cost Management User Guide](https://docs.aws.amazon.com/cost-management/latest/userguide/what-is-costmanagement.html).

### Using the Billing and Cost Management API

Use the [AWS Billing and Cost Management API Reference](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/Welcome.html) to programmatically use some AWS Cost Management features.

### Learn more

You can find more information about Billing features including presentations, virtual workshops, and blog posts on the marketing page [Cloud Financial Management with AWS](https://aws.amazon.com/aws-cost-management).

You can find virtual workshops by choosing the **Services** dropdown list
and selecting your feature.

### Get help

If you have questions about any Billing features, there are many resources available for
you. To learn more, see [Getting help with your bills and payments](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-get-answers.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

What is AWS Billing and Cost Management?

Setting up your tax information
