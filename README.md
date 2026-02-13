# clpcfdiff
Get the diff of the cluster configuration file (clp.conf).

# How to Build
1. Clone this repository.
   ```sh
   git clone https://github.com/EXPRESSCLUSTER/clpcfdiff.git
   ```
1. Move to `src` directory.
   ```sh
   cd clpcfdiff/src
   ```
1. Initialize and build clpcfdiff.
   ```sh
   go mod init clpcfdiff
   ```
   ```sh
   go mod tidy
   ```
   ```sh
   go build
   ```

# How to Use
1. Run `clpcfdiff` command as below.
   ```
   ./clpcfdiff <previous conf file> <current conf file>
   ```
   - Sample conf files
     - [Previous conf file](conf/01_previous/)
     - [Current conf file](conf/02_current/)
1. You can get the result as below.
   ```csv
   File1_Path,File1_Value,File2_Path,File2_Value
   /root,15,/root,19
   ,,/root/group[@name='failover2'],1
   ,,/root/group[@name='failover2']/comment,
   ,,/root/group[@name='failover2']/gid,1
   ,,/root/group[@name='failover2']/resource[@name='md@md2'],
   ,,/root/monitor/mdnw[@name='mdnw2'],md2
   ,,/root/monitor/mdnw[@name='mdnw2']/comment,
   ,,/root/monitor/mdnw[@name='mdnw2']/parameters,md2
   ,,/root/monitor/mdnw[@name='mdnw2']/parameters/object,md2
   ,,/root/monitor/mdnw[@name='mdnw2']/relation,LocalServer
   ,,/root/monitor/mdnw[@name='mdnw2']/relation/name,LocalServer
   ,,/root/monitor/mdnw[@name='mdnw2']/relation/type,cls
   ,,/root/monitor/mdnw[@name='mdnw2']/target,
   ,,/root/monitor/mdw[@name='mdw2'],md2
   ,,/root/monitor/mdw[@name='mdw2']/comment,
   ,,/root/monitor/mdw[@name='mdw2']/parameters,md2
   ,,/root/monitor/mdw[@name='mdw2']/parameters/object,md2
   ,,/root/monitor/mdw[@name='mdw2']/relation,LocalServer
   ,,/root/monitor/mdw[@name='mdw2']/relation/name,LocalServer
   ,,/root/monitor/mdw[@name='mdw2']/relation/type,cls
   ,,/root/monitor/mdw[@name='mdw2']/target,
   ,,/root/resource/md[@name='md2'],
   ,,/root/resource/md[@name='md2']/comment,
   ,,/root/resource/md[@name='md2']/parameters,ext4
   ,,/root/resource/md[@name='md2']/parameters/diskdev,/dev/md2/cp
   ,,/root/resource/md[@name='md2']/parameters/diskdev/cppath,/dev/md2/cp
   ,,/root/resource/md[@name='md2']/parameters/diskdev/dppath,/dev/md2/dp
   ,,/root/resource/md[@name='md2']/parameters/fs,ext4
   ,,/root/resource/md[@name='md2']/parameters/mddriver,29052
   ,,/root/resource/md[@name='md2']/parameters/mddriver/ack2port,29072
   ,,/root/resource/md[@name='md2']/parameters/mddriver/hbport,29032
   ,,/root/resource/md[@name='md2']/parameters/mddriver/port,29052
   ,,/root/resource/md[@name='md2']/parameters/mount,/mnt/md2
   ,,/root/resource/md[@name='md2']/parameters/mount/point,/mnt/md2
   ,,/root/resource/md[@name='md2']/parameters/netdev[@id='0'],mdc1
   ,,/root/resource/md[@name='md2']/parameters/netdev[@id='0']/device,400
   ,,/root/resource/md[@name='md2']/parameters/netdev[@id='0']/mdcname,mdc1
   ,,/root/resource/md[@name='md2']/parameters/netdev[@id='0']/priority,0
   ,,/root/resource/md[@name='md2']/parameters/nmppath,/dev/NMP2
   /root/trekking/configid,5311010.5492761741466077,/root/trekking/configid,5311010.11561841146518159
   /root/webmgr,15,/root/webmgr,19
   /root/webmgr/client,15,/root/webmgr/client,19
   /root/webmgr/client/objectnumber,15,/root/webmgr/client/objectnumber,19
   ```
