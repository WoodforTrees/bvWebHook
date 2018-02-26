# bvWebHook
The WebHook utility for the bitvise FTP server.

To install on a bitvise server.
Install go, and build: **go build**

Create a directory on the BV server.
Copy in the completed webhook application.

Load up the BV configuration tool,and navigate to **Edit advanced Settings**

Select the virtual account or virtual group (for all users)

Then edit the required user/group.

Select the on-upload command

If a single user uncheck the **use group default**

Set the command to point too to the application and set the working directory to the same place.

You do NOT need to add any params to the application. It takes the values direct from BV (injected into the hidden spawned shell)

Save, and exit the configuration tool. 


