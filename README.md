# Cloud Foundry Application Logs Tailer

The `cf-app-logstail` binary is used for tailing logfiles written by applications which do not satisfy the 12-factor logging to STDOUT/STDERR principle. This can be of help while applications are being re-written or for legacy applications to which you cannot change the code.

## Instructions

First set an environment variable in your application's space specifying what logs to tail into loggregator.

For example, to send `log/production.log` to STDOUT (the default) while sending log/error.log to STDERR we would do the following:

    cf set-env #app_name logstail "log/production.log:log/error.log|2"

Next add `cf-logstail` to start with the application so that it logs it's STDOUT/STDERR to the same streams as the application (eg. loggregator).

    TODO: Figure out the best way to do this.

Push your application / ensure that it restarts to get the new environment variables and the new binary:

    cf push #app_name

Check to make sure the logs are being tailed:

    cf logs #app_name

## Hat-Tip

Thanks to ActiveState for writing their [tail]( https://github.com/ActiveState/tail ) library 