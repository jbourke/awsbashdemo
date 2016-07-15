# Demonstration IP Echo

This is a demo script for performing basic AWS EC2 instance launching and tear down.  It will upload and run any given executable (referred to as a 'project'), and an IP echo binary has been provided.  It's a very simple Go app, built to run on Ubuntu.

This demo was written in just a few hours in accordance with a time limited task, and so is as basic as possible while still hitting all the major points in the process of deploying a service on AWS.  This demo relies on several pre-conditions:

* AWS CLI Tools is installed, and the version is 1.3 or higher.
* AWS user config is setup.  The profile to use can be set within the demo script if it is not 'default'.
* The used AWS account must have a default VPC or be EC2 Classic.  This does not include network setup.
* Built on a Linux desktop for a Linux environment.  Has not been tested on OSX yet.

## Why BASH?

Bash is my fast prototyper.  When most of what I am doing can be provided by command line tools, I keep it simple.  AWS CLI does a ton of heavy lifting, and the process of bringing up a server and uploading something involves easy commands that are primarily one liners.  Unless you are trying to test specific technologies, use what makes you comfortable so you can best focus.

Thus, engaging more formal automation in such an open ended demo that is severely time limited offers very little benefit from the perspective of demonstrating an understanding of the steps and pitfalls associated with working with AWS.  The purpose of a demo exercise is to show knowledge of the underlying steps.  While ansible, puppet, chef, terraform, cloudformation, or other automation tools should be used in an enterprise environment, it is useful to demonstrate knowledge of how those tools operate at low levels.

### What is the index.html included in the repo?

This was something whipped up in a few minutes to demonstrate a more practical way to accomplish echoing a client IP.  It can be hosted on S3 for considerably less trouble.  This is certainly not in the spirit of a skill demonstration.  However, in an expanded problem it could easily be a buttress solution; one which can be accomplished in very short order so that time can be gained to work on a better or more scalable solution.  In some cases, the best answer isn't the one which covers all possibilities elegantly, but many reasonably.

## Running the Demo

0. Checkout the repo and cd into the directory
1. run `./gvawsdemo new ipresponder`
⋅⋅1. This creates a new key pair, security group, and instance
⋅⋅2. Waits for the instance to become available, through both instance state and ssh connectivity
⋅⋅3. Once access is confirmed, the project is uploaded and started as a background process
⋅⋅4. The url where you can see your IP echoed back is shown as the last line
2. run `./gvawsdemo update ipresponder`
⋅⋅1. Confirm the instance exists
⋅⋅2. Uploads the new project, kills the old process, and starts it anew
⋅⋅3. The url where you can see your IP echoed back is shown as the last line
3. run `./gvawsdemo teardown`
⋅⋅1. Terminate the active instance and wait for it to report 'terminated'
⋅⋅2. Delete the security group, key pair, and local temp folder.

### Troubleshooting - so far
* SSH - To many auth failures: Clear your ssh keyring, `ssh-add -D`
* Can't write to /tmp - Script needs to save the key file locally.  Make sure /tmp is writable by your user.

### Next Steps

* Support for multiple instance could be added by including an ELB in the setup process.  This would involve adding loops in the update and teardown sections.
* Support for multiple project could be added by adding logic around the upload and restart of project.  More commands would need to be added to support adding and removing projects.
* Include custom VPC support.  A new VPC, subnets, routes, and IGW could be made during the first call to `new`.
* Support containers or git repos.  These could be sent as arguments, and a script could replace nohup to perform setup and execution.
* Don't use bash.  Literally any other language would be a better idea if this needs to be used in a formal setting.  Seriously, this is not for prod.

