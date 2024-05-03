# GSC Attestation App
This is a sample application to demonstrate the GSC enclave attestation with Intel® Trust Authority using [trustauthority-client](https://github.com/intel/trustauthority-client-for-go/)

## Instructions for Ubuntu
Use below command to install all the dependencies necessary to run this sample app.

```sh
chmod +x install.sh
./install.sh
```

## Build
Once the above install script is done, use below steps to build and run the app

```sh
make gramine
```
```sh
sudo docker run --env-file config.env --device=/dev/sgx_enclave -v /var/run/aesmd/aesm.socket:/var/run/aesmd/aesm.socket gsc-attestation-app-gsc:v0.1.0
```

## Config Definition
```env
SGX_AESM_ADDR=1
TRUSTAUTHORITY_URL=https://portal.trustauthority.intel.com
TRUSTAUTHORITY_API_URL=https://api.trustauthority.intel.com
TRUSTAUTHORITY_API_KEY=<trustauthority attestation api key>
```
