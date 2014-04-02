# fleet-talk

files for Docker St. Louis talk on [Fleet](https://github.com/coreos/fleet).

## Files

 - `app/`: a tiny Go web server that writes the value of the `HOST` env variable to the response.
 - `hellogo.web.service`: a systemd unit file to run the app
 - `hellogo.reg.service`: a systemd unit file to register the app with etcd
 - `splode`: script to "parallelize" a service by making n instances of the unit file and replacing `__ID__` with the value of n. So `./splode hellogo.web.service 3` would create 3 instances of `hellogo.web.service`, named `hellogo.web.{1,2,3}.service`. This is only necessary until [issue 212 in fleet](https://github.com/coreos/fleet/issues/212) is solved.

See also the [slide deck on blazon](http://presentboldly.com/brianhicks/managing-a-coreos-cluster-with-fleet)
