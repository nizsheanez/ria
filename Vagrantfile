VAGRANTFILE_API_VERSION = "2"

Vagrant.require_version ">= 1.6.3"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.define "boot2docker"

  config.vm.box = "yungsang/boot2docker"
  config.vm.box_check_update = false

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.synced_folder "./../../", "/gopath", type: "nfs"
  config.vm.synced_folder "/var/lib/mysql", "/var/lib/mysql", type: "nfs"

  # Uncomment below to use more than one instance at once
  # config.vm.network :forwarded_port, guest: 2375, host: 2375, auto_correct: true

  # Fix busybox/udhcpc issue
  config.vm.provision :shell do |s|
    s.inline = <<-EOT
      if ! grep -qs ^nameserver /etc/resolv.conf; then
        sudo /sbin/udhcpc
      fi
      cat /etc/resolv.conf
    EOT
  end

  # Adjust datetime after suspend and resume
  config.vm.provision :shell do |s|
    s.inline = <<-EOT
      sudo /usr/local/bin/ntpclient -s -h pool.ntp.org
      date
    EOT
  end

  config.vm.network :forwarded_port, guest: 8080, host: 8080
  config.vm.network :forwarded_port, guest: 80, host: 80
  config.vm.network :forwarded_port, guest: 3306, host: 3306
  config.vm.network :forwarded_port, guest: 9000, host: 9000
end