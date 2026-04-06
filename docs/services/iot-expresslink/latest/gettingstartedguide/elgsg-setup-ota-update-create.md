# Create a firmware update job in AWS IoT

01. Open the [AWS IoT console](https://console.aws.amazon.com/iot/home). Choose
     **Manage** then choose **Jobs**. Choose
     **Create job**, **Create FreeRTOS OTA Update Job**,
     then choose **Next**.

02. Enter a job name which is unique within your AWS account. Enter an optional
     description. Choose **Next**.

03. From the **Devices to update** dropdown, choose the thing name
     associated with your ExpressLink module when it was registered with the account.
     Choose **MQTT** as the protocol that will be used to perform the
     transfer, and unselect **HTTP** if it is selected.

04. Choose **Use my custom signed file**; this will display a form
     to be filled in. Use the details from [Prerequisites](elgsg-setup-ota-update-prereqs.md) to fill in the form.

05. In the **signature** field, enter the base64 encoded signature
     for the image. From the **Original hashing algorithm** drop down,
     select the hashing algorithm provided by the manufacturer. From the
     **Original encryption algorithm** drop down, select the encryption
     algorithm provided by the manufacturer. In the **Path name of code signing**
    **certificate on device**, enter the path name, if any is provided by the
     manufacturer. (If the path name is not provided, then you can just enter 'NA'.)

06. Choose **Upload a new file**, then **Choose file**
     and upload the image you received from the manufacturer. Choose **Create S3**
    **bucket** for the new uploaded image and proceed with creating a new bucket.
     If needed, you can also choose an existing bucket in your account by selecting
     **Browse S3** option.

07. Under **Path Name of file on device** you can enter 'NA' if the
     image is not targeted as an executable file within a filesystem.

08. From the **File type** drop down select a value (default is 0)
     to signify this is an ExpressLink firmware update as opposed to a host firmware
     update.

09. From the **role** dropdown under the **IAM role**
     section, select the OTA update role you created above. Choose **Next**.

10. Choose **Create Job**. If the job creation was successful, it
     should list the job name and state as 'in progress'.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Prerequisites

Monitor and apply a new firmware update for AWS IoT ExpressLink
