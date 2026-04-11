# Migrating custom headers out of the build specification and amplify.yml

Previously, custom HTTP headers were specified for an app either by editing the build specification
in the Amplify console or by downloading and updating the
`amplify.yml` file and saving it in the project 's root directory. It is
strongly recommended that you migrate your custom headers out of the build specification and the
`amplify.yml` file.

Specify your custom headers in the **Custom headers** section of the
Amplify console or by downloading and editing the
`customHttp.yml` file.

###### To migrate custom headers stored in the Amplify console

01. Sign in to the AWS Management Console and open the [Amplify\
     console](https://console.aws.amazon.com/amplify).

02. Choose the app to perform the custom header migration on.

03. In the navigation pane, choose **Hosting**, **Build**
    **settings**. In the **App build specification** section, you
     can review your app's buildspec.

04. Choose **Download** to save a copy of your current buildspec. You can
     reference this copy later if you need to recover any settings.

05. When the download is complete, choose **Edit**.

06. Take note of the custom header information in the file, as you will use it later in
     step 9. In the **Edit** window, delete any custom headers from the file
     and choose **Save**.

07. In the navigation pane, choose **Hosting**, **Custom**
    **headers**.

08. On the **Custom headers** page, choose
     **Edit**.

09. In the **Edit custom headers** window, enter the information for your custom headers
     that you deleted in step 6.

10. Choose **Save**.

11. Redeploy any branch that you want the new custom headers to be applied to.

###### To migrate custom headers from amplify.yml to customHttp.yml

01. Navigate to the `amplify.yml` file currently deployed in your app's
     root directory.

02. Open `amplify.yml` in the code editor of your choice.

03. Take note of the custom header information in the file, as you will use it later in
     step 8. Delete the custom headers in the file. Save and close the file.

04. Sign in to the AWS Management Console and open the [Amplify\
     console](https://console.aws.amazon.com/amplify).

05. Choose the app to set custom headers for.

06. In the navigation pane, choose **Hosting**, **Custom**
    **headers**.

07. On the **Custom headers** page, choose
     **Download**.

08. Open the downloaded `customHttp.yml` file in the code editor of
     your choice and enter the information for your custom headers that you deleted from
     `amplify.yml` in step 3.

09. Save the edited `customHttp.yml` file in your project's root
     directory. If you are working with a monorepo, save the file in the root of your
     repo.

10. Redeploy the app to apply the new custom headers.
    - For a CI/CD app, perform a new build from your Git repository that includes the
       new `customHttp.yml` file.

    - For a manual deploy app, deploy the app again in the Amplify console and include
       the new `customHttp.yml` file with artifacts that you
       upload.

###### Note

Custom headers set in the `customHttp.yml` file and deployed in the
app's root directory override the custom headers defined in the **Custom**
**headers** section of the Amplify console.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting custom headers

Monorepo custom headers

All content copied from https://docs.aws.amazon.com/.
