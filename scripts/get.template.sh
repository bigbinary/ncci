#!/bin/sh

OS=$(uname -s)
ARCH=$(uname -m)
VERSION="VERSION_PLACEHOLDER"

echo "Downloading neetoCI CLI release ${VERSION} for ${OS}_${ARCH} ..."
echo ""

readonly TMP_DIR="$(mktemp -d -t ncci-XXXX)"

trap cleanup EXIT

cleanup() {
  rm -rf "${TMP_DIR}"
}

curl --fail -L "https://github.com/bigbinary/neeto-ci-cli/releases/download/${VERSION}/ncci_${OS}_${ARCH}.tar.gz" -o ${TMP_DIR}/ncci.tar.gz


if ! [ $? -eq 0 ]; then
  echo ""
  echo "[error] Failed to download ncci CLI release for $OS $ARCH."
  echo ""
  echo "Supported versions of the neetoCI CLI are:"
  echo " - Linux_x86_64"
  echo " - Linux_i386"
  echo " - Linux_arm64"
  echo " - Darwin_i386"
  echo " - Darwin_x86_64"
  echo " - Darwin_arm64"
  echo ""
  exit 1
fi


tar -xzf ${TMP_DIR}/ncci.tar.gz -C ${TMP_DIR}
sudo chmod +x ${TMP_DIR}/ncci
sudo mv ${TMP_DIR}/ncci /usr/local/bin/


echo ""
echo "neetoCI CLI ${VERSION} for ${OS}_${ARCH} installed."
