# GoCheck
A simple go program that checks http status codes of your site and restarts services when problems are detected.

## Use
Create an hourly cronjob that runs the gocheck binary.

I do it this way because I don't trust anything to keep running for much longer than an hour.

## Future
I would like to add sendmail support so that instead of quitting after 5 failures, it will send an email notification and then quit.

I may add a more complicated conditions/actions set so that you can create an array of status codes and commands to be run when each status code is encountered. (For example, restart cgi for 502, restart php for 500, etc.)
