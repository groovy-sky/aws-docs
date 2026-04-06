# Use IAM policies to control workgroup access

To control access to workgroups, use resource-level IAM permissions or identity-based
IAM policies. Whenever you use IAM policies, make sure that you follow IAM best practices. For more information, see [Security best practices in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html) in the _IAM User Guide_.

###### Note

To access trusted identity propagation enabled workgroups, IAM Identity Center users must be
assigned to the `IdentityCenterApplicationArn` that is returned by the
response of the Athena [GetWorkGroup](https://docs.aws.amazon.com/athena/latest/APIReference/API_GetWorkGroup.html) API
action.

The following procedure is specific to Athena.

For IAM-specific information, see the links listed at the end of this section. For
information about example JSON workgroup policies, see [Example workgroup policies](https://docs.aws.amazon.com/athena/latest/ug/example-policies-workgroup.html).

###### To use the visual editor in the IAM console to create a workgroup policy

01. Sign in to the AWS Management Console and open the IAM console at [https://console.aws.amazon.com/iam/](https://console.aws.amazon.com/iam).

02. In the navigation pane on the left, choose **Policies**, and then
     choose **Create policy**.

03. On the **Visual editor** tab, choose **Choose a**
    **service**. Then choose Athena to add to the policy.

04. Choose **Select actions**, and then choose the actions to add to
     the policy. The visual editor shows the actions available in Athena. For more
     information, see [Actions,\
     resources, and condition keys for Amazon Athena](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonathena.html) in the
     _Service Authorization Reference_.

05. Choose **add actions** to type a specific action or use wildcards
     (\*) to specify multiple actions.

    By default, the policy that you are creating allows the actions that you choose.
     If you chose one or more actions that support resource-level permissions to the
     `workgroup` resource in Athena, then the editor lists the
     `workgroup` resource.

06. Choose **Resources** to specify the specific workgroups for your
     policy. For example JSON workgroup policies, see [Example workgroup policies](https://docs.aws.amazon.com/athena/latest/ug/example-policies-workgroup.html).

07. Specify the `workgroup` resource as follows:

    ```nohighlight

    arn:aws:athena:<region>:<user-account>:workgroup/<workgroup-name>
    ```

08. Choose **Review policy**, and then type a
     **Name** and a **Description** (optional) for
     the policy that you are creating. Review the policy summary to make sure that you
     granted the intended permissions.

09. Choose **Create policy** to save your new policy.

10. Attach this identity-based policy to a user, a group, or role.

For more information, see the following topics in the
_Service Authorization Reference_ and _IAM User Guide_:

- [Actions, resources, and condition\
keys for Amazon Athena](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonathena.html)

- [Creating policies with the visual editor](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html#access_policies_create-visual-editor)

- [Adding and\
removing IAM policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html)

- [Controlling access to resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_controlling.html#access_controlling-resources)

For example JSON workgroup policies, see [Example workgroup policies](https://docs.aws.amazon.com/athena/latest/ug/example-policies-workgroup.html).

For a complete list of Amazon Athena actions, see the API action names in the [Amazon Athena API Reference](../../../../reference/athena/latest/apireference.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Access to workgroups and tags

Example workgroup policies
