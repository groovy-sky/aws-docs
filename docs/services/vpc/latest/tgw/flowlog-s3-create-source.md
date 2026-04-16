---
title: "Create the AWS Transit Gateway Flow Logs source account role for Amazon S3"
---

# Create the AWS Transit Gateway Flow Logs source account role for Amazon S3

From the source account, create the source role in the AWS Identity and Access Management console.

###### To create the source account role

01. Sign in to the AWS Management Console and open the IAM console at [https://console.aws.amazon.com/iam/](https://console.aws.amazon.com/iam).

02. In the navigation pane, choose **Policies**.

03. Choose **Create policy**.

04. On the Create policy page, do the following:

1. Choose **JSON**.

2. Replace the contents of this window with the permissions
    policy at the start of this section.

3. Choose **Next: Tags** and **Next:**
**Review**.

4. Enter a name for your policy and an optional description, and
    then choose **Create policy**.

05. In the navigation pane, choose **Roles**.

06. Choose **Create role**.

07. For the **Trusted entity type**, choose **Custom**
    **trust policy**. For **Custom trust**
    **policy**, replace `"Principal": {},` with the
     following, which specifies the log delivery service. Choose
     **Next**.

    ```json

    "Principal": {
       "Service": "delivery.logs.amazonaws.com"
    },
    ```

08. On the **Add permissions** page, select the checkbox for the
     policy that you created earlier in this procedure, and then choose
     **Next**.

09. Enter a name for your role and optionally provide a description.

10. Choose **Create role**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Flow Logs

Create a Flow Log that publishes
to Amazon S3

All content copied from https://docs.aws.amazon.com/.
