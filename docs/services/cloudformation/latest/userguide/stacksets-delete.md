# Delete CloudFormation StackSets

To delete a StackSet, you must first delete all stacks in the StackSet. For information about
how to delete all stacks, see [Delete stacks from\
StackSets](stackinstances-delete.md).

###### Topics

- [Delete a StackSet (console)](#stacksets-delete-set)

- [Delete a StackSet (AWS CLI)](#stacksets-delete-set-cli)

- [Delete service roles (optional)](#stacksets-delete-roles)

## Delete a StackSet (console)

###### To delete a StackSet

1. Sign in to the AWS Management Console and open the CloudFormation console at
    [https://console.aws.amazon.com/cloudformation](https://console.aws.amazon.com/cloudformation).

2. On the navigation bar at the top of the screen, choose the AWS Region
    you created the StackSet in.

3. On the **StackSets** page, select the StackSet.

4. With the StackSet selected, choose **Delete StackSet** from
    the **Actions** menu.

5. When you are prompted to confirm that you want to delete the StackSet, choose
    **Delete StackSet**.

## Delete a StackSet (AWS CLI)

###### Note

When acting as a delegated administrator, you must include `--call-as
                        DELEGATED_ADMIN` in the command.

###### To delete a StackSet

1. Use the following **delete-stack-set** command. When you
    are prompted to confirm, type `y`, and then press
    **Enter**.

```nohighlight

aws cloudformation delete-stack-set --stack-set-name my-stackset
```

2. Verify that the StackSet was deleted by running the
    **list-stack-sets** command. The results of the
    list-stack-sets command should show your stack with a status of
    `DELETED`.

```nohighlight

aws cloudformation list-stack-sets
```

## Delete service roles (optional)

If you no longer need the IAM service roles that CloudFormation requires to perform
StackSet operations, we recommend that you delete the roles.

_For self-managed StackSets_, the roles you created. For
more information about these roles, see [Grant self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html).

_For service-managed StackSets_, the roles that were
automatically created for StackSets have the suffix
`CloudFormationStackSetsOrgAdmin` in the organization
management account, and `CloudFormationStackSetsOrgMember` in each target
account. For more information, see [Service-linked roles](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-activate-trusted-access.html#stacksets-orgs-service-linked-roles).

###### To delete a service role (console)

1. Sign in to the AWS Management Console and open the IAM console at [https://console.aws.amazon.com/iam/](https://console.aws.amazon.com/iam).

2. From the navigation pane, choose **Roles**, and then fill
    the check box next to the role that you want to delete.

3. In the **Role actions** menu at the top of the page,
    choose **Delete role**.

4. In the confirmation dialog box, choose **Yes, Delete**.
    If you are sure, you can proceed with the deletion even if the service last
    accessed data is still loading.

###### To delete a service role (AWS CLI)

- Use the following **delete-role** command. When you are
prompted to confirm, type `y`, and then press
**Enter**.

```nohighlight

aws iam delete-role --role-name role name
```

For more information about deleting roles, see [Delete roles or instance\
profiles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_manage_delete.html) in the _IAM User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Delete stacks from
StackSets

Target account gates
