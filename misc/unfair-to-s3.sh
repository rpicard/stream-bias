#!/bin/bash

# WARNING: THIS IS NOT READY. I DO NOT DO ENOUGH SHELL SCRIPTING
# TO BE CONFIDENT IN THIS YET. THIS SPINS UP EC2 INSTANCES, SO IT
# IS GOING TO BE SPENDING MONEY.

# we rely on s3put and kill_instance from boto to be available

# get the instance ID

$instance = `curl http://169.254.169.254/latest/meta-data/instance-id`

# 1. run unfair for 100 000 000 samples
#    this should take about 20 minutes on a t1.micro instances
#    eventually going to raise this, but for now we will stick
#    to a relatively small number of samples

/usr/local/bin/unfair -c rc4 -f json -s 100000000 > "${instance}-unfair.json"

# 2. use s3put to stick the output in an s3 bucket
#    manually making the s3 bucket called "unfair-test1" for now
#    creds are in the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
#    environment variables

/usr/local/bin/s3put -b "unfair-test1" "${instance}-unfair.json"

# 3. use kill_instance to commit suicide
#    pretty sure this should work with spot instances and all just
#    like any other instance

/usr/local/bin/kill_instance $instance

# 4. oh fuck there really is an after life
