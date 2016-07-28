# Github Issues.

This is a github integration around creating a workflow that is automated to move labels and update statuses to conform to a certian flow when working within github.


If an pr is merged and closed then any issues in the body that are referenced like:
    test #4
    test user/repo#4
    tests #4
    tests user/repo#4

The issue will be labeled testing.


If a pr does not ref an issue it will have a failed status.
If it does it will succeed.

When a commit is made that refs an issue it will be moved to:
    - In Progress (Or another label -- should be configiable)

Allow all features to be configed with a yml.

When a branch is made like a regex it will move the issue to in progress.
    Default: feature/issue-{IssueNumber} or feature/issue-${Full issue ref}
    Default: use repo branch is on. Can configure default branch.
    Can configure to use another repo by doing: feature/issue-${Full issue ref}

When an issue is refed on a pr move that issue to be labeled something:
    Default: Waiting To Merge
