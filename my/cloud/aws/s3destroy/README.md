# s3destroy

This was inspired by the Python-based ["s3wipe"](https://github.com/eschwim/s3wipe "So cool!"),
which whilst was almost exactly what I needed, wasn't written for some more-modern
AWS-SDK capabilities, like profiles or MFA.

Some of the ideas for dealing with very-large S3 repos resonated with me, but since
I haven't written Python in-anger for a few years and it's not one of the blessed
languages in my job, it seemed useful to port it to Go.
