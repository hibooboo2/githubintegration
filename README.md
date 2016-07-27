# Github Issues.

This is a github integration around creating a workflow that is automated to move labels and update statuses to conform to a certian flow when working within github.


If an pr is merged and closed then any issues in the body that are referenced like:
    test #4
    test user/repo#4
    tests #4
    tests user/repo#4

The issue will be labeled testing.
