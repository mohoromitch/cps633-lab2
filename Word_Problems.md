# Word Problems
1. If 'groups' of users were added to the DAC implementation, a new database would have to be created to store the groups.
These groups would have to be related back to individual users in that database, and when checking for permissions, the implementation would not only check the `users_and_permissions.db` file, but also check the groups associated with the user, and traverse the permissions for all associated groups.
Implementation wise there is not a difference between 'groups' and 'roles', in practice it may or may not matter, based on organization setup.

2. User timeout, IP whitelist/blacklist, password lockout after *n* invalid entries are some possible constraints that could be added to section 3.

3. Sessions would be helpful to add convenience to a user's interaction with the system, by reducing the amount of times the user has to enter in a password.
It would be useful for setup in the lab by allowing a user's interaction to be coupled with a client, which could by used to ensure only one login, since if one session is opened, the other would be closed.