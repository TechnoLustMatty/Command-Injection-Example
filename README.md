# Command Injection Example
 Command Injection Example written in Golang - Also will be a test to see if Synk recognizes the vulnerability in the Golang program and will create a PR to fix it.

 This program listens on the port 8080 and takes a `command` query parameter from the URL. The program then executes the `ls` command with the given `command` parameter abd returns the output to the user.

 However, this program is vulnerable command injection attacks, as a hacker can pass through a malicious command such as `:rm -rf /` which would delete all files in the server.

 A general fix for basic command injection vulnerabilities is validating and sanitizing the user input, as well as avoid passing user inputs directly into system commands without proper escaping and quoting.
