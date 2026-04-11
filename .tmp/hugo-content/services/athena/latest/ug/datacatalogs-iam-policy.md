# Control access to data catalogs with IAM policies

To control access to data catalogs, use resource-level IAM permissions or identity-based
IAM policies.

The following procedure is specific to Athena.

For IAM-specific information, see the links listed at the end of this section. For
information about example JSON data catalog policies, see [Data Catalog example policies](datacatalogs-example-policies.md).

###### To use the visual editor in the IAM console to create a data catalog policy

01. Sign in to the AWS Management Console and open the IAM console at [https://console.aws.amazon.com/iam/](https://console.aws.amazon.com/iam).

02. In the navigation pane on the left, choose **Policies**, and then
     choose **Create policy**.

03. On the **Visual editor** tab, choose **Choose a**
    **service**. Then choose Athena to add to the policy.

04. Choose **Select actions**, and then choose the actions to add to
     the policy. The visual editor shows the actions available in Athena. For more
     information, see [Actions,\
     resources, and condition keys for Amazon Athena](../../../service-authorization/latest/reference/list-amazonathena.md) in the
     _Service Authorization Reference_.

05. Choose **add actions** to type a specific action or use wildcards
     (\*) to specify multiple actions.

    By default, the policy that you are creating allows the actions that you choose.
     If you chose one or more actions that support resource-level permissions to the
     `datacatalog` resource in Athena, then the editor lists the
     `datacatalog` resource.

06. Choose **Resources** to specify the specific data catalogs for
     your policy. For example JSON data catalog policies, see [Data Catalog example policies](datacatalogs-example-policies.md).

07. Specify the `datacatalog` resource as follows:

    ```nohighlight

    arn:aws:athena:<region>:<user-account>:datacatalog/<datacatalog-name>
    ```

08. Choose **Review policy**, and then type a
     **Name** and a **Description** (optional) for
     the policy that you are creating. Review the policy summary to make sure that you
     granted the intended permissions.

09. Choose **Create policy** to save your new policy.

10. Attach this identity-based policy to a user, a group, or role and specify the
     `datacatalog` resources they can access.

For more information, see the following topics in the
_Service Authorization Reference_ and the
_IAM User Guide_:

- [Actions, resources, and\
condition keys for Amazon Athena](../../../service-authorization/latest/reference/list-amazonathena.md)

- [Creating policies with the visual editor](../../../iam/latest/userguide/access-policies-create.md#access_policies_create-visual-editor)

- [Adding and\
removing IAM policies](../../../iam/latest/userguide/access-policies-manage-attach-detach.md)

- [Controlling access to resources](../../../iam/latest/userguide/access-controlling.md#access_controlling-resources)

For example JSON data catalog policies, see [Data Catalog example policies](datacatalogs-example-policies.md).

For information about AWS Glue permissions and AWS Glue crawler permissions, see [Setting up IAM\
permissions for AWS Glue](../../../glue/latest/dg/getting-started-access.md) and [Crawler prerequisites](../../../glue/latest/dg/crawler-prereqs.md) in the _AWS Glue Developer Guide_.

For a complete list of Amazon Athena actions, see the API action names in the [Amazon Athena API Reference](../../../../reference/athena/latest/apireference.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Register a catalog from another
account

Data Catalog example policies

All content copied from https://docs.aws.amazon.com/.
