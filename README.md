# Command Injection Example
 Command Injection Example written in Golang - Also will be a test to see if Synk recognizes the vulnerability in the Golang program and will create a PR to fix it.

 This program listens on the port 8080 and takes a `command` query parameter from the URL. The program then executes the `ls` command with the given `command` parameter abd returns the output to the user.

 However, this program is vulnerable command injection attacks, as a hacker can pass through a malicious command such as `:rm -rf /` which would delete all files in the server.

 A general fix for basic command injection vulnerabilities is validating and sanitizing the user input, as well as avoid passing user inputs directly into system commands without proper escaping and quoting.

As you can see, Synk.io wasn't able to detect that this code from Golang is suseptible to command injection, however it has done very well for everything else I've needed. It's a perfect addition that every programmer needs to assist them. One of the best tools out there at the moment.

<img width="772" alt="Capture" src="https://user-images.githubusercontent.com/130214281/235314950-48414803-f34d-4af8-a535-60961a2549b8.PNG">
<img width="763" alt="Capture1" src="https://user-images.githubusercontent.com/130214281/235314951-f3d150ca-2b77-40bf-8716-04b8269b3be8.PNG">
