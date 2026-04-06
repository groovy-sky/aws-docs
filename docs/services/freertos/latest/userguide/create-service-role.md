# Create an OTA Update service role

The OTA Update service assumes this role to create and manage OTA update jobs on your behalf.

###### To create an OTA service role

01. Sign in to the [https://console.aws.amazon.com/iam/](https://console.aws.amazon.com/iam).

02. From the navigation pane, choose **Roles**.

03. Choose **Create role**.

04. Under **Select type of trusted entity**, choose **AWS**
    **Service**.

05. Choose **IoT** from the list of AWS services.

06. Under **Select your use case**, choose
     **IoT**.

07. Choose **Next: Permissions**.

08. Choose **Next: Tags**.

09. Choose **Next: Review**.

10. Enter a role name and description, and then choose **Create**
    **role**.

For more information about IAM roles, see [IAM \
Roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html).

###### Important

To address the confused deputy security issue, you must follow instructions in the
[AWS IoT Core](https://docs.aws.amazon.com/iot/latest/developerguide/cross-service-confused-deputy-prevention.html) guide.

###### To add OTA update permissions to your OTA service role

1. In the search box on the IAM console page, enter the name of your role, and
    then choose it from the list.

2. Choose **Attach policies**.

3. In the **Search** box, enter "AmazonFreeRTOSOTAUpdate",
    select **AmazonFreeRTOSOTAUpdate** from the list of filtered
    policies, and then choose **Attach policy** to attach the
    policy to your service role.

###### To add the required IAM permissions to your OTA service role

1. In the search box on the IAM console page, enter the name of your role, and
    then choose it from the list.

2. Choose **Add inline policy**.

3. Choose the **JSON** tab.

4. Copy and paste the following policy document into the text box:
JSON

```json

{
       "Version":"2012-10-17",
       "Statement": [
           {
               "Effect": "Allow",
               "Action": [
                   "iam:GetRole",
                   "iam:PassRole"
               ],
               "Resource": "arn:aws:iam::111122223333:role/your_role_name"
           }
       ]
}

```

Make sure that you replace `your_account_id`
    with your AWS account ID, and
    `your_role_name` with the name of the OTA
    service role.

5. Choose **Review policy**.

6. Enter a name for the policy, and then choose **Create**
**policy**.

###### Note

The following procedure isn't required if your Amazon S3 bucket name begins with "afr-ota".
If it does, the AWS managed policy `AmazonFreeRTOSOTAUpdate` already includes
the required permissions.

###### **To add the required Amazon S3 permissions to your OTA service role**

1. In the search box on the IAM console page, enter the name of your role, and
    then choose it from the list.

2. Choose **Add inline policy**.

3. Choose the **JSON** tab.

4. Copy and paste the following policy document into the box.
JSON

```json

{
       "Version":"2012-10-17",
       "Statement": [
           {
               "Effect": "Allow",
               "Action": [
                   "s3:GetObjectVersion",
                   "s3:GetObject",
                   "s3:PutObject"
               ],
               "Resource": [
                   "arn:aws:s3:::example-bucket/*"
               ]
           }
       ]
}

```

This policy grants your OTA service role permission to read Amazon S3 objects. Make
    sure that you replace `example-bucket` with the
    name of your bucket.

5. Choose **Review policy**.

6. Enter a name for the policy, and then choose **Create**
**policy**.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create an Amazon S3 bucket to store your update

Create an OTA user policy
