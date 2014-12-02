Setup
-----

Before this can be built you need to install libusb-dev and rtl-sdr.

### Install libusb
sudo apt-get install libusb-1.0-0-dev

### Install rtl-sdr
git clone git://git.osmocom.org/rtl-sdr.git
install cmake if needed (sudo apt-get install cmake)

cd rtl-sdr/
mkdir build
cd build
cmake ../
make
sudo make install
sudo ldconfig

### Allow non-root access to USB devicea
Plug in the USB SDR.
Run the following command to list your USB devices: *lsusb*
You should see a line that looks like this:
  Bus 001 Device 004: ID 0bda:2838 Realtek Semiconductor Corp.

The important values are *0bda* (the Vendor ID) and *2838* (the product ID).

As root, create a new file named *20.rtlsdr.rules* in the */etc/udev/rules.d* directory with the following content (replace the Vendor ID and the product ID with the values from the *ldusb* command if they're different):
  SUBSYSTEM=="usb",ATTRS{idVendor}=="0bda",ATTRS{idProduct}=="2838",GROUP="adm",MODE="0666",SYMLINK+="rtl_sdr"

Unplug the USB device then restart udev: *sudo restart udev*

Plug back in the USB device.

Source
------

Get the source code by running the following command after installing Go and setting up your GOPATH environment variable (see https://golang.org/doc/code.html#Organization):
  go get github.com/inpursuit/goadsb
